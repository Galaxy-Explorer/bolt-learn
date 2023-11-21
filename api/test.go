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

type page struct {
	id       pgid      // 物理上page的id
	flags    uint16    // 页的类型，（1，分支页类型，2，叶子页类型，4，元数据页，16，空闲链表页泪习惯）
	count    uint16    // 该页上有多少个元素，以叶子节点为例，只要计算leafHeader的个数就可以
	overflow uint32    // 数据是否在该页已经装不下了，这个时候就需要多个页来承载这些个节点数据，TODO:个人觉得overflow是记录了溢出了多少个页，这些页的pgid是连续的，所以很容易找到溢出的页；
	ptr      uintptr   // 底层的leafElement，branchElement是一个列表，ptr指向的就是这个列表的指针，这样就能够找到数组的起始位置，来根据位移定位数据
}

const bucketHeaderSize = int(unsafe.Sizeof(bucket{}))
const pageHeaderSize = int(unsafe.Offsetof(((*page)(nil)).ptr))


type leafPageElement struct {
	flags uint32
	pos   uint32
	ksize uint32
	vsize uint32
}

type pgid uint64

type Cursor struct {
	bucket *Bucket
}
type Bucket struct {
	*bucket
	name string
}

type bucket struct {
	root     pgid   // page id of the bucket's root-level page
	sequence uint64 // monotonically incrementing, used by NextSequence()
}


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
	t := &Cursor{ &Bucket{}}
	t.bucket.name = "test"
	fmt.Printf("%+v\n", t.bucket)
}