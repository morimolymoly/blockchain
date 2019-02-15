package block

import (
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
