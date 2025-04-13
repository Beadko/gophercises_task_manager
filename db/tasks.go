package db

import (
	"encoding/binary"
	"fmt"
	"time"

	"github.com/boltdb/bolt"
)

var (
	activeBucket    = []byte("active")
	completedBucket = []byte("completed")
	db              *bolt.DB
)

type Task struct {
	Key       int
	Value     string
	CreatedAt time.Time
}

func Init(dbpath string) error {
	var err error
	db, err = bolt.Open(dbpath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(activeBucket)
		return err
	})
}

func CreateTask(t string) (int, error) {
	var id int
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(activeBucket)
		id64, _ := b.NextSequence()
		id = int(id64)
		return b.Put(itob(id), []byte(t))
	})
	if err != nil {
		return -1, nil
	}

	return id, nil
}

func AllTasks() ([]Task, error) {
	var tasks []Task
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(activeBucket)
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			tasks = append(tasks, Task{
				Key:   btoi(k),
				Value: string(v),
			})
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func DeleteTask(key int) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(activeBucket)
		return b.Delete(itob(key))
	})
}

func DoTask(key int) error {
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(completedBucket)
		if err != nil {
			return nil
		}

		a := tx.Bucket(activeBucket)
		c := tx.Bucket(completedBucket)

		task := a.Get(itob(key))
		if task == nil {
			return fmt.Errorf("Error: task not found")
		}

		return c.Put(itob(key), []byte(task))
	})
}

func CompletedTasks() ([]Task, error) {
	var tasks []Task
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(completedBucket)
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			t, err := time.Parse(time.RFC3339, string(v))
			if err != nil {
				return err
			}
			if t.After(time.Now().Truncate(24 * time.Hour)) {
				tasks = append(tasks, Task{
					Key:   btoi(k),
					Value: string(v),
				})
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}
