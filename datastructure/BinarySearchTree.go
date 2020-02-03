package TestProject

import (
	"strconv"
)

type BinarySearchTree struct {
	Root  *BinarySearchNode
	index int
}

type BinarySearchNode struct {
	Data  int
	Left  *BinarySearchNode
	Right *BinarySearchNode
}

func NewBinarySearchTree() *BinarySearchTree {
	return new(BinarySearchTree)
}

func NewBinarySearchTreeNode(s int) *BinarySearchNode {
	return &BinarySearchNode{
		Data:  s,
		Left:  nil,
		Right: nil,
	}
}

func checkBinarySearchTree(tree *BinarySearchTree) {
	if tree == nil {
		panic("no binarySearchTree created")
	}
}

func (tree *BinarySearchTree) InitBinarySearchTree(str_p []string, str_m []string) *BinarySearchNode {
	checkBinarySearchTree(tree)
	if len(str_m) <= 0 || len(str_p) <= 0 {
		return nil
	}
	var index = searchRootIndex(str_p, str_m)
	if str_m[index] != "#" {
		str, _ := strconv.Atoi(str_p[0])
		root := NewBinarySearchTreeNode(str)
		root.Left = tree.InitBinarySearchTree(str_p[1:index+1], str_m[:index])
		root.Right = tree.InitBinarySearchTree(str_p[index+1:], str_m[index+1:])
		tree.Root = root
		return root
	}
	return nil
}

func (tree *BinarySearchTree) SearchMinElementInBinarySearchTree(root *BinarySearchNode) *BinarySearchNode {
	checkBinarySearchTree(tree)
	if root == nil {
		return nil
	}
	if root.Left == nil {
		return root
	} else {
		return tree.SearchMinElementInBinarySearchTree(root.Left)
	}
}

func (tree *BinarySearchTree) SearchMaxElementInBinarySearchTree(root *BinarySearchNode) *BinarySearchNode {
	checkBinarySearchTree(tree)
	if root == nil {
		return nil
	}
	if root.Right == nil {
		return root
	} else {
		return tree.SearchMaxElementInBinarySearchTree(root.Right)
	}
}

func (tree *BinarySearchTree) SearchElementInBinarySearchTree(root *BinarySearchNode, element int) *BinarySearchNode {
	checkBinarySearchTree(tree)
	if root == nil {
		return nil
	}
	if root.Data < element {
		return tree.SearchElementInBinarySearchTree(root.Right, element)
	} else if root.Data > element {
		return tree.SearchElementInBinarySearchTree(root.Left, element)
	} else {
		return root
	}
}
