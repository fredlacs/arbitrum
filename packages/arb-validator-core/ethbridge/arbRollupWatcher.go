/*
 * Copyright 2020, Offchain Labs, Inc.
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

package ethbridge

import (
	"context"
	"errors"
	"math/big"
	"strings"

	errors2 "github.com/pkg/errors"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethbridge/rollup"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"
)

var rollupCreatedID ethcommon.Hash
var stakeCreatedID ethcommon.Hash
var challengeStartedID ethcommon.Hash
var challengeCompletedID ethcommon.Hash
var rollupRefundedID ethcommon.Hash
var rollupPrunedID ethcommon.Hash
var rollupStakeMovedID ethcommon.Hash
var rollupAssertedID ethcommon.Hash
var rollupConfirmedID ethcommon.Hash
var confirmedAssertionID ethcommon.Hash

func init() {
	parsedRollup, err := abi.JSON(strings.NewReader(rollup.ArbRollupABI))
	if err != nil {
		panic(err)
	}
	rollupCreatedID = parsedRollup.Events["RollupCreated"].ID()
	stakeCreatedID = parsedRollup.Events["RollupStakeCreated"].ID()
	challengeStartedID = parsedRollup.Events["RollupChallengeStarted"].ID()
	challengeCompletedID = parsedRollup.Events["RollupChallengeCompleted"].ID()
	rollupRefundedID = parsedRollup.Events["RollupStakeRefunded"].ID()
	rollupPrunedID = parsedRollup.Events["RollupPruned"].ID()
	rollupStakeMovedID = parsedRollup.Events["RollupStakeMoved"].ID()
	rollupAssertedID = parsedRollup.Events["RollupAsserted"].ID()
	rollupConfirmedID = parsedRollup.Events["RollupConfirmed"].ID()
	confirmedAssertionID = parsedRollup.Events["ConfirmedAssertion"].ID()
}

type ethRollupWatcher struct {
	ArbRollup *rollup.ArbRollup

	rollupAddress ethcommon.Address
	client        *ethclient.Client
}

func newRollupWatcher(
	rollupAddress ethcommon.Address,
	client *ethclient.Client,
) (*ethRollupWatcher, error) {
	arbitrumRollupContract, err := rollup.NewArbRollup(rollupAddress, client)
	if err != nil {
		return nil, errors2.Wrap(err, "Failed to connect to arbRollup")
	}

	return &ethRollupWatcher{
		ArbRollup:     arbitrumRollupContract,
		rollupAddress: rollupAddress,
		client:        client,
	}, nil
}

func (vm *ethRollupWatcher) generateTopics() [][]ethcommon.Hash {
	addressIndex := ethcommon.Hash{}
	copy(
		addressIndex[:],
		ethcommon.LeftPadBytes(vm.rollupAddress.Bytes(), 32),
	)
	return [][]ethcommon.Hash{
		{
			stakeCreatedID,
			challengeStartedID,
			challengeCompletedID,
			rollupRefundedID,
			rollupPrunedID,
			rollupStakeMovedID,
			rollupAssertedID,
			rollupConfirmedID,
			confirmedAssertionID,
		},
	}
}

func (vm *ethRollupWatcher) GetEvents(
	ctx context.Context,
	blockId *common.BlockId,
	_ *big.Int,
) ([]arbbridge.Event, error) {
	bh := blockId.HeaderHash.ToEthHash()
	rollupLogs, err := vm.client.FilterLogs(ctx, ethereum.FilterQuery{
		BlockHash: &bh,
		Addresses: []ethcommon.Address{vm.rollupAddress},
		Topics:    vm.generateTopics(),
	})
	if err != nil {
		return nil, err
	}

	events := make([]arbbridge.Event, 0, len(rollupLogs))
	for _, evmLog := range rollupLogs {
		event, err := vm.processEvents(
			getLogChainInfo(evmLog),
			evmLog,
		)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

func (vm *ethRollupWatcher) GetAllEvents(
	ctx context.Context,
	fromBlock *big.Int,
	toBlock *big.Int,
) ([]arbbridge.Event, error) {
	inboxLogs, err := vm.client.FilterLogs(ctx, ethereum.FilterQuery{
		FromBlock: fromBlock,
		ToBlock:   toBlock,
		Addresses: []ethcommon.Address{vm.rollupAddress},
		Topics:    vm.generateTopics(),
	})
	if err != nil {
		return nil, err
	}

	events := make([]arbbridge.Event, 0, len(inboxLogs))
	for _, evmLog := range inboxLogs {
		event, err := vm.processEvents(getLogChainInfo(evmLog), evmLog)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

func (vm *ethRollupWatcher) processEvents(
	chainInfo arbbridge.ChainInfo,
	ethLog types.Log,
) (arbbridge.Event, error) {
	switch ethLog.Topics[0] {
	case stakeCreatedID:
		eventVal, err := vm.ArbRollup.ParseRollupStakeCreated(ethLog)
		if err != nil {
			return nil, err
		}
		return arbbridge.StakeCreatedEvent{
			ChainInfo: chainInfo,
			Staker:    common.NewAddressFromEth(eventVal.Staker),
			NodeHash:  eventVal.NodeHash,
		}, nil
	case challengeStartedID:
		eventVal, err := vm.ArbRollup.ParseRollupChallengeStarted(ethLog)
		if err != nil {
			return nil, err
		}
		return arbbridge.ChallengeStartedEvent{
			ChainInfo:  chainInfo,
			Asserter:   common.NewAddressFromEth(eventVal.Asserter),
			Challenger: common.NewAddressFromEth(eventVal.Challenger),
			ChallengeType: valprotocol.ChildType(
				eventVal.ChallengeType.Uint64(),
			),
			ChallengeContract: common.NewAddressFromEth(
				eventVal.ChallengeContract,
			),
		}, nil
	case challengeCompletedID:
		eventVal, err := vm.ArbRollup.ParseRollupChallengeCompleted(ethLog)
		if err != nil {
			return nil, err
		}
		return arbbridge.ChallengeCompletedEvent{
			ChainInfo: chainInfo,
			Winner:    common.NewAddressFromEth(eventVal.Winner),
			Loser:     common.NewAddressFromEth(eventVal.Loser),
			ChallengeContract: common.NewAddressFromEth(
				eventVal.ChallengeContract,
			),
		}, nil
	case rollupRefundedID:
		eventVal, err := vm.ArbRollup.ParseRollupStakeRefunded(ethLog)
		if err != nil {
			return nil, err
		}
		return arbbridge.StakeRefundedEvent{
			ChainInfo: chainInfo,
			Staker:    common.NewAddressFromEth(eventVal.Staker),
		}, nil
	case rollupPrunedID:
		eventVal, err := vm.ArbRollup.ParseRollupPruned(ethLog)
		if err != nil {
			return nil, err
		}
		return arbbridge.PrunedEvent{
			ChainInfo: chainInfo,
			Leaf:      eventVal.Leaf,
		}, nil
	case rollupStakeMovedID:
		eventVal, err := vm.ArbRollup.ParseRollupStakeMoved(ethLog)
		if err != nil {
			return nil, err
		}
		return arbbridge.StakeMovedEvent{
			ChainInfo: chainInfo,
			Staker:    common.NewAddressFromEth(eventVal.Staker),
			Location:  eventVal.ToNodeHash,
		}, nil
	case rollupAssertedID:
		eventVal, err := vm.ArbRollup.ParseRollupAsserted(ethLog)
		if err != nil {
			return nil, err
		}
		params := &valprotocol.AssertionParams{
			NumSteps: eventVal.NumSteps,
			TimeBounds: &protocol.TimeBounds{
				LowerBoundBlock:     common.NewTimeBlocks(eventVal.TimeBounds[0]),
				UpperBoundBlock:     common.NewTimeBlocks(eventVal.TimeBounds[1]),
				LowerBoundTimestamp: eventVal.TimeBounds[2],
				UpperBoundTimestamp: eventVal.TimeBounds[3],
			},
			ImportedMessageCount: eventVal.ImportedMessageCount,
		}
		claim := &valprotocol.AssertionClaim{
			AfterInboxTop:         eventVal.Fields[2],
			ImportedMessagesSlice: eventVal.Fields[3],
			AssertionStub: &valprotocol.ExecutionAssertionStub{
				AfterHash:        eventVal.Fields[4],
				DidInboxInsn:     eventVal.DidInboxInsn,
				NumGas:           eventVal.NumArbGas,
				FirstMessageHash: [32]byte{},
				LastMessageHash:  eventVal.Fields[5],
				FirstLogHash:     [32]byte{},
				LastLogHash:      eventVal.Fields[6],
			},
		}
		return arbbridge.AssertedEvent{
			ChainInfo:    chainInfo,
			PrevLeafHash: eventVal.Fields[0],
			Disputable: valprotocol.NewDisputableNode(
				params,
				claim,
				eventVal.Fields[1],
				eventVal.InboxCount,
			),
		}, nil
	case rollupConfirmedID:
		eventVal, err := vm.ArbRollup.ParseRollupConfirmed(ethLog)
		if err != nil {
			return nil, err
		}
		return arbbridge.ConfirmedEvent{
			ChainInfo: chainInfo,
			NodeHash:  eventVal.NodeHash,
		}, nil
	case confirmedAssertionID:
		eventVal, err := vm.ArbRollup.ParseConfirmedAssertion(ethLog)
		if err != nil {
			return nil, err
		}
		return arbbridge.ConfirmedAssertionEvent{
			ChainInfo:   chainInfo,
			LogsAccHash: hashSliceToHashes(eventVal.LogsAccHash),
		}, nil
	default:
		return nil, errors2.New("unknown arbitrum event type")
	}
}

func (vm *ethRollupWatcher) GetParams(
	ctx context.Context,
) (valprotocol.ChainParams, error) {
	callOpts := &bind.CallOpts{Context: ctx}
	rawParams, err := vm.ArbRollup.VmParams(callOpts)
	if err != nil {
		return valprotocol.ChainParams{}, err
	}
	stakeRequired, err := vm.ArbRollup.GetStakeRequired(callOpts)
	if err != nil {
		return valprotocol.ChainParams{}, err
	}
	return valprotocol.ChainParams{
		StakeRequirement: stakeRequired,
		GracePeriod: common.TimeTicks{
			Val: rawParams.GracePeriodTicks,
		},
		MaxExecutionSteps:       rawParams.MaxExecutionSteps,
		MaxBlockBoundsWidth:     rawParams.MaxBlockBoundsWidth,
		MaxTimestampBoundsWidth: rawParams.MaxTimestampBoundsWidth,
		ArbGasSpeedLimitPerTick: rawParams.ArbGasSpeedLimitPerTick.Uint64(),
	}, nil
}

func (vm *ethRollupWatcher) InboxAddress(
	ctx context.Context,
) (common.Address, error) {
	addr, err := vm.ArbRollup.GlobalInbox(&bind.CallOpts{Context: ctx})
	return common.NewAddressFromEth(addr), err
}

func (vm *ethRollupWatcher) GetCreationInfo(
	ctx context.Context,
) (common.Hash, *common.BlockId, common.Hash, error) {
	addressIndex := ethcommon.Hash{}
	copy(
		addressIndex[:],
		ethcommon.LeftPadBytes(vm.rollupAddress.Bytes(), 32),
	)
	logs, err := vm.client.FilterLogs(ctx, ethereum.FilterQuery{
		Addresses: []ethcommon.Address{vm.rollupAddress},
		Topics:    [][]ethcommon.Hash{{rollupCreatedID}},
	})
	if err != nil {
		return common.Hash{}, nil, common.Hash{}, err
	}
	if len(logs) != 1 {
		return common.Hash{},
			nil,
			common.Hash{},
			errors.New("more than one chain created with same address")
	}
	ev, err := vm.ArbRollup.ParseRollupCreated(logs[0])
	if err != nil {
		return common.Hash{}, nil, common.Hash{}, err
	}

	return common.NewHashFromEth(logs[0].TxHash), getLogBlockID(logs[0]), ev.InitVMHash, nil
}

func (vm *ethRollupWatcher) GetVersion(ctx context.Context) (string, error) {
	return vm.ArbRollup.VERSION(&bind.CallOpts{Context: ctx})
}

func (vm *ethRollupWatcher) IsStaked(address common.Address) (bool, error) {
	return vm.ArbRollup.IsStaked(nil, address.ToEthAddress())
}
