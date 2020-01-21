package TestProject

import "fmt"

type BinaryTreeNode struct {
	Data  interface{}
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

type BinaryTree struct {
	Root  *BinaryTreeNode
	index int
}

func NewBinaryTreeNode(vaule interface{}) *BinaryTreeNode {
	return &BinaryTreeNode{
		Data:  vaule,
		Left:  nil,
		Right: nil,
	}
}

func checkBinaryTree(tree *BinaryTree) {
	if tree == nil {
		panic("no binaryTree created")
	}
}
func (tree *BinaryTree) CreatePreBinaryTree(n *BinaryTreeNode, str []string) *BinaryTreeNode {
	checkBinaryTree(tree)
	if str[tree.index] != "#" {
		n = NewBinaryTreeNode(str[tree.index])
		tree.index++
		n.Left = tree.CreatePreBinaryTree(n.Left, str)
		tree.index++
		n.Right = tree.CreatePreBinaryTree(n.Right, str)
	}
	tree.Root = n
	return n
}

func (tree *BinaryTree) CreateMidBinaryTree(str_p []string, str_m []string) *BinaryTreeNode {
	checkBinaryTree(tree)
	if len(str_m) <= 0 || len(str_p) <= 0 {
		return nil
	}
	var index = searchRootIndex(str_p, str_m)
	if str_m[index] != "#" {
		root := NewBinaryTreeNode(str_p[0])
		root.Left = tree.CreateMidBinaryTree(str_p[1:index+1], str_m[:index])
		root.Right = tree.CreateMidBinaryTree(str_p[index+1:], str_m[index+1:])
		tree.Root = root
		return root
	}
	return nil
}

func searchRootIndex(str_p []string, str_m []string) int {
	for i, v := range str_m {
		if v == str_p[0] {
			return i
		}
	}
	return -1
}

func (tree *BinaryTree) CreatePostBinaryTree(n *BinaryTreeNode, str []string) *BinaryTreeNode {
	checkBinaryTree(tree)
	item := str[len(str)-1-tree.index]
	tree.index++
	if item != "#" {
		n := NewBinaryTreeNode(item)
		n.Right = tree.CreatePostBinaryTree(n.Right, str)
		n.Left = tree.CreatePostBinaryTree(n.Left, str)
		fmt.Printf("%v", n)
		tree.Root = n
	}
	return n
}
