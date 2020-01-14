package TestProject

import (
	"testing"
)

func TestCreateTwLinkedList(t *testing.T){
	twLinkedList:=newTwLinkedList(nil)
	t.Log(twLinkedList)
}

func initTwLinkedList() *TwLinkedList{
	twLinkedList:=newTwLinkedList(nil)
	return twLinkedList
}

func TestTwLinkedListAddF(t *testing.T){
	twLinkedList:=initTwLinkedListAddValue()
	for twLinkedList!=nil{
		t.Log(twLinkedList)
		twLinkedList=twLinkedList.nextNode
	}
}

func initTwLinkedListAddValue() *TwLinkedList{
	twLinkedList:=initTwLinkedList()
	twLinkedList.addHead("1111")
	twLinkedList.addHead("2222")
	twLinkedList.addHead("3333")
	twLinkedList.addHead("4444")
	return twLinkedList
}

func TestTwLinkedListAddT(t *testing.T){
	twLinkedList:=initTwLinkedList()
	twLinkedList.addTail("1111")
	twLinkedList.addTail("2222")
	twLinkedList.addTail("3333")
	twLinkedList.addTail("4444")
	for twLinkedList!=nil{
		t.Log(twLinkedList)
		twLinkedList=twLinkedList.nextNode
	}
}

func TestTwLinkedListInsert(t *testing.T){
	twLinkedList:=initTwLinkedListAddValue()
    twLinkedList.insert(3,"5555")
	for twLinkedList!=nil{
		t.Log(twLinkedList)
		twLinkedList=twLinkedList.nextNode
	}
}

func TestTwLinkedListLength(t *testing.T){
	twLinkedList:=initTwLinkedListAddValue()
	len:=twLinkedList.length()
	t.Log(len)
}

func TestTwLinkedListIsEmpty(t *testing.T){
	twLinkedList:=initTwLinkedList()
	if twLinkedList.isEmpty(){
		t.Log("twLinkedList is empty")
	}else{
		t.Log("twLinkedList is not empty")
	}
}

func TestTwLinkedListGetValue(t *testing.T){
	twLinkedList:=initTwLinkedListAddValue()
	val:=twLinkedList.getValue(4)
	t.Log(val)
	for twLinkedList!=nil{
		t.Log(twLinkedList)
		twLinkedList=twLinkedList.nextNode
	}
}

func TestTwLinkedListDelete(t *testing.T){
	twLinkedList:=initTwLinkedListAddValue()
	twLinkedList.delete("1111")
	twLinkedList.delete("2222")
	twLinkedList.delete("3333")
	twLinkedList.delete("4444")
	for twLinkedList!=nil{
		t.Log(twLinkedList)
		twLinkedList=twLinkedList.nextNode
	}
}

func TestTwLinkedListDeleteF(t *testing.T){
	twLinkedList:=initTwLinkedListAddValue()
	twLinkedList.deleteF()
	for twLinkedList!=nil{
		t.Log(twLinkedList)
		twLinkedList=twLinkedList.nextNode
	}
}

func TestTwLinkedListDeleteT(t *testing.T){
	twLinkedList:=initTwLinkedListAddValue()
	twLinkedList.deleteT()
	twLinkedList.deleteT()
	twLinkedList.deleteT()
	for twLinkedList!=nil{
		t.Log(twLinkedList)
		twLinkedList=twLinkedList.nextNode
	}
}

func TestTwLinkedListClear(t *testing.T){
	twLinkedList:=initTwLinkedListAddValue()
	twLinkedList.clear()
	for twLinkedList!=nil{
		t.Log(twLinkedList)
		twLinkedList=twLinkedList.nextNode
	}
}

