package TestProject

import (
	"encoding/json"
	"testing"
)

func TestQueueLinked_Create(t *testing.T) {
	queueLinked := NewQueueLinked()
	t.Log(queueLinked)
}

func TestQueueLinked_EnQueue(t *testing.T) {
	queueLinked := NewQueueLinked()
	queueLinked.EnQueue("1111")
	queueLinked.EnQueue("2222")
	queueLinked.EnQueue("3333")
	json1, _ := json.Marshal(queueLinked)
	t.Log(string(json1))
}

func TestQueueLinked_DeQueue(t *testing.T) {
	queueLinked := NewQueueLinked()
	queueLinked.EnQueue("1111")
	queueLinked.EnQueue("2222")
	queueLinked.EnQueue("3333")
	val := queueLinked.DeQueue()
	t.Log(val)
	json1, _ := json.Marshal(queueLinked)
	t.Log(string(json1))
}

func TestQueueLinked_GetFront(t *testing.T) {
	queueLinked := NewQueueLinked()
	queueLinked.EnQueue("1111")
	queueLinked.EnQueue("2222")
	queueLinked.EnQueue("3333")
	val := queueLinked.GetFront()
	t.Log(val)
	json1, _ := json.Marshal(queueLinked)
	t.Log(string(json1))
}

func TestQueueLinked_Length(t *testing.T) {
	queueLinked := NewQueueLinked()
	queueLinked.EnQueue("1111")
	queueLinked.EnQueue("2222")
	queueLinked.EnQueue("3333")
	queueLinked.EnQueue("3333")
	len := queueLinked.Length()
	t.Log(len)
}

func TestQueueLinked_IsEmpty(t *testing.T) {
	queueLinked := NewQueueLinked()
	queueLinked.EnQueue("1111")
	if queueLinked.IsEmpty() {
		t.Log("queueLinked is empty")
	} else {
		t.Log("queueLinked is not empty")
	}
}

func TestQueueLinked_Clear(t *testing.T) {
	queueLinked := NewQueueLinked()
	queueLinked.EnQueue("1111")
	queueLinked.EnQueue("2222")
	queueLinked.EnQueue("3333")
	queueLinked.EnQueue("3333")
	queueLinked.Clear()
	t.Log(queueLinked)
}

func TestQueueLinked_Distroy(t *testing.T) {
	queueLinked := NewQueueLinked()
	queueLinked.EnQueue("1111")
	queueLinked.EnQueue("2222")
	queueLinked.EnQueue("3333")
	queueLinked.EnQueue("3333")
	queueLinked.Clear()
	t.Log(queueLinked)
	queueLinked.Distroy(&queueLinked)
	t.Log(queueLinked)
}
