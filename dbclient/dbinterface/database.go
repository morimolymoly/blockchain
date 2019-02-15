package dbinterface

// Database ... database client
type Database interface {
	Put(bucket string, key string, item []byte) error
	Get(bucket string, key string) ([]byte, error)
	ExistBucket(bucket string) (bool, error)
	CreateBucket(bucket string) error
	Close() error
}
