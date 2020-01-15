package TestProject

type Node struct {
	Data interface{}
	Next *Node
}

type StackLinked struct {
	Top    *Node
	Length int
}

func newStackLinked() *StackLinked {
	stackLinked := new(StackLinked)
	return stackLinked
}

func newNode(top *Node, value interface{}) *Node {
	node := new(Node)
	node.Data = value
	node.Next = top
	return node
}
func checkStackLinked(stack *StackLinked) {
	if stack == nil {
		panic("no stack created")
	}
}
func (stack *StackLinked) Push(value interface{}) {
	checkStackLinked(stack)
	node := newNode(stack.Top, value)
	stack.Top = node
	stack.Length++
}

func (stack *StackLinked) Pop() *Node {
	checkStackLinked(stack)
	var top = stack.Top
	stack.Top = stack.Top.Next
	stack.Length--
	return top
}

func (stack *StackLinked) IsEmpty() bool{
	checkStackLinked(stack)
	if stack.Length==0{
		return true
	}
	return false
}

func (stack *StackLinked) Clear() {
	checkStackLinked(stack)
	for stack.Length>0{
		stack.Pop()
	}
}

func (stack *StackLinked) Distroy(s **StackLinked) {
	checkStackLinked(stack)
	*s=nil
}
