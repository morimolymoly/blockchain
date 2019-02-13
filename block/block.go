package block

import (
	"bytes"
	"crypto/sha256"
	"strconv"
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
}

// SetHash ... setting hash
func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)

	b.Hash = hash[:]
}

// NewBlock ... Create new block
func NewBlock(data string, prevBlockHash []byte) *Block {
	b := &Block{
		Timestamp:     time.Now().Unix(),
		Data:          []byte(data),
		PrevBlockHash: prevBlockHash,
		Hash:          []byte{},
	}
	b.SetHash()
	return b
}
