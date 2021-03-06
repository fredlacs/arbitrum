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
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"errors"
	"log"
	"math/big"
	"sort"
	"sync"
	"time"

	errors2 "github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/message"
)

const maxTransactions = 200

const signatureLength = 65
const recoverBitPos = signatureLength - 1

type DecodedBatchTx struct {
	tx     message.BatchTx
	pubkey []byte
}

func NewDecodedBatchTx(tx message.BatchTx, key ecdsa.PublicKey) DecodedBatchTx {
	return DecodedBatchTx{
		tx:     tx,
		pubkey: elliptic.Marshal(crypto.S256(), key.X, key.Y),
	}
}

type Server struct {
	rollupAddress common.Address
	globalInbox   arbbridge.GlobalInbox

	sync.Mutex
	valid        bool
	transactions []DecodedBatchTx
}

// NewServer returns a new instance of the Server class
func NewServer(
	ctx context.Context,
	globalInbox arbbridge.GlobalInbox,
	rollupAddress common.Address,
) *Server {
	server := &Server{
		rollupAddress: rollupAddress,
		globalInbox:   globalInbox,
		valid:         true,
	}

	go func() {
		ticker := time.NewTicker(time.Second * 5)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return

			case <-ticker.C:
				server.Lock()
				// Keep sending in spin loop until we can't anymore
				sentFull := false
				for server.valid && len(server.transactions) >= maxTransactions {
					server.sendBatch(ctx)
					sentFull = true
				}
				// If we have've sent any batches, send a partial
				if !sentFull && server.valid && len(server.transactions) > 0 {
					server.sendBatch(ctx)
				}
				server.Unlock()
			}
		}
	}()
	return server
}

// prepareTransactions reorders the transactions such that the position of each
// user is maintained, but the transactions of that user are swapped to be in
// sequence number order
func prepareTransactions(txes []DecodedBatchTx) []message.BatchTx {
	transactionsBySender := make(map[string][]DecodedBatchTx)
	for _, tx := range txes {
		senderPubkey := hexutil.Encode(tx.pubkey)
		transactionsBySender[senderPubkey] = append(transactionsBySender[senderPubkey], tx)
	}

	for _, txes := range transactionsBySender {
		sort.SliceStable(txes, func(i, j int) bool {
			return txes[i].tx.SeqNum.Cmp(txes[j].tx.SeqNum) < 0
		})
	}

	batchTxes := make([]message.BatchTx, 0, len(txes))
	for _, tx := range txes {
		senderPubkey := hexutil.Encode(tx.pubkey)
		nextTx := transactionsBySender[senderPubkey][0]
		transactionsBySender[senderPubkey] = transactionsBySender[senderPubkey][1:]
		batchTxes = append(batchTxes, nextTx.tx)
	}
	return batchTxes
}

func (m *Server) sendBatch(ctx context.Context) {
	var txes []DecodedBatchTx

	if len(m.transactions) > maxTransactions {
		txes = m.transactions[:maxTransactions]
		m.transactions = m.transactions[maxTransactions:]
	} else {
		txes = m.transactions
		m.transactions = nil
	}
	m.Unlock()

	log.Println("Submitting batch with", len(txes), "transactions")

	err := m.globalInbox.DeliverTransactionBatchNoWait(
		ctx,
		m.rollupAddress,
		prepareTransactions(txes),
	)

	m.Lock()
	if err != nil {
		log.Println("Transaction aggregator failed: ", err)
		m.valid = false
	}
}

// SendTransaction takes a request signed transaction message from a client
// and puts it in a queue to be included in the next transaction batch
func (m *Server) SendTransaction(
	ctx context.Context,
	args *SendTransactionArgs,
) (*SendTransactionReply, error) {
	toBytes, err := hexutil.Decode(args.To)
	if err != nil {
		return nil, errors2.Wrap(err, "error decoding to")
	}
	var to common.Address
	copy(to[:], toBytes)

	sequenceNum, valid := new(big.Int).SetString(args.SequenceNum, 10)
	if !valid {
		return nil, errors.New("Invalid sequence num")
	}

	valueInt, valid := new(big.Int).SetString(args.Value, 10)
	if !valid {
		return nil, errors.New("Invalid value")
	}

	data, err := hexutil.Decode(args.Data)
	if err != nil {
		return nil, errors2.Wrap(err, "error decoding data")
	}

	pubkeyBytes, err := hexutil.Decode(args.Pubkey)
	if err != nil {
		return nil, errors2.Wrap(err, "error decoding pubkey")
	}

	signature, err := hexutil.Decode(args.Signature)
	if err != nil {
		return nil, errors2.Wrap(err, "error decoding signature")
	}

	if len(signature) != signatureLength {
		return nil, errors.New("signature of wrong length")
	}

	// Convert sig with normalized v
	if signature[recoverBitPos] == 27 {
		signature[recoverBitPos] = 0
	} else if signature[recoverBitPos] == 28 {
		signature[recoverBitPos] = 1
	}

	txDataHash := message.BatchTxHash(
		m.rollupAddress,
		to,
		sequenceNum,
		valueInt,
		data,
	)

	messageHash := hashing.SoliditySHA3WithPrefix(txDataHash[:])

	if !crypto.VerifySignature(
		pubkeyBytes,
		messageHash[:],
		signature[:len(signature)-1],
	) {
		return nil, errors.New("Invalid signature")
	}

	var sigData [signatureLength]byte
	copy(sigData[:], signature)

	tx := message.BatchTx{
		To:     to,
		SeqNum: sequenceNum,
		Value:  valueInt,
		Data:   data,
		Sig:    sigData,
	}

	txHash := message.BatchTxHash(
		m.rollupAddress,
		tx.To,
		tx.SeqNum,
		tx.Value,
		tx.Data,
	)

	var pubkey *ecdsa.PublicKey
	var pubkeyErr error
	if len(pubkeyBytes) == 33 {
		pubkey, pubkeyErr = crypto.DecompressPubkey(pubkeyBytes)
	} else {
		x, y := elliptic.Unmarshal(crypto.S256(), pubkeyBytes)
		pubkey = &ecdsa.PublicKey{Curve: crypto.S256(), X: x, Y: y}
	}

	if pubkeyErr != nil || pubkey.X == nil || pubkey.Y == nil {
		log.Println("Got tx: ", tx, "with hash", txHash, "from pubkey", hexutil.Encode(pubkeyBytes))
	} else {
		sender := crypto.PubkeyToAddress(*pubkey)
		log.Println("Got tx: ", tx, "with hash", txHash, "from", hexutil.Encode(sender[:]))
	}

	m.Lock()
	defer m.Unlock()

	if !m.valid {
		return nil, errors.New("Tx aggregator is not running")
	}

	m.transactions = append(m.transactions, DecodedBatchTx{
		tx:     tx,
		pubkey: pubkeyBytes,
	})

	return &SendTransactionReply{}, nil
}
