package dbclient

import (
	"fmt"

	"github.com/morimolymoly/blockchain/dbclient/dbinterface"
	"github.com/morimolymoly/blockchain/dbclient/leveldb"
)

// NewDatabase ... create Database client
func NewDatabase(dbtype string, path string) (dbinterface.Database, error) {
	if dbtype == "leveldb" {
		return leveldb.NewLevelDB(path)
	}
	return nil, fmt.Errorf("Undefined database type")
}
