package block

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"time"
)

// Block ... block
type Block struct {
	// when this block is creted
	Timestamp    int64
	Transactions []*Transaction
	// hash of the previsous block
	PrevBlockHash []byte
	// header
	Hash []byte
	// counter of pow
	Nonce int
}

// Serialize ... serialize block to bytes
func (b *Block) Serialize() ([]byte, error) {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(b)
	if err != nil {
		return nil, err
	}

	return result.Bytes(), nil
}

// DeserializeBlock ... deserialize to block
func DeserializeBlock(d []byte) (*Block, error) {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&block)
	if err != nil {
		return nil, err
	}

	return &block, nil
}

// HashTransactions ... get hash of transactions
func (b *Block) HashTransactions() []byte {
	var txHashes [][]byte
	var txHash [32]byte

	for _, tx := range b.Transactions {
		txHashes = append(txHashes, tx.ID)
	}
	txHash = sha256.Sum256(bytes.Join(txHashes, []byte{}))

	return txHash[:]
}

// NewBlock ... Create new block
func NewBlock(transactions []*Transaction, prevBlockHash []byte) *Block {
	b := &Block{
		Timestamp:     time.Now().Unix(),
		Transactions:  transactions,
		PrevBlockHash: prevBlockHash,
	}
	pow := NewProofOfWork(b)
	nonce, hash := pow.Run()
	b.Hash = hash[:]
	b.Nonce = nonce
	return b
}
