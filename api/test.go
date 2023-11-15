package main

import (
	"fmt"
	"sync"
	"unsafe"
)

type Counter struct {
	mu    sync.Mutex
	count int
}
func (c *Counter) Add(n int) {
	c.count += n
}

func (c *Counter) Value() int {
	return c.count
}

const bucketHeaderSize = int(unsafe.Sizeof(bucket{}))

type bucket struct {
	root     pgid   // page id of the bucket's root-level page
	sequence uint64 // monotonically incrementing, used by NextSequence()
}

type leafPageElement struct {
	flags uint32
	pos   uint32
	ksize uint32
	vsize uint32
}

type pgid uint64

func main() {

	//pool := sync.Pool{
	//	New: func() interface{} {
	//		return &Counter{}
	//	},
	//}
	//counter := &Counter{}
	//
	//for i := 0; i < 10; i++ {
	//	go func(pool sync.Pool, i int) {
	//		counter = pool.Get().(*Counter)
	//		fmt.Printf("i = %d, before %d ", i, counter.Value())
	//		counter.Add(i)
	//		fmt.Printf("after %d\n", counter.Value())
	//		pool.Put(counter)
	//	}(pool, i)
	//}
	//
	//
	//time.Sleep(1 * time.Second)
	//fmt.Println(counter.Value())

	//const bucketHeaderSize = int(unsafe.Sizeof(bucket{}))
	//const leafPageElementSize = int(unsafe.Sizeof(leafPageElement{}))
	//fmt.Println(bucketHeaderSize, leafPageElementSize)
	//var value = make([]byte, 100)
	//fmt.Println(value)
	fmt.Println(4 & 1)

}