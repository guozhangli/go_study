package TestProject

var count int32

type LinkedBlockingQueue struct {
	front    *Node
	rear     *Node
	capatity int32
}

func NewLInkedBlockingQueue(capacity int32) *LinkedBlockingQueue {
	return &LinkedBlockingQueue{
		front:    nil,
		rear:     nil,
		capatity: capacity,
	}
}

//从队列中取出元素，如果队列为空，则阻塞当前线程，直到其中有元素为止
func (lbq *LinkedBlockingQueue) take() interface{} {
	return nil
}

//向队列中插入元素，如果队列已满，则阻塞线程，直到出现空间为止
func (lbq *LinkedBlockingQueue) put(value interface{}) {

}

//向队列中插入元素，如果队列已满，则返回false，否则返回true（非阻塞）
func (lbq *LinkedBlockingQueue) offer(value interface{}) bool {

}

//从队列中取出元素，如果队列为空，则返回null
func (lbq *LinkedBlockingQueue) poll() interface{} {
	return nil
}

//从队列中返回元素，但不删除元素，如果队列为空，则返回null
func (lbq *LinkedBlockingQueue) peek() interface{} {
	return nil
}

func (lbq *LinkedBlockingQueue) size() int32 {
	return count
}
