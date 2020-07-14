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

package message

import (
	"bytes"
	"fmt"
	common2 "github.com/ethereum/go-ethereum/common"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

func TestL2MessageSerialization(t *testing.T) {
	pk, err := crypto.GenerateKey()
	if err != nil {
		t.Fatal(err)
	}
	tx := NewRandomTransaction()
	txData := tx.AsData()
	if len(txData) != TransactionHeaderSize+len(tx.Data) {
		t.Error("serialized tx has incorrect size")
	}

	l2Messages := []AbstractL2Message{
		tx,
		NewRandomContractTransaction(),
		NewRandomCall(),
		NewRandomTransactionBatch(10, common.RandAddress(), pk),
	}

	for _, msg := range l2Messages {
		t.Run(fmt.Sprintf("%T", msg), func(t *testing.T) {
			l2 := L2Message{Msg: msg}
			data := l2.AsData()
			decoded, err := NewL2MessageFromData(data)
			if err != nil {
				t.Fatal(err)
			}
			if bytes.Equal(decoded.AsData(), data) {
				t.Fatal("decoded l2 message not equal")
			}
		})
	}

}

func TestMarshaledBytesHash(t *testing.T) {
	data := make([]byte, 0)
	for i := 0; i < 67; i++ {
		data = append(data, byte(i))
	}
	hash := marshaledBytesHash(data)
	correct := common.NewHashFromEth(common2.HexToHash("0x4fc384a19926e9ff7ec8f2376a0d146dc273031df1db4d133236d209700e4780"))
	if hash != correct {
		t.Fatal("incorrect result", hash, correct)
	}
}
