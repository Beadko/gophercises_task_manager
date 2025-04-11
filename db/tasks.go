package db

import (
	"encoding/binary"
	"time"

	"github.com/boltdb/bolt"
)

var (
	taskBucket = []byte("tasks")
	db         *bolt.DB
)

type Task struct {
	Key   int
	Value string
}

func Init(dbpath string) error {
	var err error
	db, err = bolt.Open(dbpath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(taskBucket)
		return err
	})
}

func CreateTask(t string) (int, error) {
	var id int
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		id64, _ := b.NextSequence()
		id = int(id64)
		return b.Put(itob(id), []byte(t))
	})
	if err != nil {
		return -1, nil
	}

	return id, nil
}

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}
