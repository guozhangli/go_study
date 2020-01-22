package TestProject

import (
	"fmt"
)

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
		tree.Root = n
		return n
	}
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
		tree.Root = n
		return n
	}
	return n
}

func (tree *BinaryTree) CreateLevelBinaryTree(n *BinaryTreeNode, str []string) *BinaryTreeNode {
	checkBinaryTree(tree)
	var count = tree.index
	item := str[count]
	if item != "#" {
		n := NewBinaryTreeNode(item)
		indexLeft := count*2 + 1
		if indexLeft < len(str) && str[indexLeft] != "#" {
			tree.index = count*2 + 1
			n.Left = tree.CreateLevelBinaryTree(n.Left, str)
		}
		indexLight := count*2 + 2
		if indexLight < len(str) && str[indexLight] != "#" {
			tree.index = count*2 + 2
			n.Right = tree.CreateLevelBinaryTree(n.Right, str)
		}
		tree.Root = n
		return n
	}
	return n
}

func (tree *BinaryTree) BinaryTreePreOrder(n *BinaryTreeNode) {
	checkBinaryTree(tree)
	if n != nil {
		fmt.Println(n.Data)
		tree.BinaryTreePreOrder(n.Left)
		tree.BinaryTreePreOrder(n.Right)
	}
}

func (tree *BinaryTree) BinaryTreeMidOrder(n *BinaryTreeNode) {
	checkBinaryTree(tree)
	if n != nil {
		tree.BinaryTreeMidOrder(n.Left)
		fmt.Println(n.Data)
		tree.BinaryTreeMidOrder(n.Right)
	}
}

func (tree *BinaryTree) BinaryTreePostOrder(n *BinaryTreeNode) {
	checkBinaryTree(tree)
	if n != nil {
		tree.BinaryTreePostOrder(n.Left)
		tree.BinaryTreePostOrder(n.Right)
		fmt.Println(n.Data)
	}
}

func (tree *BinaryTree) BinaryTreePreOrderStack() {
	if tree.Root != nil {
		stack := NewStackLinked()
		root := tree.Root
		for {
			for root != nil {
				stack.Push(root)
				fmt.Println(root.Data)
				root = root.Left
			}
			if stack.IsEmpty() {
				break
			}
			root = stack.Pop().(*Node).Data.(*BinaryTreeNode)
			root = root.Right
		}
	}
}

func (tree *BinaryTree) BinaryTreeMidOrderStack() {
	if tree.Root != nil {
		stack := NewStackLinked()
		root := tree.Root
		for {
			for root != nil {
				stack.Push(root)
				root = root.Left
			}
			if stack.IsEmpty() {
				break
			}
			root = stack.Pop().(*Node).Data.(*BinaryTreeNode)
			fmt.Println(root.Data)
			root = root.Right
		}
	}
}

func (tree *BinaryTree) BinaryTreePostOrderStack() {
	if tree.Root != nil {
		stack := NewStackLinked()
		root := tree.Root
		var s []string
		for {
			for root != nil {
				stack.Push(root)
				s = append(s, root.Data.(string))
				root = root.Right
			}
			if stack.IsEmpty() {
				break
			}
			root = stack.Pop().(*Node).Data.(*BinaryTreeNode)
			root = root.Left
		}
		for i := len(s) - 1; i >= 0; i-- {
			fmt.Println(s[i])
		}
	}
}

func (tree *BinaryTree) BinaryTreeLevelOrderQueue() {
	if tree.Root != nil {
		queue := NewQueueLinked()
		root := tree.Root
		queue.EnQueue(root)
		fmt.Println(root.Data)
		for !queue.IsEmpty() {
			root = queue.DeQueue().(*Node).Data.(*BinaryTreeNode)
			if root.Left != nil {
				queue.EnQueue(root.Left)
				fmt.Println(root.Left.Data)
			}
			if root.Right != nil {
				queue.EnQueue(root.Right)
				fmt.Println(root.Right.Data)
			}
		}
	}
}
