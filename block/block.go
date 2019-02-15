package block

import (
	"bytes"
	"encoding/gob"
	"time"
)

// Block ... block
type Block struct {
	// when this block is creted
	Timestamp int64
	Data      []byte
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

// NewBlock ... Create new block
func NewBlock(data string, prevBlockHash []byte) *Block {
	b := &Block{
		Timestamp:     time.Now().Unix(),
		Data:          []byte(data),
		PrevBlockHash: prevBlockHash,
	}
	pow := NewProofOfWork(b)
	nonce, hash := pow.Run()
	b.Hash = hash[:]
	b.Nonce = nonce
	return b
}
