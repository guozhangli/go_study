package TestProject

import (
	"sync"
	"sync/atomic"
)

var (
	count    int32
	takeLock = new(sync.Mutex)
	putLock  = new(sync.Mutex)
	notEmpty = sync.NewCond(takeLock)
	notFull  = sync.NewCond(putLock)
)

type LinkedBlockingQueue struct {
	Head     *Node
	Last     *Node
	Capatity int32
}

func NewLinkedBlockingQueue(capacity int32) *LinkedBlockingQueue {
	return &LinkedBlockingQueue{
		Head:     nil,
		Last:     nil,
		Capatity: capacity,
	}
}

//从对头删除一个元素
func (lbq *LinkedBlockingQueue) dequeue() interface{} {
	if lbq.Head == nil {
		return nil
	}
	lbq.Head = lbq.Head.Next
	return lbq.Head.Data
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
func (lbq *LinkedBlockingQueue) take() interface{} {
	takeLock.Lock()
	for atomic.LoadInt32(&count) == 0 {
		notEmpty.Wait()
	}
	item := lbq.dequeue()
	c := atomic.AddInt32(&count, -1)
	if c > 1 {
		notEmpty.Signal()
	}
	takeLock.Unlock()
	if c < lbq.Capatity {
		notFull.Signal()
	}
	return item
}

//向队列中插入元素，如果队列已满，则阻塞线程，直到出现空间为止
func (lbq *LinkedBlockingQueue) put(value interface{}) {
	putLock.Lock()
	for atomic.LoadInt32(&count) == lbq.Capatity {
		notFull.Wait()
	}
	lbq.enqueue(value)
	c := atomic.AddInt32(&count, 1)
	if c < lbq.Capatity {
		notFull.Signal()
	}
	putLock.Unlock()
	if c > 0 {
		notEmpty.Signal()
	}
}

//向队列中插入元素，如果队列已满，则返回false，否则返回true（非阻塞）
func (lbq *LinkedBlockingQueue) offer(value interface{}) bool {
	putLock.Lock()
	if atomic.LoadInt32(&count) == lbq.Capatity {
		return false
	}
	lbq.enqueue(value)
	c := atomic.AddInt32(&count, 1)
	if c < lbq.Capatity {
		notFull.Signal()
	}
	putLock.Unlock()
	if c > 0 {
		notEmpty.Signal()
	}
	return true
}

//从队列中取出元素，如果队列为空，则返回null(非阻塞）
func (lbq *LinkedBlockingQueue) poll() interface{} {
	takeLock.Lock()
	if atomic.LoadInt32(&count) == 0 {
		return nil
	}
	item := lbq.dequeue()
	c := atomic.AddInt32(&count, -1)
	if c > 1 {
		notEmpty.Signal()
	}
	takeLock.Unlock()
	if c < lbq.Capatity {
		notFull.Signal()
	}
	return item
}

//从队列中返回元素，但不删除元素，如果队列为空，则返回null
func (lbq *LinkedBlockingQueue) peek() interface{} {
	return nil
}

func (lbq *LinkedBlockingQueue) size() int32 {
	return count
}

//清空队列元素
func (lbq *LinkedBlockingQueue) clear() {
	putLock.Lock()
	takeLock.Lock()
	for node := lbq.Head; node != nil; node = node.Next {
		node.Data = nil
	}
	lbq.Head = lbq.Last
	if atomic.LoadInt32(&count) == lbq.Capatity {
		notFull.Signal()
	}
	takeLock.Unlock()
	putLock.Unlock()
}
