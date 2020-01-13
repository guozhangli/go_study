package TestProject

import "testing"


func initLinkedList() *LinkedList{
	linkedList:=new(LinkedList)
	return linkedList
}
func TestCreateNode(t *testing.T){
	linkedList:=initLinkedList()
	if linkedList!=nil{
		t.Log(linkedList)
	}
}

func TestLinkedListAdd(t *testing.T){
	linkedList:=initLinkedList1()
	for ; linkedList != nil; {
		t.Log(linkedList)
		linkedList = linkedList.nextNode
	}
}

func TestLinkedListHeadAdd(t *testing.T){
	linkedList:=initLinkedList()
	linkedList.headAdd("1111")
	linkedList.headAdd("2222")
	linkedList.headAdd("3333")
	linkedList.headAdd("4444")
	for ; linkedList != nil; {
		t.Log(linkedList)
		linkedList = linkedList.nextNode
	}
}

func initLinkedList1() *LinkedList{
	linkedList:=initLinkedList()
	linkedList.tailAdd("1111")
	linkedList.tailAdd("2222")
	linkedList.tailAdd("3333")
	linkedList.tailAdd("4444")
	return linkedList
}

func TestLinkedListGetVaule(t *testing.T){
	linkedList:=initLinkedList1()
	linkedList.getVaule(1)
}

func TestLinkedListLength(t *testing.T){
	linkedList:=initLinkedList1()
	len:=linkedList.length()
	t.Log(len)
}

