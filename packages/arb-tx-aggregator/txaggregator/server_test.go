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

package txaggregator

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/message"
	"math/big"
	"testing"
)

func TestPrepareTransactions(t *testing.T) {
	type testCase struct {
		raw    []DecodedBatchTx
		sorted []message.BatchTx
		label  string
	}

	chain := common.RandAddress()
	keys := make([]*ecdsa.PrivateKey, 0)
	for i := 0; i < 10; i++ {
		pk, err := crypto.GenerateKey()
		if err != nil {
			t.Fatal()
		}
		keys = append(keys, pk)
	}

	cases := make([]testCase, 0)
	cases = append(cases, func() testCase {
		decodedTxes := make([]DecodedBatchTx, 0)
		sortedTxes := make([]message.BatchTx, 0)
		for i := 0; i < 10; i++ {
			batchTx := message.NewRandomBatchTx(chain, keys[0])
			batchTx.SeqNum = big.NewInt(int64(i))
			decoded := NewDecodedBatchTx(batchTx, keys[0].PublicKey)
			decodedTxes = append(decodedTxes, decoded)
			sortedTxes = append(sortedTxes, decoded.tx)
		}
		return testCase{
			raw:    decodedTxes,
			sorted: sortedTxes,
			label:  "inorder",
		}
	}())
	cases = append(cases, func() testCase {
		decodedTxes := make([]DecodedBatchTx, 0)
		sortedTxes := make([]message.BatchTx, 0)
		for i := 0; i < 10; i++ {
			batchTx := message.NewRandomBatchTx(chain, keys[0])
			batchTx.SeqNum = big.NewInt(9 - int64(i))
			decoded := NewDecodedBatchTx(batchTx, keys[0].PublicKey)
			decodedTxes = append(decodedTxes, decoded)
		}
		for i := range decodedTxes {
			sortedTxes = append(sortedTxes, decodedTxes[len(decodedTxes)-1-i].tx)
		}
		return testCase{
			raw:    decodedTxes,
			sorted: sortedTxes,
			label:  "reverse",
		}
	}())

	cases = append(cases, func() testCase {
		decodedTxes := make([]DecodedBatchTx, 0)
		sortedTxes := make([]message.BatchTx, 0)
		for i := 0; i < 10; i++ {
			batchTx := message.NewRandomBatchTx(chain, keys[i])
			batchTx.SeqNum = big.NewInt(9 - int64(i))
			decoded := NewDecodedBatchTx(batchTx, keys[i].PublicKey)
			decodedTxes = append(decodedTxes, decoded)
			sortedTxes = append(sortedTxes, decoded.tx)
		}
		return testCase{
			raw:    decodedTxes,
			sorted: sortedTxes,
			label:  "reverseDifferentKeys",
		}
	}())

	for _, tc := range cases {
		t.Run(tc.label, func(t *testing.T) {
			sortedTxesCal := prepareTransactions(tc.raw)
			t.Log("correct:", tc.sorted)
			t.Log("calculated:", sortedTxesCal)
			if len(sortedTxesCal) != len(tc.sorted) {
				t.Fatal("sorted is wrong length")
			}
			for i, tx := range tc.sorted {
				if !tx.Equals(sortedTxesCal[i]) {
					t.Error("tx in wrong order")
					break
				}
			}
		})
	}

}
