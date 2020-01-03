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

import "./Staking.sol";


contract Confirming is Staking {
    // invalid path proof
    string constant PLACE_PATH_PROOF = "PLACE_PATH_PROOF";
    // invalid path proof
    string constant RECOV_PATH_PROOF = "RECOV_PATH_PROOF";
    // Invalid conflict proof
    string constant RECOV_CONFLICT_PROOF = "RECOV_CONFLICT_PROOF";


    bytes32 private latestConfirmedPriv;

    event RollupStakeRefunded(address staker);

    function recoverStakeConfirmed(bytes32[] calldata proof) external {
        _recoverStakeConfirmed(msg.sender, proof);
    }

    function recoverStakeOld(address payable stakerAddress, bytes32[] calldata proof) external {
        require(proof.length > 0);
        _recoverStakeConfirmed(stakerAddress, proof);
    }

    function recoverStakeMooted(
        address payable stakerAddress,
        bytes32 disputableHash,
        bytes32[] calldata latestConfirmedProof,
        bytes32[] calldata nodeProof
    )
        external
    {
        Staker storage staker = getValidStaker(stakerAddress);
        require(
            RollupUtils.isConflict(
                staker.location,
                disputableHash,
                latestConfirmed(),
                latestConfirmedProof,
                nodeProof
            ),
            RECOV_CONFLICT_PROOF
        );
        deleteStakerWithPayout(stakerAddress);

        emit RollupStakeRefunded(stakerAddress);
    }

    function init(
        bytes32 _vmState,
        uint128 _stakeRequirement,
        address _challengeFactoryAddress
    )
        internal
    {
        Staking.init(_stakeRequirement, _challengeFactoryAddress);

        // VM protocol state
        bytes32 vmProtoStateHash = RollupUtils.protoStateHash(_vmState, Value.hashEmptyTuple(), Value.hashEmptyTuple(), 0);
        updateLatestConfirmed(
            RollupUtils.childNodeHash(
                0,
                0,
                0,
                0,
                vmProtoStateHash
            )
        );
    }

    function latestConfirmed() internal view returns (bytes32) {
        return latestConfirmedPriv;
    }

    function updateLatestConfirmed(bytes32 node) internal {
        latestConfirmedPriv = node;
    }

    function _recoverStakeConfirmed(address payable stakerAddress, bytes32[] memory proof) private {
        Staker storage staker = getValidStaker(stakerAddress);
        require(RollupUtils.isPath(staker.location, latestConfirmed(), proof), RECOV_PATH_PROOF);
        deleteStakerWithPayout(stakerAddress);

        emit RollupStakeRefunded(stakerAddress);
    }
}