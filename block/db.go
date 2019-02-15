package block

import (
	"github.com/morimolymoly/blockchain/dbclient"
	"github.com/morimolymoly/blockchain/dbclient/dbinterface"
)

const databasePath = "uoooo.db"
const blockchainBucket = "blockchain2"
const dbtype = "leveldb"

type blockdb struct {
	db dbinterface.Database
}

func (b *blockdb) CheckBlockchainIsNull() (bool, error) {
	return b.db.ExistBucket(blockchainBucket)
}

func (b *blockdb) CreateBlockchainBucket() error {
	return b.db.CreateBucket(blockchainBucket)
}

func (b *blockdb) PutBlock(block *Block) error {
	bytes, err := block.Serialize()
	if err != nil {
		return err
	}
	err = b.db.Put(blockchainBucket, string(block.Hash), bytes)
	if err != nil {
		return err
	}
	return b.PutTip(block.Hash)
}

// GetBlock
func (b *blockdb) GetBlock(hash string) (*Block, error) {
	block, err := b.db.Get(blockchainBucket, hash)
	if err != nil {
		return nil, err
	}
	return DeserializeBlock(block)
}

func (b *blockdb) PutTip(hash []byte) error {
	return b.db.Put(blockchainBucket, "l", hash)
}

func (b *blockdb) GetTip() ([]byte, error) {
	return b.db.Get(blockchainBucket, "l")
}

func getNewBlockDB() (*blockdb, error) {
	db, err := dbclient.NewDatabase(dbtype, databasePath)
	if err != nil {
		return nil, err
	}
	b := &blockdb{
		db: db,
	}
	return b, nil
}
