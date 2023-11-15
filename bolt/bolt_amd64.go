package bolt

// maxMapSize represents the largest mmap size supported by Bolt.
const maxMapSize = 0xFFFFFFFFFFFF // 256TB

// maxAllocSize is the size used when creating array pointers.
// maxAllocSize 是指分配一个足够大的数组，来暂时存放(k, v)，如果是切片，那切片在扩容的时候，地址可能会发生变化
const maxAllocSize = 0x7FFFFFFF

// Are unaligned load/stores broken on this arch?
var brokenUnaligned = false
