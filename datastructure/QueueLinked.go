package TestProject

type QueueLinked struct {
	Front int
	Rear  int
	node  *Node
}

func NewQueueLinked() Queue {
	queueLinked := new(QueueLinked)
	return queueLinked
}

func (queue *QueueLinked) DeQueue() interface{} {
	panic("implement me")
}

func (queue *QueueLinked) EnQueue(value interface{}) {
	panic("implement me")
}

func (queue *QueueLinked) IsEmpty() bool {
	panic("implement me")
}

func (queue *QueueLinked) Length() int {
	panic("implement me")
}

func (queue *QueueLinked) Clear() {
	panic("implement me")
}

func (queue *QueueLinked) GetFront() interface{} {
	panic("implement me")
}

func (queue *QueueLinked) Distroy(q *Queue) {
	panic("implement me")
}