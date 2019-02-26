hchan结构体定义位置 runtime/chan.go:hchan
###1、结构
```
type hchan struct {
    qcount uint // total data in the queue 队列中已有的元素个数
    
    dataqsiz uint // size of the circular queue 缓存队列大小(环形队列),ch := make(chan int, 10) => 就是这里这个 10
     
    buf unsafe.Pointer // points to an array of dataqsiz elements 指向channel缓存队列的指针
    
    elemsize uint16 // 通过channel传递的元素大小
    
    closed uint32 // 是否已被关闭,关闭时 closed==1
    
    elemtype *_type // element type 通过channel传递的元素类型(runtime._type 是 runtime包中的结构体)
    
    sendx uint // send index channel 中发送元素在队列中的索引,当sendx==dataqsiz时,sendx清零,回到对列头
    
    recvx uint // receive index channel 中接受元素在队列中的索引,当recvx==dataqsiz时,recvx清零,回到队列头
    
    recvq waitq // list of recv waiters 等待从channel中接收元素的协程列表
    
    sendq waitq // list of send waiters 等待向channel中发送元素的协程列表

    // lock protects all fields in hchan, as well as several fields in sudogs blocked on this channel.
    // 锁使得结构体和阻塞在该channel上的一些sudog中的字段是并发安全的
    
    // Do not change another G's status while holding this lock(in particular, do not ready a G), as this can deadlock with stack shrinking.
    // 当协程持有锁的时候不要修改其他协程的状态(尤其是不要将其他协程的状态变成ready),因为这可能导致堆栈减少而死锁
    
    // mutex是互斥锁,没有竞争的情况跟spin lock(自旋锁)一样快(只是几条用户级别的命令),发生竞争的时候会 sleep in kernel
    lock mutex
}
```
###2、初始化
```
func makechan(t *chantype, size int) *hchan {
    elem := t.elem

    // compiler checks this but be safe.
    if elem.size >= 1<<16 {
        throw("makechan: invalid channel element type")
    }
    if hchanSize%maxAlign != 0 || elem.align > maxAlign {
        throw("makechan: bad alignment")
    }
    
    // 校验size,必须大于等于0且缓存大小要小于heap中能分配的大小(_MaxMem是可分配的堆大小)
    if size < 0 || uintptr(size) > maxSliceCap(elem.size) || uintptr(size)*elem.size > _MaxMem-hchanSize {
        panic(plainError("makechan: size out of range"))
    }

    // Hchan does not contain pointers interesting for GC when elements stored in buf do not contain pointers.
    // 如果 buf 中的元素不包含有指针,那么hchan中也不包含指针,此时hchan与GC无关
    
    // buf points into the same allocation, elemtype is persistent.
    // SudoG's are referenced from their owning thread so they can't be collected.
    // SudoG无法被GC
    
    var c *hchan
    switch {
    case size == 0 || elem.size == 0: // 队列大小或元素大小为0
        c = (*hchan)(mallocgc(hchanSize, nil, true))
        // Race detector uses this location for synchronization.
        c.buf = unsafe.Pointer(c)
    case elem.kind&kindNoPointers != 0:
        // 非指针类型元素,直接分配(hchanSize+uintptr(size)*elem.size)大小的连续空间,小对象从P关联的缓存free list分配,大对象(>32kB)直接从heap中分配
        c = (*hchan)(mallocgc(hchanSize+uintptr(size)*elem.size, nil, true))
        c.buf = add(unsafe.Pointer(c), hchanSize)
    default: // 指针类型元素
        c = new(hchan)
        c.buf = mallocgc(uintptr(size)*elem.size, elem, true)
    }

    c.elemsize = uint16(elem.size)
    c.elemtype = elem
    c.dataqsiz = uint(size)

    return c
}
```
###3、发送
```
// 通过chansend1(),调用chansend(),其中block参数为true
func chansend(c *hchan, ep unsafe.Pointer, block bool, callerpc uintptr) bool {
	if c == nil {
		if !block {
			return false
		}
		// nil channel 发送数据会永远阻塞下去.(发生 panic 情况是 channel 被 closed 了,不是 nil channel)
		// gopark是执行挂起当前 goroutine 的操作(此处参数unlockf(用于唤醒g)为nil,即没有人来唤醒它,系统进入死锁.故channel必须被初始化之后才能使用,否则死锁)
		gopark(nil, nil, "chan send (nil chan)", traceEvGoStop, 2)
		throw("unreachable")
	}

	if debugChan {
		print("chansend: chan=", c, "\n")
	}

	if raceenabled { // 有竞争
		racereadpc(unsafe.Pointer(c), callerpc, funcPC(chansend))
	}

	// Fast path: check for failed non-blocking operation without acquiring the lock.
	// 在不获取锁的情况下检查失败的非阻塞操作

	// After observing that the channel is not closed, we observe that the channel is not ready for sending.
	// 在观察到channel未关闭时,同时发现channel是没有准备好发送的
	
	// Each of these observations is a single word-sized read (first c.closed and second c.recvq.first or c.qcount depending on kind of channel).
	// 每个观察值都是一个单字大小的读取(第一个c.closed和第二个c.recvq.first或c.qcount取决于channel类型)
	
	// Because a closed channel cannot transition from 'ready for sending' to 'not ready for sending', 
	// 因为关闭的channel无法从“准备发送”转换为“未准备发送”
	
	// even if the channel is closed between the two observations,
	// 即使两次观测之间的channel是已关闭的
	
	// they imply a moment between the two when the channel was both not yet closed and not ready for sending. 
	// 意味着两个channel之间存在一个瞬间,即既没有关闭,也没有准备好发送
	
	// We behave as if we observed the channel at that moment,and report that the send cannot proceed.
	// 我们的行为就好像我们当时观察到了channel,并报告发送无法继续.

	// It is okay if the reads are reordered here: if we observe that the channel is not ready for sending and then observe that it is not closed,
	// 如果在这里重新排序读取是可行的：如果我们观察到channel没有准备好发送,然后观察到它也没有关闭
	
	// that implies that the channel wasn't closed during the first observation.
	// 这意味着在第一次观察期间通道没有关闭
	
	if !block && c.closed == 0 && ((c.dataqsiz == 0 && c.recvq.first == nil) ||
		(c.dataqsiz > 0 && c.qcount == c.dataqsiz)) { // 非阻塞、未closed、(有cap且有接受者 或 有cap且cap=len即满了)
		return false
	}

	var t0 int64
	if blockprofilerate > 0 {
		t0 = cputicks()
	}
	
    // channel接收数据之前获取互斥锁
	lock(&c.lock)

	if c.closed != 0 { // channel 已被关闭,panic
		unlock(&c.lock)
		panic(plainError("send on closed channel"))
	}

    // 从接收队列中获取一个*sudog,sudog represents a goroutine in a wait list.也就是说发送到channel的数据会优先检查接收等待队列,如果有协程等待取数,就直接给它,发完解锁,操作完成
	if sg := c.recvq.dequeue(); sg != nil {// dequeue出队
		// Found a waiting receiver. We pass the value we want to send directly to the receiver, bypassing the channel buffer (if any).
		send(c, sg, ep, func() { unlock(&c.lock) }, 3) // send()方法会将数据写到从队列里取出来的sg中,通过goready()唤醒sg.g(即等待的协程),进行后续处理
		// c是empty的channel,ep被拷贝进了sg.elem中,然后唤醒goroutine,此时goroutine是ready状态可以被P调用,最后解锁c(sg必须从recvq删除).
		return true
	}

    
    // 没有接收协程在等待,则去检查channel的缓存队列是否还有空位,若有则将数据放到缓存队列中
	if c.qcount < c.dataqsiz { // len < cap 即还有位置,数据放进缓存
		// Space is available in the channel buffer. Enqueue the element to send.
		qp := chanbuf(c, c.sendx)
		if raceenabled {
			raceacquire(qp)
			racerelease(qp)
		}
		typedmemmove(c.elemtype, qp, ep) // 将ep指向的类型内存拷贝进qp地址空间
		c.sendx++ // 将发送 index +1
		if c.sendx == c.dataqsiz { // 环形队列,当发送index加到最大时置0
			c.sendx = 0
		}
		c.qcount++ // channel中元素数量增加
		unlock(&c.lock)
		return true
	}

	if !block {
		unlock(&c.lock)
		return false
	}
	
    // channel已满,或者无缓冲的channel
    
	// Block on the channel. Some receiver will complete our operation for us.
	// 阻塞住发送协程,等有合适机会再将数据发送出去
	gp := getg() // 获取当前协程对象g的指针
	mysg := acquireSudog() // 生成一个sudog
	mysg.releasetime = 0
	if t0 != 0 {
		mysg.releasetime = -1
	}
	// No stack splits between assigning elem and enqueuing mysg on gp.waiting where copystack can find it.
	mysg.elem = ep
	mysg.waitlink = nil
	mysg.g = gp
	mysg.isSelect = false
	mysg.c = c
	gp.waiting = mysg
	gp.param = nil
	c.sendq.enqueue(mysg) // 将当前这个发送 goroutine 打包后的 sudog 入队到 channel 的 sendq 队列中
	goparkunlock(&c.lock, "chan send", traceEvGoBlockSend, 3) // 将当前goroutine状态置成waiting状态,然后解锁c

	// someone woke us up.
	// 以下是被唤醒后要执行的,先判断当前是不是合法的休眠中
	if mysg != gp.waiting {
		throw("G waiting list is corrupted")
	}
	gp.waiting = nil
	if gp.param == nil {
		if c.closed == 0 {
			throw("chansend: spurious wakeup")
		}
		panic(plainError("send on closed channel"))
	}
	gp.param = nil
	if mysg.releasetime > 0 {
		blockevent(mysg.releasetime-t0, 2)
	}
	mysg.c = nil
	releaseSudog(mysg)
	return true
}

// send processes a send operation on an empty channel c.
// The value ep sent by the sender is copied to the receiver sg.
// The receiver is then woken up to go on its merry way.
// Channel c must be empty and locked.  send unlocks c with unlockf.
// sg must already be dequeued from c.
// ep must be non-nil and point to the heap or the caller's stack.
func send(c *hchan, sg *sudog, ep unsafe.Pointer, unlockf func(), skip int) {
    // receiver 的 sudog 已经在对应区域分配过空间,仅需把数据拷贝过去
	if raceenabled {
		if c.dataqsiz == 0 {
			racesync(c, sg)
		} else {
			// Pretend we go through the buffer, even though
			// we copy directly. Note that we need to increment
			// the head/tail locations only when raceenabled.
			qp := chanbuf(c, c.recvx)
			raceacquire(qp)
			racerelease(qp)
			raceacquireg(sg.g, qp)
			racereleaseg(sg.g, qp)
			c.recvx++
			if c.recvx == c.dataqsiz {
				c.recvx = 0
			}
			c.sendx = c.recvx // c.sendx = (c.sendx+1) % c.dataqsiz
		}
	}
	if sg.elem != nil {
		sendDirect(c.elemtype, sg, ep)
		sg.elem = nil
	}
	gp := sg.g
	unlockf()
	gp.param = unsafe.Pointer(sg)
	if sg.releasetime != 0 {
		sg.releasetime = cputicks()
	}
	goready(gp, skip+1)
}
```
###4、接收
```
// chanrecv receives on channel c and writes the received data to ep.
// 从c接收数据并将数据写入到ep

// ep may be nil, in which case received data is ignored.
// 如果ep是nil接收到的数据被忽略

// If block == false and no elements are available, returns (false, false).if block == false 
// 并且channel是empty,返回(false,false)

// Otherwise, if c is closed, zeros *ep and returns (true, false).
// 如果channel已经关闭,ep存的是类型零值,返回(true,false)

// Otherwise, fills in *ep with an element and returns (true, true).
// 否则ep存储元素地址,返回(true,true)

// A non-nil ep must point to the heap or the caller's stack.
// 非nil的ep指向堆或者调用者的帧栈.

// 通过chanrecv1(),调用chanrecv(),其中block参数为true
func chanrecv(c *hchan, ep unsafe.Pointer, block bool) (selected, received bool) {
	// raceenabled: don't need to check ep, as it is always on the stack or is new memory allocated by reflect.

	if debugChan {
		print("chanrecv: chan=", c, "\n")
	}
	
    
	if c == nil { // 如果在 nil channel 上进行 recv 操作也会永远阻塞
		if !block { // 非阻塞的情况下要直接返回,非阻塞出现在一些 select 的场景中
			return
		}
		gopark(nil, nil, "chan receive (nil chan)", traceEvGoStop, 2) // 当前 goroutine: Grunning -> Gwaiting
		throw("unreachable")
	}

	// Fast path: check for failed non-blocking operation without acquiring the lock.

	// After observing that the channel is not ready for receiving, we observe that the
	// channel is not closed. Each of these observations is a single word-sized read
	// (first c.sendq.first or c.qcount, and second c.closed).
	// Because a channel cannot be reopened, the later observation of the channel
	// being not closed implies that it was also not closed at the moment of the
	// first observation. We behave as if we observed the channel at that moment
	// and report that the receive cannot proceed.

	// The order of operations is important here: reversing the operations can lead to incorrect behavior when racing with a close.
	
	
	if !block && (c.dataqsiz == 0 && c.sendq.first == nil || c.dataqsiz > 0 && atomic.Loaduint(&c.qcount) == 0) &&
		atomic.Load(&c.closed) == 0 { // 非阻塞且没内容可收的情况下直接返回,此时两个返回值就是 false,false
		return
	}

	var t0 int64
	if blockprofilerate > 0 {
		t0 = cputicks()
	}

	lock(&c.lock)

	if c.closed != 0 && c.qcount == 0 { // channel未关闭并且channel没有元素,直接结束
		if raceenabled {
			raceacquire(unsafe.Pointer(c))
		}
		unlock(&c.lock)
		if ep != nil {
			typedmemclr(c.elemtype, ep) // 清空ep指向的内存
		}
		return true, false
	}

	if sg := c.sendq.dequeue(); sg != nil {
		// Found a waiting sender. If buffer is size 0, receive value directly from sender.
		// 找到一个等待协程,如果是无缓冲的channel,直接将sudog.elem拷贝进ep
		
		// Otherwise, receive from head of queue and add sender's value to the tail of the queue.
		// 否则通过计算recvx得到qp,将ep拷贝进qp地址空间
		
		// (both map to the same buffer slot because the queue is full).
		// 然后唤醒goroutine,此时goroutine是ready状态可以被P调用,最后解锁c,sg必须从sendq删除
		
		recv(c, sg, ep, func() {unlock(&c.lock) }, 3)
		return true, true
	}

    //如果channel中有元素,直接从内存数组qp中移动到ep中,之后将qp指向的内存清空,recvx++
	if c.qcount > 0 {
		// Receive directly from queue
		// 根据recv index获取内存地址
		qp := chanbuf(c, c.recvx)
		if raceenabled {
			raceacquire(qp)
			racerelease(qp)
		}
		if ep != nil {
			typedmemmove(c.elemtype, ep, qp) // 将qp指向的内存拷贝到ep
		}
		typedmemclr(c.elemtype, qp) // 清空qp指向的内存
		c.recvx++ // 接收索引 +1
		if c.recvx == c.dataqsiz { // 若队列已空
			c.recvx = 0
		}
		c.qcount-- // buffer 元素计数 -1
		unlock(&c.lock)
		return true, true
	}

	if !block {
		unlock(&c.lock)
		return false, false
	}

	// no sender available: block on this channel.阻塞在<-ch操作上,使用当前waiting状态的goroutine,ep(元素地址) 构造*sudog
	gp := getg()
	mysg := acquireSudog()
	mysg.releasetime = 0
	if t0 != 0 {
		mysg.releasetime = -1
	}
	// No stack splits between assigning elem and enqueuing mysg
	// on gp.waiting where copystack can find it.
	mysg.elem = ep
	mysg.waitlink = nil
	gp.waiting = mysg
	mysg.g = gp
	mysg.isSelect = false
	mysg.c = c
	gp.param = nil
	c.recvq.enqueue(mysg) // *sudog进入receiver队列
	goparkunlock(&c.lock, "chan receive", traceEvGoBlockRecv, 3) // 将当前goroutine状态置成waiting状态,然后解锁c,goroutine通过调用goready(gp) 变成runnable

	// someone woke us up
	if mysg != gp.waiting {
		throw("G waiting list is corrupted")
	}
	gp.waiting = nil
	if mysg.releasetime > 0 {
		blockevent(mysg.releasetime-t0, 2)
	}
	closed := gp.param == nil
	gp.param = nil
	mysg.c = nil
	releaseSudog(mysg)
	return true, !closed
}

// recv processes a receive operation on a full channel c.
// There are 2 parts:
// 1) The value sent by the sender sg is put into the channel
//    and the sender is woken up to go on its merry way.
// 2) The value received by the receiver (the current G) is
//    written to ep.
// For synchronous channels, both values are the same.
// For asynchronous channels, the receiver gets its data from
// the channel buffer and the sender's data is put in the
// channel buffer.
// Channel c must be full and locked. recv unlocks c with unlockf.
// sg must already be dequeued from c.
// A non-nil ep must point to the heap or the caller's stack.
func recv(c *hchan, sg *sudog, ep unsafe.Pointer, unlockf func(), skip int) {
	if c.dataqsiz == 0 { // 当channel为无缓存channel时,直接将发送协程中的数据,拷贝给接收者
		if raceenabled {
			racesync(c, sg)
		}
		if ep != nil {
			// copy data from sender
			recvDirect(c.elemtype, sg, ep)
		}
	} else { // channel有缓存,根据缓存的接收游标,从缓存队列中取出一个拷贝给接受者;将发送协程中的数据,放到空出来的缓存位置中,游标下移;channel接收操作解锁;唤醒取出的发送协程;阻塞接收协程
		// Queue is full. Take the item at the
		// head of the queue. Make the sender enqueue
		// its item at the tail of the queue. Since the
		// queue is full, those are both the same slot.
		qp := chanbuf(c, c.recvx)
		if raceenabled {
			raceacquire(qp)
			racerelease(qp)
			raceacquireg(sg.g, qp)
			racereleaseg(sg.g, qp)
		}
		// copy data from queue to receiver
		if ep != nil {
			typedmemmove(c.elemtype, ep, qp)
		}
		// copy data from sender to queue
		typedmemmove(c.elemtype, qp, sg.elem)
		c.recvx++
		if c.recvx == c.dataqsiz {
			c.recvx = 0
		}
		c.sendx = c.recvx // c.sendx = (c.sendx+1) % c.dataqsiz
	}
	sg.elem = nil
	gp := sg.g
	unlockf()
	gp.param = unsafe.Pointer(sg)
	if sg.releasetime != 0 {
		sg.releasetime = cputicks()
	}
	goready(gp, skip+1)
}
```
###5、关闭
```
func closechan(c *hchan) {
	if c == nil {
		panic(plainError("close of nil channel"))
	}

	lock(&c.lock)
	if c.closed != 0 {
		unlock(&c.lock)
		panic(plainError("close of closed channel"))
	}

	if raceenabled {
		callerpc := getcallerpc()
		racewritepc(unsafe.Pointer(c), callerpc, funcPC(closechan))
		racerelease(unsafe.Pointer(c))
	}

	c.closed = 1

	var glist *g // MPG中的G
	
	// 以下顺序执行的两个for循环：
	// 1:清空recerver queue,将sg.elem指向的内存清空
	// 2:清空sender queue, sg.elem=nil,解锁channel,并将receiver queue和sender queue中获得的goroutine全部置成runnable状态,可以被P调度

	// release all readers
	for {
		sg := c.recvq.dequeue()
		if sg == nil { // 弹出的 sudog 是 nil说明读队列已经空了
			break
		}
		if sg.elem != nil {
			typedmemclr(c.elemtype, sg.elem) // 释放对应的内存
			sg.elem = nil
		}
		if sg.releasetime != 0 {
			sg.releasetime = cputicks()
		}
		gp := sg.g
		gp.param = nil
		if raceenabled {
			raceacquireg(gp, unsafe.Pointer(c))
		}
		gp.schedlink.set(glist)
		glist = gp
	}

	// release all writers (they will panic)
	// 将所有挂在 channel 上的 writer 从 sendq 中弹出,该操作会使所有 writer panic
	for {
		sg := c.sendq.dequeue()
		if sg == nil {
			break
		}
		sg.elem = nil
		if sg.releasetime != 0 {
			sg.releasetime = cputicks()
		}
		gp := sg.g
		gp.param = nil
		if raceenabled {
			raceacquireg(gp, unsafe.Pointer(c))
		}
		gp.schedlink.set(glist)
		glist = gp
	}
	unlock(&c.lock)

	// Ready all Gs now that we've dropped the channel lock.
	for glist != nil {
		gp := glist
		glist = glist.schedlink.ptr()
		gp.schedlink = 0
		goready(gp, 3)
	}
}
```
###n、其余
```
// 用于获取buf指向的数组的第i个元素
func chanbuf(c *hchan, i uint) unsafe.Pointer {
	return add(c.buf, uintptr(i)*uintptr(c.elemsize))
}
```