package TestProject

import (
	"testing"
)

func initArrayList(capacity uint) *ArrayList {
	list := newArray(capacity)
	return list
}

func TestArrayListIsEmpty(t *testing.T) {
	list := initArrayList(8)
	len := list.IsEmpty()
	if len {
		t.Log("list is empty")
	} else {
		t.Error("list not empty")
	}
}

func TestArrayListAdd(t *testing.T) {
	list := initArrayList(8)
	for i := 0; i < 8; i++ {
		if list.Add(i) {
			t.Log("ok")
		} else {
			t.Error("err")
		}
	}
	t.Log(list)

}

func TestArrayListSize(t *testing.T) {
	list := initArrayList(8)
	list.Add("111")
	list.Add("222")
	list.Add("333")
	if list.Length() == 3 {
		t.Log(3)
	} else {
		t.Error("err")
	}
}

func TestArrayListInsert(t *testing.T) {
	list := initArrayList(8)
	if list.Insert(1, "adadf") {
		t.Log(list)
	} else {
		t.Error(list)
	}
}

func TestArrayListGetValue(t *testing.T) {
	list := initArrayList(8)
	list.Add("111")
	list.Add("222")
	list.Add("333")
	val := list.GetValue(2)
	if val != nil {
		t.Log(val)
	} else {
		t.Error(val)
	}
}

func TestArrayListDelete(t *testing.T) {
	list := initArrayList(8)
	list.Add("111")
	list.Add("222")
	list.Add("333")
	list.Add("444")
	if list.Delete("444") {
		t.Log(list)
	} else {
		t.Error(list)
	}
}

func TestArrayListAddConcurrence(t *testing.T) {
	list := initArrayList(10000)
	for i := 0; i < 10000; i++ {
		go list.Add(i)
	}
	t.Log(list)
}
