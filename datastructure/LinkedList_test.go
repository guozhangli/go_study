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
	linkedList.headAdd("1111")
	linkedList.headAdd("2222")
	linkedList.headAdd("3333")
	linkedList.headAdd("4444")
	for linkedList != nil {
		t.Log(linkedList)
		linkedList = linkedList.nextNode
	}
}

func initLinkedList1() *LinkedList {
	linkedList := initLinkedList()
	linkedList.tailAdd("1111")
	linkedList.tailAdd("2222")
	linkedList.tailAdd("3333")
	linkedList.tailAdd("4444")
	return linkedList
}

func TestLinkedListGetVaule(t *testing.T) {
	linkedList := initLinkedList1()
	for i := 0; i < 8; i++ {
		data := linkedList.getVaule(i)
		t.Log(data)
	}
}

func TestLinkedListLength(t *testing.T) {
	linkedList := initLinkedList1()
	len := linkedList.length()
	t.Log(len)
}

func TestLinkedListDelete(t *testing.T) {
	linkedList := initLinkedList1()
	linkedList.delete("1111")
	for linkedList != nil {
		t.Log(linkedList)
		linkedList = linkedList.nextNode
	}
}

func TestLinkedListClear(t *testing.T) {
	linkedList := initLinkedList1()
	linkedList.clear()
	t.Log(linkedList)
}

func TestLinkedListMidInsert(t *testing.T){
	linkedList := initLinkedList1()
	if linkedList.insert(3,"5555"){
		for linkedList != nil {
			t.Log(linkedList)
			linkedList = linkedList.nextNode
		}
	}else{
		for linkedList != nil {
			t.Error(linkedList)
			linkedList = linkedList.nextNode
		}
	}
}

func TestLinkedListDelHead(t *testing.T){
	linkedList := initLinkedList1()
	linkedList.deleteF()
	for linkedList != nil {
		t.Log(linkedList)
		linkedList = linkedList.nextNode
	}
}

func TestLinkedListDelTail(t *testing.T){
	linkedList := initLinkedList1()
	linkedList.deleteT()
	linkedList.deleteT()
	linkedList.deleteT()
	linkedList.deleteT()
	for linkedList != nil {
		t.Log(linkedList)
		linkedList = linkedList.nextNode
	}
}
