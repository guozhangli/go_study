package TestProject

import "testing"

func initQueueArray() Queue {
	queueArray := NewQueueArray(5)
	return queueArray
}
func TestCreateQueueArray(t *testing.T) {
	queueArray := initQueueArray()
	t.Log(queueArray)
}

func TestQueueArray_Length(t *testing.T) {
	queueArray := initQueueArray()
	len := queueArray.Length()
	t.Log(len)
}

func TestQueueArray_EnQueue(t *testing.T) {
	queueArray := initQueueArray()
	queueArray.EnQueue("1111")
	queueArray.EnQueue("2222")
	queueArray.EnQueue("3333")
	t.Log(queueArray)
}

func TestQueueArray_DeQueue(t *testing.T) {
	queueArray := initQueueArray()
	queueArray.EnQueue("1111")
	queueArray.EnQueue("2222")
	queueArray.EnQueue("3333")
	queueArray.DeQueue()
	t.Log(queueArray)
}
