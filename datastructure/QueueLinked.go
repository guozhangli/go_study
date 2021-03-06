package TestProject

type QueueLinked struct {
	Front *Node
	Rear  *Node
}

func (queue *QueueLinked) IsExist(value interface{}) bool {
	panic("implement me")
}

func NewQueueLinked() Queue {
	queueLinked := new(QueueLinked)
	return queueLinked
}

func checkQueueLinked(queue *QueueLinked) {
	if queue == nil {
		panic("no queueLinked created")
	}
}

//从队列头部取出元素
func (queue *QueueLinked) DeQueue() interface{} {
	checkQueueLinked(queue)
	if queue.Front == nil {
		return nil
	}
	var top = queue.Front
	queue.Front = queue.Front.Next
	return top
}

//向队列尾部加入元素
func (queue *QueueLinked) EnQueue(value interface{}) {
	checkQueueLinked(queue)
	node := NewNode(nil, value)
	if queue.Rear != nil {
		queue.Rear.Next = node
	}
	queue.Rear = node
	if queue.Front == nil {
		queue.Front = node
	}

}

func (queue *QueueLinked) IsEmpty() bool {
	checkQueueLinked(queue)
	if queue.Front == nil {
		return true
	}
	return false
}

func (queue *QueueLinked) Length() int {
	checkQueueLinked(queue)
	if queue.Front == nil {
		return 0
	}
	var front = queue.Front
	var count int
	for front != nil {
		front = front.Next
		count++
	}
	return count
}

func (queue *QueueLinked) Clear() {
	checkQueueLinked(queue)
	for queue.Length() > 0 {
		queue.DeQueue()
	}
	queue.Rear = nil
}

//得到队首元素
func (queue *QueueLinked) GetFront() interface{} {
	checkQueueLinked(queue)
	if queue.Front == nil {
		return nil
	}
	return queue.Front.Data
}

func (queue *QueueLinked) Distroy(q *Queue) {
	checkQueueLinked(queue)
	*q = nil
}
