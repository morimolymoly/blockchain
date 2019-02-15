package block

import "reflect"

// BlockchainIterator ... iterator
type BlockchainIterator struct {
	currentHash []byte
	db          *blockdb
	Final       bool
}

// Iterator ... get iterator
func (bc *Blockchain) Iterator() *BlockchainIterator {
	bci := &BlockchainIterator{
		currentHash: bc.tip,
		db:          bc.db,
		Final:       false,
	}
	return bci
}

// Next ... get previous block
func (i *BlockchainIterator) Next() (*Block, error) {
	block, err := i.db.GetBlock(string(i.currentHash))
	if err != nil {
		return nil, err
	}
	i.currentHash = block.PrevBlockHash
	i.Final = reflect.DeepEqual(i.currentHash, []byte{})
	return block, nil
}
