package TestProject

import "fmt"

type BinaryTreeNode struct {
	Data  interface{}
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

func NewBinaryTree(vaule interface{}) *BinaryTreeNode {
	return &BinaryTreeNode{
		Data:  vaule,
		Left:  nil,
		Right: nil,
	}
}

func CreateBinaryTree(node *BinaryTreeNode, str []string, count int) *BinaryTreeNode {
	fmt.Println("test1")
	fmt.Println(str[count], count)
	if node == nil {
		node = NewBinaryTree(str[count])
	}
	if str[count] != "#" {
		count++
		node.Left = CreateBinaryTree(node.Left, str, count)
		fmt.Println("test2")
		count++
		node.Right = CreateBinaryTree(node.Right, str, count)
		fmt.Println("test3")
	}
	return node
}
