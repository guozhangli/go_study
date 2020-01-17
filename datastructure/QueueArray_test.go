package TestProject

import (
	"testing"
)

func initQueueArray() Queue {
	queueArray := NewQueueArray()
	return queueArray
}

func TestCreateQueueArray(t *testing.T) {
	queueArray := initQueueArray()
	t.Log(queueArray)
}

func TestQueueArray_IsFull(t *testing.T) {
	queueArray := initQueueArray()
	if IsFull(queueArray.(*QueueArray)) {
		t.Log("queueArray is full")
	} else {
		t.Log("queueArray is not full")
	}
}

func TestQueueArray_Length(t *testing.T) {
	queueArray := initQueueArray()
	queueArray.EnQueue("1111")
	queueArray.EnQueue("2222")
	queueArray.EnQueue("3333")
	queueArray.EnQueue("4444")
	queueArray.EnQueue("5555")
	queueArray.EnQueue("6666")
	len := queueArray.Length()
	t.Log(len)
}

func TestQueueArray_EnQueue(t *testing.T) {
	queueArray := initQueueArray()
	queueArray.EnQueue("1111")
	queueArray.EnQueue("2222")
	queueArray.EnQueue("3333")
	queueArray.EnQueue("4444")
	queueArray.EnQueue("5555")
	queueArray.EnQueue("6666")
	queueArray.EnQueue("7777")
	queueArray.EnQueue("8888")
	queueArray.EnQueue("9999")

	t.Log(queueArray)
}

func TestQueueArray_DeQueue(t *testing.T) {
	queueArray := initQueueArray()
	queueArray.EnQueue("1111")
	queueArray.EnQueue("2222")
	queueArray.EnQueue("3333")
	queueArray.EnQueue("4444")
	queueArray.DeQueue()
	queueArray.EnQueue("5555")
	queueArray.DeQueue()
	t.Log(queueArray)
}

func TestQueueArray_GetFront(t *testing.T) {
	queueArray := initQueueArray()
	queueArray.EnQueue("1111")
	queueArray.EnQueue("2222")
	queueArray.EnQueue("3333")
	queueArray.EnQueue("4444")
	val := queueArray.GetFront()
	t.Log(queueArray)
	t.Log(val)
}

func TestQueueArray_IsEmpty(t *testing.T) {
	queueArray := initQueueArray()
	if queueArray.IsEmpty() {
		t.Log("queueArray is empty")
	} else {
		t.Log("queueArray is not empty")
	}
}

func TestQueueArray_Clear(t *testing.T) {
	queueArray := initQueueArray()
	queueArray.EnQueue("1111")
	queueArray.EnQueue("2222")
	queueArray.EnQueue("3333")
	queueArray.EnQueue("4444")
	queueArray.Clear()
	queueArray.EnQueue("3333")
	queueArray.EnQueue("4444")
	queueArray.EnQueue("3333")
	queueArray.EnQueue("4444")
	queueArray.EnQueue("3333")
	queueArray.EnQueue("4444")
	queueArray.EnQueue("3333")
	queueArray.EnQueue("4444")
	queueArray.EnQueue("3333")

	if IsFull(queueArray.(*QueueArray)) {
		t.Log("queueArray is full")
	} else {
		t.Log("queueArray is not full")
	}
	t.Log(queueArray)
}

func TestQueueArray_Distroy(t *testing.T) {
	queueArray := initQueueArray()
	queueArray.EnQueue("1111")
	queueArray.EnQueue("2222")
	queueArray.EnQueue("3333")
	queueArray.EnQueue("4444")
	queueArray.Clear()
	t.Log(queueArray)
	queueArray.Distroy(&queueArray)
	t.Log(queueArray)
}
