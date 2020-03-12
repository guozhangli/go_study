package TestProject

import (
	"strconv"
)

type BinarySearchTree struct {
	Root *BinarySearchNode
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

func (tree *BinarySearchTree) SearchMinElementInBinarySearchTreeForeach(root *BinarySearchNode) *BinarySearchNode {
	checkBinarySearchTree(tree)
	if root == nil {
		return nil
	}
	for root.Left != nil {
		root = root.Left
	}
	return root
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

func (tree *BinarySearchTree) SearchMaxElementInBinarySearchTreeForeach(root *BinarySearchNode) *BinarySearchNode {
	checkBinarySearchTree(tree)
	if root == nil {
		return nil
	}
	for root.Right != nil {
		root = root.Right
	}
	return root
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

func (tree *BinarySearchTree) SearchElementInBinarySearchTreeForeach(root *BinarySearchNode, element int) *BinarySearchNode {
	checkBinarySearchTree(tree)
	if root == nil {
		return nil
	}
	for root != nil {
		if root.Data < element {
			root = root.Right
		} else if root.Data > element {
			root = root.Left
		} else {
			return root
		}
	}
	return root
}

func (tree *BinarySearchTree) InsertElementInBinarySearchTree(root *BinarySearchNode, element int) {
	checkBinarySearchTree(tree)
	node := NewBinarySearchTreeNode(element)
	if root == nil {
		root = node
	} else {
		var pre *BinarySearchNode
		for root != nil {
			if root.Data < element {
				if root != nil {
					pre = root
				}
				root = root.Right
			} else if root.Data > element {
				if root != nil {
					pre = root
				}
				root = root.Left
			} else {
				pre = nil
				break
			}
		}
		if pre != nil {
			if pre.Data > node.Data {
				pre.Left = node
			} else {
				pre.Right = node
			}
		}
	}
}

func (tree *BinarySearchTree) DeleteElementInBinarySearchTree(root **BinarySearchNode, element int) {
	checkBinarySearchTree(tree)
	var temp *BinarySearchNode
	for *root != nil {
		if (*root).Data > element {
			root = &(*root).Left
		} else if (*root).Data < element {
			root = &(*root).Right
		} else {
			if (*root).Left == nil && (*root).Right == nil {
				*root = nil
			} else if (*root).Left != nil && (*root).Right != nil {
				temp = tree.SearchMaxElementInBinarySearchTreeForeach((*root).Left)
				(*root).Data = temp.Data
				tree.DeleteElementInBinarySearchTree(&(*root).Left, (*root).Data)
			} else if (*root).Left != nil && (*root).Right == nil {
				*root = (*root).Left
			} else if (*root).Left == nil && (*root).Right != nil {
				*root = (*root).Right
			}
		}
	}
}

func (tree *BinarySearchTree) GetHight() int {
	checkBinarySearchTree(tree)
	queue := NewQueueLinked()
	queue.EnQueue(tree.Root)
	queue.EnQueue(nil)
	var level int
	for !queue.IsEmpty() {
		var node *BinarySearchNode
		temp := queue.DeQueue().(*Node).Data
		if temp != nil {
			node = temp.(*BinarySearchNode)
		}
		if node == nil {
			if !queue.IsEmpty() {
				queue.EnQueue(nil)
			}
			level++
		} else {
			if node.Left != nil {
				queue.EnQueue(node.Left)
			}
			if node.Right != nil {
				queue.EnQueue(node.Right)
			}
		}
	}
	return level
}
