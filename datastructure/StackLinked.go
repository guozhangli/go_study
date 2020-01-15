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
