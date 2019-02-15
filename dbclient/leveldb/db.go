package leveldb

import (
	"fmt"

	"github.com/morimolymoly/blockchain/dbclient/dbinterface"
	l "github.com/syndtr/goleveldb/leveldb"
)

func getPlace(bucket string, key string) []byte {
	place := []byte(fmt.Sprintf("%s/%s", bucket, key))
	return place
}

// LevelDB ... client for leveldb
type LevelDB struct {
	db *l.DB
}

// ExistBucket ... return true if blockchain is null
func (d LevelDB) ExistBucket(bucket string) (bool, error) {
	data, err := d.db.Get([]byte(bucket), nil)

	if err.Error() == l.ErrNotFound.Error() {
		return true, nil
	}

	if err != nil {
		return true, err
	}
	return (data == nil), nil
}

// CreateBucket ... create bucket
func (d LevelDB) CreateBucket(bucket string) error {
	err := d.db.Put([]byte(bucket), []byte{}, nil)
	if err != nil {
		return err
	}
	return nil
}

// Put ... put data to db
func (d LevelDB) Put(bucket string, key string, item []byte) error {
	place := getPlace(bucket, key)
	err := d.db.Put(place, item, nil)
	if err != nil {
		return err
	}
	return nil
}

// Get ... get data from db
func (d LevelDB) Get(bucket string, key string) ([]byte, error) {
	place := getPlace(bucket, key)
	data, err := d.db.Get(place, nil)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Close ... close db
func (d LevelDB) Close() error {
	err := d.db.Close()
	return err
}

// NewLevelDB ... create database client
func NewLevelDB(path string) (dbinterface.Database, error) {
	db, err := l.OpenFile(path, nil)
	if err != nil {
		return nil, err
	}
	ret := LevelDB{
		db: db,
	}
	return ret, nil
}
