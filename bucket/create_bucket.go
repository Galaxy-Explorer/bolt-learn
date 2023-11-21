package main

import (
	"github.com/ZhengHe-MD/learn-bolt/bolt"
	"log"
)

func main() {
	db, err := bolt.Open("1.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_ = db.Update(func(tx *bolt.Tx) error {
		b1, _ := tx.CreateBucketIfNotExists([]byte("bucket1"))
		b11, _ := b1.CreateBucketIfNotExists([]byte("bucket11"))
		b11.Put([]byte("k11"), []byte("v11"))
		return err
	})
}