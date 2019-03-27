package leveldb

import (
	"time"

	"github.com/golang/leveldb"
	"github.com/golang/leveldb/db"
)

const (
	// degradationWarnInterval specifies how often warning should be printed if the
	// leveldb database cannot keep up with requested writes.
	degradationWarnInterval = time.Minute

	// minCache is the minimum amount of memory in megabytes to allocate to leveldb
	// read and write caching, split half and half.
	minCache = 16

	// minHandles is the minimum number of files handles to allocate to the open
	// database files.
	minHandles = 16

	// metricsGatheringInterval specifies the interval to retrieve leveldb database
	// compaction, io and pause stats to report to the user.
	metricsGatheringInterval = 3 * time.Second
)

type Database struct {
	fn string // Filename of db struct
	db *leveldb.DB
}

func New(file string, cache int, handles int, namespace string) (*Database, error) {
	if cache < minCache {
		cache = minCache
	}

	if handles < minHandles {
		handles = minHandles
	}

	dB, err := leveldb.Open(file, &db.Options{
		WriteBufferSize:      2048,
		BlockSize:            1024 * 10,
		BlockRestartInterval: 2,
	})

	if err != nil {
		return nil, err
	}

	ldb := &Database{
		fn: file,
		db: dB,
	}

	return ldb, nil

}

func (d *Database) Save(height string, i []byte) bool {
	d.db.Set([]byte(height), []byte(i), &db.WriteOptions{Sync: false})
	return true
}

func (d *Database) Read(key string) ([]byte, error) {
	r, err := d.db.Get([]byte(key), nil)
	return r, err
}

func (d *Database) Open() (*Database, error) {

	dB, err := leveldb.Open(d.fn, &db.Options{

		WriteBufferSize:      2048,
		BlockSize:            1024 * 10,
		BlockRestartInterval: 2})
	//	defer d.db.Close()
	ldb := &Database{
		fn: d.fn,
		db: dB,
	}

	return ldb, err
}

func (d *Database) SetName(name string) *Database {
	d.fn = name
	return d
}
