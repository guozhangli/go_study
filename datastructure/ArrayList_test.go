package TestProject

import (
	"testing"
)

func initArrayList() *ArrayList{
	list:=newArray(8)
	return list
}

func TestArrayListIsEmpty(t *testing.T){
	list:=initArrayList()
	len:=list.isEmpty()
	if len{
		t.Log("list is empty")
	}else{
		t.Error("list not empty")
	}
}

func TestArrayListAdd(t *testing.T){
	list:=initArrayList()
	for i:=0;i<8;i++{
		if list.add(i){
			t.Log("ok")
		}else{
			t.Error("err")
		}
	}

}

func TestArrayListSize(t *testing.T){
	list:=initArrayList()
	list.add("111")
	list.add("222")
	list.add("333")
	if list.size()==3{
		t.Log(3)
	}else{
		t.Error("err")
	}
}

func TestArrayListInsert(t *testing.T){
	list:=initArrayList()
	if list.insert(1,"adadf"){
		t.Log(list)
	}else{
		t.Error(list)
	}
}

func TestArrayListGetValue(t *testing.T){
	list:=initArrayList()
	list.add("111")
	list.add("222")
	list.add("333")
	val:=list.getValue(2)
	if val!=nil{
		t.Log(val)
	}else{
		t.Error(val)
	}
}

func TestArrayListDelete(t *testing.T){
	list:=initArrayList()
	list.add("111")
	list.add("222")
	list.add("333")
	list.add("444")
	if list.delete("444"){
		t.Log(list)
	}else{
		t.Error(list)
	}
}

func TestArrayListAddConcurrence(t *testing.T){
	list:=initArrayList()
	for i:=0;i<8;i++{
		go list.add(i)
	}
	t.Log(list)
}

