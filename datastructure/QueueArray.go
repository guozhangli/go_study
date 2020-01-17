package TestProject

type QueueArray struct {
	Array []interface{}
}

func NewQueueArray(cap int) Queue {
	queueArray := &QueueArray{
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
	checkQueueArray(queue)
	var top = queue.Array[0]
	queue.Array = append(queue.Array[:0], queue.Array[1:]...)
	return top
}

func (queue *QueueArray) EnQueue(value interface{}) {
	checkQueueArray(queue)
	queue.Array = append(queue.Array, value)
}

func (queue *QueueArray) IsEmpty() bool {
	checkQueueArray(queue)
	if queue.Length() == 0 {
		return true
	}
	return false
}

func (queue *QueueArray) Length() int {
	checkQueueArray(queue)
	return len(queue.Array)
}

func (queue *QueueArray) Clear() {
	checkQueueArray(queue)
	queue.Array = append(queue.Array[0:0])
}

func (queue *QueueArray) GetFront() interface{} {
	checkQueueArray(queue)
	return queue.Array[0]
}

func (queue *QueueArray) Distroy(q *Queue) {
	queue.Array = nil
	*q = nil
}
