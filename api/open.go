package main

import (
	"fmt"
	"github.com/ZhengHe-MD/learn-bolt/bolt"
	"log"
)

func main() {
	db, err := bolt.Open("1.db", 0600, nil)
	fmt.Printf("%+v\n", bolt.DefaultOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}


