package TestProject

const capactity = 10

type QueueArray struct {
	Front int
	Rear  int
	Array [capactity]interface{}
}

func NewQueueArray() Queue {
	var arr [capactity]interface{}
	queueArray := &QueueArray{
		Front: 0,
		Rear:  0,
		Array: arr,
	}
	return queueArray
}

func checkQueueArray(queueArray *QueueArray) {
	if queueArray == nil {
		panic("no queueArray created")
	}
}

func IsFull(queue *QueueArray) bool {
	checkQueueArray(queue)
	return (queue.Rear+1)%capactity == queue.Front
}

func (queue *QueueArray) DeQueue() interface{} {
	checkQueueArray(queue)
	if queue.IsEmpty() {
		return nil
	}
	var top = queue.Array[queue.Front]
	queue.Array[queue.Front] = nil
	var front = queue.Front
	front++
	queue.Front = front % capactity
	return top
}

func (queue *QueueArray) EnQueue(value interface{}) {
	checkQueueArray(queue)
	if !IsFull(queue) {
		queue.Array[queue.Rear] = value
		var rear = queue.Rear
		rear++
		queue.Rear = rear % capactity
	} else {
		panic("queueArray is full")
	}
}

func (queue *QueueArray) IsEmpty() bool {
	checkQueueArray(queue)
	return queue.Front == queue.Rear
}

func (queue *QueueArray) Length() int {
	checkQueueArray(queue)
	return (queue.Rear - queue.Front + capactity) % capactity
}

func (queue *QueueArray) Clear() {
	checkQueueArray(queue)
	for !queue.IsEmpty() {
		queue.DeQueue()
	}
}

func (queue *QueueArray) GetFront() interface{} {
	checkQueueArray(queue)
	if queue.IsEmpty() {
		return nil
	}
	return queue.Array[queue.Front]
}

func (queue *QueueArray) Distroy(q *Queue) {
	checkQueueArray(queue)
	*q = nil
}
