hmap结构体定义位置 runtime/map.go:hmap
###1、结构
```
type hmap struct {
	count     int // map的大小，即len()的值，代指map中的键值对个数
	
	flags     uint8 // 状态标识，主要是 goroutine 写入和扩容机制的相关状态控制，并发读写的判断条件之一
	
	B         uint8  // 桶，最大可容纳的元素数量，值为负载因子（默认 6.5）*2^B，是2的指数
	
	noverflow uint16 // approximate number of overflow buckets; see incrnoverflow for details
	
	hash0     uint32 // hash seed

	buckets    unsafe.Pointer // array of 2^B Buckets. may be nil if count==0.
	
	oldbuckets unsafe.Pointer // previous bucket array of half the size, non-nil only when growing
	
	nevacuate  uintptr        // progress counter for evacuation (buckets less than this have been evacuated)

	extra *mapextra // optional fields
}

type mapextra struct {
	// If both key and value do not contain pointers and are inline, then we mark bucket
	// type as containing no pointers. This avoids scanning such maps.
	// However, bmap.overflow is a pointer. In order to keep overflow buckets
	// alive, we store pointers to all overflow buckets in hmap.extra.overflow and hmap.extra.oldoverflow.
	// overflow and oldoverflow are only used if key and value do not contain pointers.
	// overflow contains overflow buckets for hmap.buckets.
	// oldoverflow contains overflow buckets for hmap.oldbuckets.
	// The indirection allows to store a pointer to the slice in hiter.
	overflow    *[]*bmap
	oldoverflow *[]*bmap

	// nextOverflow holds a pointer to a free overflow bucket.
	nextOverflow *bmap
}
```