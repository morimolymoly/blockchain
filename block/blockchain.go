package block

import "fmt"

// Blockchain ... blockchain
type Blockchain struct {
	tip []byte
	db  *blockdb
}

// AddBlock ... add block to blockchain
func (bc *Blockchain) AddBlock(data string) error {
	tip, err := bc.db.GetTip()
	if err != nil {
		return err
	}
	newBlock := NewBlock(data, tip)

	err = bc.db.PutBlock(newBlock)
	return err
}

// NewGenesisBlock ... generate FIRTST block in our blockchain
func NewGenesisBlock() *Block {
	fmt.Printf("Hmm... There is no block in blockchain!\nLet's create Genesis Block!\n")
	return NewBlock("Genesis Block", []byte{})
}

// NewBlockchain ... create blockchain with genesis block
func NewBlockchain() (*Blockchain, error) {
	db, err := getNewBlockDB()
	if err != nil {
		return nil, err
	}

	null, err := db.CheckBlockchainIsNull()
	if err != nil {
		return nil, err
	}

	var tip []byte
	if null {
		db.CreateBlockchainBucket()
		genesis := NewGenesisBlock()
		err := db.PutBlock(genesis)
		if err != nil {
			return nil, err
		}
		tip = genesis.Hash
	} else {
		h, err := db.GetTip()
		if err != nil {
			return nil, err
		}
		tip = h
	}
	return &Blockchain{
		tip: tip,
		db:  db,
	}, nil
}
