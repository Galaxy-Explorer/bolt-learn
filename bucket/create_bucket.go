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
		tx.CreateBucketIfNotExists([]byte("bucket1"))
		tx.CreateBucketIfNotExists([]byte("bucket2"))
		tx.CreateBucketIfNotExists([]byte("bucket3"))
		tx.CreateBucketIfNotExists([]byte("bucket4"))
		tx.CreateBucketIfNotExists([]byte("bucket5"))
		tx.CreateBucketIfNotExists([]byte("bucket6"))
		tx.CreateBucketIfNotExists([]byte("bucket7"))
		tx.CreateBucketIfNotExists([]byte("bucket8"))
		tx.CreateBucketIfNotExists([]byte("bucket9"))
		tx.CreateBucketIfNotExists([]byte("bucket10"))
		tx.CreateBucketIfNotExists([]byte("bucket11"))
		tx.CreateBucketIfNotExists([]byte("bucket12"))
		tx.CreateBucketIfNotExists([]byte("bucket13"))
		tx.CreateBucketIfNotExists([]byte("bucket14"))
		tx.CreateBucketIfNotExists([]byte("bucket15"))
		tx.CreateBucketIfNotExists([]byte("bucket16"))
		tx.CreateBucketIfNotExists([]byte("bucket17"))
		tx.CreateBucketIfNotExists([]byte("bucket18"))
		tx.CreateBucketIfNotExists([]byte("bucket20"))
		tx.CreateBucketIfNotExists([]byte("bucket21"))
		tx.CreateBucketIfNotExists([]byte("bucket22"))
		//_, err := tx.CreateBucketIfNotExists([]byte("bucket1"))
		//return err
		return nil
	})
}