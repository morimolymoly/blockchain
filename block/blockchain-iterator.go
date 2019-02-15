package block

import "reflect"

// BlockchainIterator ... iterator
type BlockchainIterator struct {
	currentHash []byte
	db          *blockdb
}

// Iterator ... get iterator
func (bc *Blockchain) Iterator() *BlockchainIterator {
	bci := &BlockchainIterator{bc.tip, bc.db}
	return bci
}

// Next ... get previous block
func (i *BlockchainIterator) Next() (*Block, bool, error) {
	block, err := i.db.GetBlock(string(i.currentHash))
	if err != nil {
		return nil, false, err
	}
	i.currentHash = block.PrevBlockHash
	final := reflect.DeepEqual(i.currentHash, []byte{})
	return block, final, nil
}
