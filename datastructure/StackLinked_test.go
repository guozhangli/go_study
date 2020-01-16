package TestProject

import (
	json "encoding/json"
	"fmt"
	"testing"
)

func TestStackLinkedCreate(t *testing.T) {
	sl := NewStackLinked()
	t.Log(sl)
}

func TestStackLinkedPush(t *testing.T) {
	sl := NewStackLinked()
	sl.Push("1111")
	sl.Push("2222")
	sl.Push("3333")
	t.Logf("%+v", sl)
	json, _ := json.MarshalIndent(sl, "", " ")
	fmt.Printf("%s\n", json)
}

func TestStackLinkedPop(t *testing.T) {
	sl := NewStackLinked()
	sl.Push("1111")
	sl.Push("2222")
	sl.Push("3333")
	sl.Push("4444")
	top := sl.Pop()
	t.Logf("%+v", top)
	top1 := sl.Pop()
	t.Logf("%+v", top1)
	json, _ := json.MarshalIndent(sl, "", " ")
	fmt.Printf("%s\n", json)
}

func TestStackLinkedLength(t *testing.T) {
	sl := NewStackLinked()

	json, _ := json.MarshalIndent(sl, "", " ")
	fmt.Printf("%s\n", json)

	if sl.IsEmpty() {
		t.Log("stackLinked is empty")
	} else {
		t.Log("stackLinked is not empty")
	}

}

func TestStackLinkedClear(t *testing.T) {
	sl := NewStackLinked()
	sl.Push("1111")
	sl.Push("2222")
	sl.Push("3333")
	sl.Push("4444")
	sl.Clear()
	t.Log(sl)
}

func TestStackLinkedDistroy(t *testing.T) {
	sl := NewStackLinked()
	sl.Push("1111")
	sl.Push("2222")
	sl.Push("3333")
	sl.Push("4444")
	sl.Clear()
	t.Log(sl)
	sl.Distroy(&sl)
	t.Log(sl)
}

