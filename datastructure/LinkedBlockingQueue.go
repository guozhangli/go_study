package TestProject

import (
	"sync"
	"sync/atomic"
)

var (
	takeLock = new(sync.Mutex)
	putLock  = new(sync.Mutex)
	notEmpty = sync.NewCond(takeLock)
	notFull  = sync.NewCond(putLock)
)

type LinkedBlockingQueue struct {
	Head     *Node
	Last     *Node
	Capatity int32
	Count    int32
}

func NewLinkedBlockingQueue(capacity int32) *LinkedBlockingQueue {
	node := NewNode(nil, nil)
	lbq := &LinkedBlockingQueue{
		Capatity: capacity,
		Count:    0,
	}
	lbq.Head = node
	lbq.Last = lbq.Head
	return lbq
}

//从对头删除一个元素
func (lbq *LinkedBlockingQueue) dequeue() interface{} {
	var h = lbq.Head
	var top = h.Next
	h.Next = h //help GC
	lbq.Head = top
	var item = top.Data
	top.Data = nil
	return item
}

//从队尾添加一个元素
func (lbq *LinkedBlockingQueue) enqueue(data interface{}) {
	node := NewNode(nil, data)
	lbq.Last.Next = node
	lbq.Last = lbq.Last.Next
}

//向队列中插入元素，如果队列已满，则阻塞线程，直到出现空间为止
func (lbq *LinkedBlockingQueue) Put(value interface{}) {
	putLock.Lock()
	c := int32(-1)
	for atomic.LoadInt32(&lbq.Count) == lbq.Capatity {
		notFull.Wait()
	}
	lbq.enqueue(value)
	c = atomic.AddInt32(&lbq.Count, 1)
	if c < lbq.Capatity {
		notFull.Signal()
	}
	putLock.Unlock()
	if c == 1 {
		signalNotEmpty()
	}
}

//向队列中插入元素，如果队列已满，则返回false，否则返回true（非阻塞）
func (lbq *LinkedBlockingQueue) Offer(value interface{}) bool {
	if atomic.LoadInt32(&lbq.Count) == lbq.Capatity {
		return false
	}
	c := int32(-1)
	putLock.Lock()
	if atomic.LoadInt32(&lbq.Count) < lbq.Capatity {
		lbq.enqueue(value)
		c = atomic.AddInt32(&lbq.Count, 1)
		if c < lbq.Capatity {
			notFull.Signal()
		}
	}
	putLock.Unlock()
	if c == 1 {
		signalNotEmpty()
	}
	return c >= 1
}

func signalNotEmpty() {
	defer takeLock.Unlock()
	takeLock.Lock()
	notEmpty.Signal()
}

//从队列中取出元素，如果队列为空，则阻塞当前线程，直到其中有元素为止
func (lbq *LinkedBlockingQueue) Take() interface{} {
	c := int32(-1)
	takeLock.Lock()
	for atomic.LoadInt32(&lbq.Count) == 0 {
		notEmpty.Wait()
	}
	item := lbq.dequeue()
	c = atomic.AddInt32(&lbq.Count, -1)
	if c > 0 {
		notEmpty.Signal()
	}
	takeLock.Unlock()
	if c == lbq.Capatity-1 {
		signalNotFull()
	}
	return item
}

//从队列中取出元素，如果队列为空，则返回null(非阻塞）
func (lbq *LinkedBlockingQueue) Poll() interface{} {
	if atomic.LoadInt32(&lbq.Count) == 0 {
		return nil
	}
	var item interface{}
	c := int32(-1)
	takeLock.Lock()
	if atomic.LoadInt32(&lbq.Count) > 0 {
		item = lbq.dequeue()
		c = atomic.AddInt32(&lbq.Count, -1)
		if c > 0 {
			notEmpty.Signal()
		}
	}
	takeLock.Unlock()
	if c == lbq.Capatity-1 {
		signalNotFull()
	}
	return item
}

func signalNotFull() {
	defer putLock.Unlock()
	putLock.Lock()
	notFull.Signal()
}

//从队列中返回元素，但不删除元素，如果队列为空，则返回null
func (lbq *LinkedBlockingQueue) peek() interface{} {
	return nil
}

func (lbq *LinkedBlockingQueue) size() int32 {
	return atomic.LoadInt32(&lbq.Count)
}

//清空队列元素
func (lbq *LinkedBlockingQueue) clear() {
	putLock.Lock()
	takeLock.Lock()
	for node := lbq.Head; node != nil; node = node.Next {
		node.Data = nil
	}
	lbq.Head = lbq.Last
	if atomic.LoadInt32(&lbq.Count) == lbq.Capatity {
		notFull.Signal()
	}
	takeLock.Unlock()
	putLock.Unlock()
}
