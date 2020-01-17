package TestProject

type Node struct {
	Data interface{}
	Next *Node
}

func NewNode(next *Node, value interface{}) *Node {
	node := new(Node)
	node.Data = value
	node.Next = next
	return node
}
