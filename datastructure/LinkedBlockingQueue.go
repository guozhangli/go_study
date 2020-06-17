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
	return &LinkedBlockingQueue{
		Head:     nil,
		Last:     nil,
		Capatity: capacity,
		Count:    0,
	}
}

//从对头删除一个元素
func (lbq *LinkedBlockingQueue) dequeue() *Node {
	if lbq.Head == nil {
		return nil
	}
	var top = lbq.Head
	lbq.Head = lbq.Head.Next
	return top
}

//从队尾添加一个元素
func (lbq *LinkedBlockingQueue) enqueue(data interface{}) {
	node := NewNode(nil, data)
	if lbq.Last != nil {
		lbq.Last.Next = node
	}
	lbq.Last = node
	if lbq.Head == nil {
		lbq.Head = node
	}
}

//从队列中取出元素，如果队列为空，则阻塞当前线程，直到其中有元素为止
func (lbq *LinkedBlockingQueue) Take() interface{} {
	defer takeLock.Unlock()
	takeLock.Lock()
	for atomic.LoadInt32(&lbq.Count) == 0 {
		notEmpty.Wait()
	}
	item := lbq.dequeue()
	c := atomic.AddInt32(&lbq.Count, -1)
	if c > 1 {
		notEmpty.Signal()
	}
	if c < lbq.Capatity {
		notFull.Signal()
	}
	return item.Data
}

//向队列中插入元素，如果队列已满，则阻塞线程，直到出现空间为止
func (lbq *LinkedBlockingQueue) Put(value interface{}) {
	defer putLock.Unlock()
	putLock.Lock()
	for atomic.LoadInt32(&lbq.Count) == lbq.Capatity {
		notFull.Wait()
	}
	lbq.enqueue(value)
	c := atomic.AddInt32(&lbq.Count, 1)
	if c < lbq.Capatity {
		notFull.Signal()
	}
	if c > 0 {
		notEmpty.Signal()
	}
}

//向队列中插入元素，如果队列已满，则返回false，否则返回true（非阻塞）
func (lbq *LinkedBlockingQueue) Offer(value interface{}) bool {
	defer putLock.Unlock()
	putLock.Lock()
	if atomic.LoadInt32(&lbq.Count) == lbq.Capatity {
		return false
	}
	lbq.enqueue(value)
	c := atomic.AddInt32(&lbq.Count, 1)
	if c < lbq.Capatity {
		notFull.Signal()
	}
	if c > 0 {
		notEmpty.Signal()
	}
	return true
}

//从队列中取出元素，如果队列为空，则返回null(非阻塞）
func (lbq *LinkedBlockingQueue) Poll() interface{} {
	defer takeLock.Unlock()
	takeLock.Lock()
	if atomic.LoadInt32(&lbq.Count) == 0 {
		return nil
	}
	item := lbq.dequeue()
	c := atomic.AddInt32(&lbq.Count, -1)
	if c > 1 {
		notEmpty.Signal()
	}
	if c < lbq.Capatity {
		notFull.Signal()
	}
	return item.Data
}

//从队列中返回元素，但不删除元素，如果队列为空，则返回null
func (lbq *LinkedBlockingQueue) peek() interface{} {
	return nil
}

func (lbq *LinkedBlockingQueue) size() int32 {
	return lbq.Count
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
