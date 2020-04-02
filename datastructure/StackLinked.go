package TestProject

type StackLinked struct {
	Top *Node
	Len int
}

func NewStackLinked() Stack {
	stackLinked := new(StackLinked)
	return stackLinked
}

func checkStackLinked(stack *StackLinked) {
	if stack == nil {
		panic("no stack created")
	}
}
func (stack *StackLinked) Push(value interface{}) {
	checkStackLinked(stack)
	node := NewNode(stack.Top, value)
	stack.Top = node
	stack.Len++
}

func (stack *StackLinked) Pop() interface{} {
	checkStackLinked(stack)
	var top = stack.Top
	if top != nil {
		stack.Top = top.Next
		stack.Len--
	}
	return top
}

func (stack *StackLinked) Pop2() interface{} {
	checkStackLinked(stack)
	var top = stack.Top
	if top != nil {
		stack.Top = top.Next
		stack.Len--
	}
	return top.Data
}

func (stack *StackLinked) IsEmpty() bool {
	checkStackLinked(stack)
	if stack.Len == 0 {
		return true
	}
	return false
}

func (stack *StackLinked) Length() int {
	return stack.Len
}

func (stack *StackLinked) Clear() {
	checkStackLinked(stack)
	for stack.Len > 0 {
		stack.Pop()
	}
}

func (stack *StackLinked) Distroy(s *Stack) {
	checkStackLinked(stack)
	*s = nil
}

func (stack *StackLinked) GetTop() interface{} {
	checkStackLinked(stack)
	if stack.Top != nil {
		return stack.Top.Data
	}
	return nil
}
