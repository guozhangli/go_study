package TestProject

import "testing"

func TestCycLinkedList(t *testing.T) {
	cycLinkedList := NewCyclicLinkedList(nil)
	t.Log(cycLinkedList)
}

func initCycLinkedList() *CyclicLinkedList {
	cycLinkedList := NewCyclicLinkedList(nil)
	return cycLinkedList
}

func TestCycLinkedListHeadAdd(t *testing.T) {
	cycLinkedList := initCycLinkedList()
	head := cycLinkedList
	cycLinkedList.HeadAdd("1111")
	cycLinkedList.HeadAdd("2222")
	cycLinkedList.HeadAdd("3333")
	cycLinkedList.HeadAdd("4444")
	cycLinkedList.Print(head)
}

func TestCycLinkedListTailAdd(t *testing.T) {
	cycLinkedList, head := initCycLinkedListAddValue()
	cycLinkedList.Print(head)
}

func TestCycLinkedListInsert(t *testing.T) {
	cycLinkedList, head := initCycLinkedListAddValue()
	if cycLinkedList.insert(4, "5555") {
		cycLinkedList.Print(head)
	} else {
		t.Log("insert failed")
	}

}

func initCycLinkedListAddValue() (*CyclicLinkedList, *CyclicLinkedList) {
	cycLinkedList := initCycLinkedList()
	head := cycLinkedList
	cycLinkedList.TailAdd("1111")
	cycLinkedList.TailAdd("2222")
	cycLinkedList.TailAdd("3333")
	cycLinkedList.TailAdd("4444")
	return cycLinkedList, head
}

func TestCycLinkedListLength(t *testing.T) {
	cycLinkedList, _ := initCycLinkedListAddValue()
	len := cycLinkedList.length()
	t.Log(len)
}
