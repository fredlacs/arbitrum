/*
 * Copyright 2019, Offchain Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

pragma solidity ^0.5.3;

import "./RollupUtils.sol";
import "../libraries/RollupTime.sol";

import "../challenge/ChallengeUtils.sol";
import "../challenge/ChallengeType.sol";
import "../challenge/IChallengeFactory.sol";

import "../arch/Protocol.sol";


contract Staking is ChallengeType {

    // VM already initialized"
    string constant INIT_TWICE = "INIT_TWICE";
    // Challenge factory must be nonzero
    string constant INIT_NONZERO = "INIT_NONZERO";

    // Invalid staker
    string constant INV_STAKER = "INV_STAKER";

    // must supply stake value
    string constant STK_AMT = "STK_AMT";
    // Staker already exists
    string constant ALRDY_STAKED = "ALRDY_STAKED";

    // Challenge can only be resolved by spawned contract
    string constant RES_CHAL_SENDER = "RES_CHAL_SENDER";

    // Stakers must have a conflict over pending top
    string constant PND_CHAL_TYPE = "PND_CHAL_TYPE";

    // Stakers must have a conflict over messages
    string constant MSGS_CHAL_TYPE = "MSGS_CHAL_TYPE";

    // Stakers must have a conflict over execution
    string constant EXEC_CHAL_TYPE = "EXEC_CHAL_TYPE";

    // staker1 staked after deadline
    string constant STK1_DEADLINE = "STK1_DEADLINE";
    // staker2 staked after deadline
    string constant STK2_DEADLINE = "STK2_DEADLINE";
    // staker1 already in a challenge
    string constant STK1_IN_CHAL = "STK1_IN_CHAL";
    // staker2 already in a challenge
    string constant STK2_IN_CHAL = "STK1_IN_CHAL";
    // Child types must be ordered
    string constant TYPE_ORDER = "TYPE_ORDER";
    // Invalid child type
    string constant INVLD_CHLD_TYPE = "INVLD_CHLD_TYPE";
    // Invalid staker1 proof
    string constant STK1_PROOF = "STK1_PROOF";
    // Invalid staker2 proof
    string constant STK2_PROOF = "STK2_PROOF";

    uint256 internal constant VALID_CHILD_TYPE = 3;
    uint256 internal constant MAX_CHILD_TYPE = 3;

    IChallengeFactory public challengeFactory;

    struct Staker {
        bytes32 location;
        uint128 creationTimeBlocks;
        bool inChallenge;
    }

    uint128 private stakeRequirement;
    mapping(address => Staker) private stakers;
    uint256 private stakerCount;

    event RollupStakeCreated(
        address staker,
        bytes32 nodeHash
    );

    event RollupChallengeStarted(
        address asserter,
        address challenger,
        uint256    challengeType,
        address challengeContract
    );

    event RollupChallengeCompleted(
        address challengeContract,
        address winner,
        address loser
    );

    function startChallenge(
        address payable asserterAddress,
        address payable challengerAddress,
        bytes32 node,
        uint256 disputableDeadlineTicks,
        uint256[2] memory stakerPositions,
        bytes32[2] memory vmProtoHashes,
        bytes32[] memory proof1,
        bytes32[] memory proof2,
        bytes32 challenge1DataHash,
        uint128 challenge1PeriodTicks,
        bytes32 challenge2NodeHash
    )
        public
    {
        Staker storage asserter = getValidStaker(asserterAddress);
        Staker storage challenger = getValidStaker(challengerAddress);

        require(RollupTime.blocksToTicks(asserter.creationTimeBlocks) < disputableDeadlineTicks, STK1_DEADLINE);
        require(RollupTime.blocksToTicks(challenger.creationTimeBlocks) < disputableDeadlineTicks, STK2_DEADLINE);
        require(!asserter.inChallenge, STK1_IN_CHAL);
        require(!challenger.inChallenge, STK2_IN_CHAL);
        require(stakerPositions[0] < stakerPositions[1], TYPE_ORDER);
        require(
            RollupUtils.isPath(
                RollupUtils.childNodeHash(
                    node,
                    disputableDeadlineTicks,
                    keccak256(
                        abi.encodePacked(
                            challenge1DataHash,
                            challenge1PeriodTicks
                        )
                    ),
                    stakerPositions[0],
                    vmProtoHashes[0]
                ),
                asserter.location,
                proof1
            ),
            STK1_PROOF
        );
        require(
            RollupUtils.isPath(
                RollupUtils.childNodeHash(
                    node,
                    disputableDeadlineTicks,
                    challenge2NodeHash,
                    stakerPositions[1],
                    vmProtoHashes[1]
                ),
                challenger.location,
                proof2
            ),
            STK2_PROOF
        );

        asserter.inChallenge = true;
        challenger.inChallenge = true;

        address newChallengeAddr = challengeFactory.createChallenge(
            asserterAddress,
            challengerAddress,
            challenge1PeriodTicks,
            challenge1DataHash,
            stakerPositions[0]
        );

        emit RollupChallengeStarted(
            asserterAddress,
            challengerAddress,
            stakerPositions[0],
            newChallengeAddr
        );
    }

    function resolveChallenge(address payable winner, address loser) external {
        address sender = msg.sender;
        bytes32 codehash;
        assembly { codehash := extcodehash(sender) }
        address challengeContract1 = challengeFactory.generateCloneAddress(address(winner), loser, codehash);
        address challengeContract2 = challengeFactory.generateCloneAddress(address(winner), loser, codehash);
        require(challengeContract1 == msg.sender || challengeContract2 == msg.sender, RES_CHAL_SENDER);
        Staker storage winningStaker = getValidStaker(address(winner));
        winner.transfer(stakeRequirement / 2);
        winningStaker.inChallenge = false;
        deleteStaker(loser);

        emit RollupChallengeCompleted(msg.sender, address(winner), loser);
    }

    function init(
        uint128 _stakeRequirement,
        address _challengeFactoryAddress
    )
        internal
    {
        require(address(challengeFactory) == address(0), INIT_TWICE);
        require(_challengeFactoryAddress != address(0), INIT_NONZERO);

        challengeFactory = IChallengeFactory(_challengeFactoryAddress);

        // VM parameters
        stakeRequirement = _stakeRequirement;
    }

    function getValidStaker(address _stakerAddress) internal view returns (Staker storage) {
        Staker storage staker = stakers[_stakerAddress];
        require(staker.location != 0x00, INV_STAKER);
        return staker;
    }

    function createStake(
        bytes32 location
    )
        internal
    {
        require(msg.value == stakeRequirement, STK_AMT);
        require(stakers[msg.sender].location != 0x00, ALRDY_STAKED);
        stakers[msg.sender] = Staker(
            location,
            uint128(block.number),
            false
        );
        stakerCount++;

        emit RollupStakeCreated(msg.sender, location);
    }

    function deleteStakerWithPayout(address payable _stakerAddress) internal {
        deleteStaker(_stakerAddress);
        _stakerAddress.transfer(stakeRequirement);
    }

    function getStakerCount() internal view returns(uint) {
        return stakerCount;
    }

    function deleteStaker(address _stakerAddress) private {
        delete stakers[_stakerAddress];
        stakerCount--;
    }
}