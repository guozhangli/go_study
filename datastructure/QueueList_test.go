package TestProject

import "testing"

func initQueueList() Queue {
	queueList := NewQueueList(5)
	return queueList
}
func TestCreateQueueList(t *testing.T) {
	queueList := initQueueList()
	t.Log(queueList)
}

func TestQueueList_Length(t *testing.T) {
	queueList := initQueueList()
	queueList.EnQueue("1111")
	queueList.EnQueue("2222")
	queueList.EnQueue("3333")
	queueList.EnQueue("4444")
	queueList.EnQueue("5555")
	queueList.EnQueue("6666")
	len := queueList.Length()
	t.Log(len)
}

func TestQueueList_EnQueue(t *testing.T) {
	queueList := initQueueList()
	queueList.EnQueue("1111")
	queueList.EnQueue("2222")
	queueList.EnQueue("3333")
	queueList.EnQueue("4444")
	queueList.EnQueue("5555")
	queueList.EnQueue("6666")
	t.Log(queueList)
}

func TestQueueList_DeQueue(t *testing.T) {
	queueList := initQueueList()
	queueList.EnQueue("1111")
	queueList.EnQueue("2222")
	queueList.EnQueue("3333")
	queueList.EnQueue("4444")
	queueList.DeQueue()
	queueList.EnQueue("5555")
	queueList.DeQueue()
	t.Log(queueList)
}

func TestQueueList_GetFront(t *testing.T) {
	queueList := initQueueList()
	queueList.EnQueue("1111")
	queueList.EnQueue("2222")
	queueList.EnQueue("3333")
	queueList.EnQueue("4444")
	val := queueList.GetFront()
	t.Log(val)
}

func TestQueueList_IsEmpty(t *testing.T) {
	queueList := initQueueList()
	if queueList.IsEmpty() {
		t.Log("queueList is empty")
	} else {
		t.Log("queueList is not empty")
	}
}

func TestQueueList_Clear(t *testing.T) {
	queueList := initQueueList()
	queueList.EnQueue("1111")
	queueList.EnQueue("2222")
	queueList.EnQueue("3333")
	queueList.EnQueue("4444")
	queueList.Clear()
	t.Log(queueList)
}

func TestQueueList_Distroy(t *testing.T) {
	queueList := initQueueList()
	queueList.EnQueue("1111")
	queueList.EnQueue("2222")
	queueList.EnQueue("3333")
	queueList.EnQueue("4444")
	queueList.Clear()
	t.Log(queueList)
	queueList.Distroy(&queueList)
	t.Log(queueList)
}
