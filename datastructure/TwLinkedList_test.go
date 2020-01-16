package TestProject

import (
	"testing"
)

func TestCreateTwLinkedList(t *testing.T) {
	twLinkedList := newTwLinkedList(nil)
	t.Log(twLinkedList)
}

func initTwLinkedList() *TwLinkedList {
	twLinkedList := newTwLinkedList(nil)
	return twLinkedList
}

func TestTwLinkedListAddF(t *testing.T) {
	twLinkedList := initTwLinkedListAddValue()
	for twLinkedList != nil {
		t.Log(twLinkedList)
		twLinkedList = twLinkedList.nextNode
	}
}

func initTwLinkedListAddValue() *TwLinkedList {
	twLinkedList := initTwLinkedList()
	twLinkedList.HeadAdd("1111")
	twLinkedList.HeadAdd("2222")
	twLinkedList.HeadAdd("3333")
	twLinkedList.HeadAdd("4444")
	return twLinkedList
}

func TestTwLinkedListAddT(t *testing.T) {
	twLinkedList := initTwLinkedList()
	twLinkedList.TailAdd("1111")
	twLinkedList.TailAdd("2222")
	twLinkedList.TailAdd("3333")
	twLinkedList.TailAdd("4444")
	for twLinkedList != nil {
		t.Log(twLinkedList)
		twLinkedList = twLinkedList.nextNode
	}
}

func TestTwLinkedListInsert(t *testing.T) {
	twLinkedList := initTwLinkedListAddValue()
	twLinkedList.Insert(3, "5555")
	for twLinkedList != nil {
		t.Log(twLinkedList)
		twLinkedList = twLinkedList.nextNode
	}
}

func TestTwLinkedListLength(t *testing.T) {
	twLinkedList := initTwLinkedListAddValue()
	len := twLinkedList.Length()
	t.Log(len)
}

func TestTwLinkedListIsEmpty(t *testing.T) {
	twLinkedList := initTwLinkedList()
	if twLinkedList.IsEmpty() {
		t.Log("twLinkedList is empty")
	} else {
		t.Log("twLinkedList is not empty")
	}
}

func TestTwLinkedListGetValue(t *testing.T) {
	twLinkedList := initTwLinkedListAddValue()
	val := twLinkedList.GetValue(4)
	t.Log(val)
	for twLinkedList != nil {
		t.Log(twLinkedList)
		twLinkedList = twLinkedList.nextNode
	}
}

func TestTwLinkedListDelete(t *testing.T) {
	twLinkedList := initTwLinkedListAddValue()
	twLinkedList.Delete("1111")
	twLinkedList.Delete("2222")
	twLinkedList.Delete("3333")
	twLinkedList.Delete("4444")
	for twLinkedList != nil {
		t.Log(twLinkedList)
		twLinkedList = twLinkedList.nextNode
	}
}

func TestTwLinkedListDeleteF(t *testing.T) {
	twLinkedList := initTwLinkedListAddValue()
	twLinkedList.DeleteF()
	for twLinkedList != nil {
		t.Log(twLinkedList)
		twLinkedList = twLinkedList.nextNode
	}
}

func TestTwLinkedListDeleteT(t *testing.T) {
	twLinkedList := initTwLinkedListAddValue()
	twLinkedList.DeleteT()
	twLinkedList.DeleteT()
	twLinkedList.DeleteT()
	for twLinkedList != nil {
		t.Log(twLinkedList)
		twLinkedList = twLinkedList.nextNode
	}
}

func TestTwLinkedListClear(t *testing.T) {
	twLinkedList := initTwLinkedListAddValue()
	twLinkedList.Clear()
	for twLinkedList != nil {
		t.Log(twLinkedList)
		twLinkedList = twLinkedList.nextNode
	}
}
