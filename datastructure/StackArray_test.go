package TestProject

import "testing"

func TestStackArrayCreate(t *testing.T) {
	stackArray := NewStackArray(5)
	t.Log(stackArray)
}

func initStack() Stack {
	stackArray := NewStackArray(5)
	return stackArray
}

func TestStackArrayIsEmpty(t *testing.T) {
	sa := initStack()
	if sa.IsEmpty() {
		t.Log("stack is empty")
	} else {
		t.Log("stack is not empty")
	}
}

func TestStackArrayLength(t *testing.T) {
	sa := initStack()
	len := sa.Length()
	t.Log(len)
}

func TestStackArrayPush(t *testing.T) {
	sa := initStack()
	sa.Push("1111")
	sa.Push("2222")
	sa.Push("3333")
	len := sa.Length()
	t.Log(len)
	t.Log(sa)
}

func TestStackArrayPop(t *testing.T) {
	sa := initStack()
	sa.Push("1111")
	sa.Push("2222")
	sa.Push("3333")
	sa.Push("4444")
	sa.Push("5555")
	sa.Push("6666")
	sa.Push("7777")
	sa.Push("8888")
	sa.Push("9999")
	val := sa.Pop()
	val1 := sa.Pop()
	t.Log(sa)
	t.Log(val)
	t.Log(val1)
}

func TestStackArrayPop_1(t *testing.T) {
	sa := initStack()
	val := sa.Pop()
	val1 := sa.Pop()
	t.Log(sa)
	t.Log(val)
	t.Log(val1)
}

func TestStackArrayClear(t *testing.T) {
	sa := initStack()
	sa.Push("1111")
	sa.Push("2222")
	sa.Push("3333")
	sa.Push("4444")
	sa.Push("5555")
	sa.Push("6666")
	sa.Push("7777")
	sa.Push("8888")
	sa.Push("9999")
	sa.Clear()
	t.Log(sa)
}

func TestStackArrayDistory(t *testing.T) {
	sa := initStack()
	sa.Push("1111")
	sa.Push("2222")
	sa.Push("3333")
	sa.Push("4444")
	sa.Push("5555")
	sa.Push("6666")
	sa.Push("7777")
	sa.Push("8888")
	sa.Push("9999")
	top := sa.GetTop()
	t.Log(top)
	sa.Clear()
	t.Log(sa)
	sa.Distroy(&sa)
	t.Log(sa)
}
