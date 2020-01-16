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

func checkQueueArray(queue *QueueArray) {
	if queue == nil {
		panic("no queueArray created")
	}
}

func (queue *QueueArray) DeQueue() interface{} {
	panic("implement me")
}

func (queue *QueueArray) EnQueue(value interface{}) {
	checkQueueArray(queue)
	queue.Array = append(queue.Array, value)
	var rear_index int
	queue.Rear = rear_index % cap(queue.Array)
}

func (queue *QueueArray) IsEmpty() bool {
	panic("implement me")
}

func (queue *QueueArray) Length() int {
	checkQueueArray(queue)
	if queue.Front == queue.Rear {
		return 0
	} else if queue.Rear > queue.Front {
		return queue.Rear - queue.Front
	} else {
		return cap(queue.Array) - queue.Front + queue.Rear
	}

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
