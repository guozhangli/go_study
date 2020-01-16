package TestProject

import (
	"testing"
)

func initLinkedList() *LinkedList {
	linkedList := new(LinkedList)
	return linkedList
}
func TestCreateNode(t *testing.T) {
	linkedList := initLinkedList()
	if linkedList != nil {
		t.Log(linkedList)
	}
}

func TestLinkedListAdd(t *testing.T) {
	linkedList := initLinkedList1()
	for linkedList != nil {
		t.Log(linkedList)
		linkedList = linkedList.nextNode
	}
}

func TestLinkedListHeadAdd(t *testing.T) {
	linkedList := initLinkedList()
	linkedList.HeadAdd("1111")
	linkedList.HeadAdd("2222")
	linkedList.HeadAdd("3333")
	linkedList.HeadAdd("4444")
	for linkedList != nil {
		t.Log(linkedList)
		linkedList = linkedList.nextNode
	}
}

func initLinkedList1() *LinkedList {
	linkedList := initLinkedList()
	linkedList.TailAdd("1111")
	linkedList.TailAdd("2222")
	linkedList.TailAdd("3333")
	linkedList.TailAdd("4444")
	return linkedList
}

func TestLinkedListGetVaule(t *testing.T) {
	linkedList := initLinkedList1()
	for i := 0; i < 8; i++ {
		data := linkedList.GetVaule(i)
		t.Log(data)
	}
}

func TestLinkedListLength(t *testing.T) {
	linkedList := initLinkedList1()
	len := linkedList.Length()
	t.Log(len)
}

func TestLinkedListDelete(t *testing.T) {
	linkedList := initLinkedList1()
	linkedList.Delete("1111")
	for linkedList != nil {
		t.Log(linkedList)
		linkedList = linkedList.nextNode
	}
}

func TestLinkedListClear(t *testing.T) {
	linkedList := initLinkedList1()
	linkedList.Clear()
	t.Log(linkedList)
}

func TestLinkedListMidInsert(t *testing.T) {
	linkedList := initLinkedList1()
	if linkedList.Insert(3, "5555") {
		for linkedList != nil {
			t.Log(linkedList)
			linkedList = linkedList.nextNode
		}
	} else {
		for linkedList != nil {
			t.Error(linkedList)
			linkedList = linkedList.nextNode
		}
	}
}

func TestLinkedListDelHead(t *testing.T) {
	linkedList := initLinkedList1()
	linkedList.DeleteF()
	for linkedList != nil {
		t.Log(linkedList)
		linkedList = linkedList.nextNode
	}
}

func TestLinkedListDelTail(t *testing.T) {
	linkedList := initLinkedList1()
	linkedList.DeleteT()
	linkedList.DeleteT()
	linkedList.DeleteT()
	linkedList.DeleteT()
	for linkedList != nil {
		t.Log(linkedList)
		linkedList = linkedList.nextNode
	}
}
