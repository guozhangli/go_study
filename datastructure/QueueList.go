package TestProject

import (
	"fmt"
	"reflect"
)

type QueueList struct {
	Array []interface{}
}

func NewQueueList() Queue {
	queueList := &QueueList{
		Array: make([]interface{}, 0),
	}
	return queueList
}

func checkQueueList(queue *QueueList) {
	if queue == nil {
		panic("no queueList created")
	}
}

func (queue *QueueList) DeQueue() interface{} {
	checkQueueList(queue)
	var top = queue.Array[0]
	queue.Array = append(queue.Array[:0], queue.Array[1:]...)
	return top
}

func (queue *QueueList) EnQueue(value interface{}) {
	checkQueueList(queue)
	queue.Array = append(queue.Array, value)
}

func (queue *QueueList) IsEmpty() bool {
	checkQueueList(queue)
	if queue.Length() == 0 {
		return true
	}
	return false
}

func (queue *QueueList) Length() int {
	checkQueueList(queue)
	return len(queue.Array)
}

func (queue *QueueList) Clear() {
	checkQueueList(queue)
	queue.Array = append(queue.Array[0:0])
}

func (queue *QueueList) GetFront() interface{} {
	checkQueueList(queue)
	return queue.Array[0]
}

func (queue *QueueList) Distroy(q *Queue) {
	queue.Array = nil
	*q = nil
}

func (queue *QueueList) IsExist(value interface{}) bool {
	checkQueueList(queue)
	for _, v := range queue.Array {
		if reflect.DeepEqual(v, value) {
			return true
		}
	}
	return false
}

func (queue *QueueList) Print() {
	for _, v := range queue.Array {
		fmt.Println(v)
	}
}
