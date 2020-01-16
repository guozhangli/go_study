package TestProject

type QueueArray struct {
	Front int
	Rear  int
	Array []interface{}
}

func NewQueueArray(cap int) Queue {
	queueArray := &QueueArray{
		Front: 0,
		Rear:  0,
		Array: make([]interface{}, 0, cap),
	}
	return queueArray
}

func (queue *QueueArray) DeQueue() interface{} {
	panic("implement me")
}

func (queue *QueueArray) EnQueue(value interface{}) {
	panic("implement me")
}

func (queue *QueueArray) IsEmpty() bool {
	panic("implement me")
}

func (queue *QueueArray) Length() int {
	panic("implement me")
}

func (queue *QueueArray) Clear() {
	panic("implement me")
}

func (queue *QueueArray) GetFront() interface{} {
	panic("implement me")
}

func (queue *QueueArray) Distroy(q *Queue) {
	panic("implement me")
}
