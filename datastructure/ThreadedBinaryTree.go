package TestProject

import "fmt"

type ThreadedBinaryTreeNode struct {
	Left  *ThreadedBinaryTreeNode
	Right *ThreadedBinaryTreeNode
	LTag  int
	RTag  int
	Data  interface{}
}
type ThreadedBinaryTree struct {
	Root  *ThreadedBinaryTreeNode
	Index int
}

func NewThreadedBinaryNode(value interface{}) *ThreadedBinaryTreeNode {
	return &ThreadedBinaryTreeNode{
		Left:  nil,
		Right: nil,
		LTag:  0,
		RTag:  0,
		Data:  value,
	}
}

func NewThreadedBinaryTree() *ThreadedBinaryTree {
	return new(ThreadedBinaryTree)
}

func checkThreadedBinaryTree(tree *ThreadedBinaryTree) {
	if tree == nil {
		panic("no threadedBinaryTree created")
	}
}

func (tree *ThreadedBinaryTree) CreateMidThreadedBinaryTree(str_p []string, str_m []string) *ThreadedBinaryTreeNode {
	checkThreadedBinaryTree(tree)
	if len(str_m) <= 0 || len(str_p) <= 0 {
		return nil
	}
	var index = searchRootIndex(str_p, str_m)
	if str_m[index] != "#" {
		root := NewThreadedBinaryNode(str_p[0])
		root.Left = tree.CreateMidThreadedBinaryTree(str_p[1:index+1], str_m[:index])
		root.Right = tree.CreateMidThreadedBinaryTree(str_p[index+1:], str_m[index+1:])
		tree.Root = root
		return root
	}
	return nil
}

var pre *ThreadedBinaryTreeNode

func (tree *ThreadedBinaryTree) ThreadedBinaryTreeThreading(n *ThreadedBinaryTreeNode) {
	checkThreadedBinaryTree(tree)
	if n != nil {
		tree.ThreadedBinaryTreeThreading(n.Left)
		if n.Left == nil {
			n.LTag = 1
			n.Left = pre
		}
		if pre != nil && pre.Right == nil {
			pre.RTag = 1
			pre.Right = n
		}
		pre = n
		tree.ThreadedBinaryTreeThreading(n.Right)
	}
}

func (tree *ThreadedBinaryTree) ThreadedBinaryTreeMidForEachAsc() {
	checkThreadedBinaryTree(tree)
	node := tree.Root
	for node != nil && node.LTag == 0 {
		node = node.Left
	}
	for node != nil {
		fmt.Println(node.Data)
		if node.RTag == 1 {
			node = node.Right
		} else {
			node = node.Right
			for node != nil && node.LTag == 0 {
				node = node.Left
			}
		}
	}

}

func (tree *ThreadedBinaryTree) ThreadedBinaryTreeMidForEachDesc() {
	checkThreadedBinaryTree(tree)
	node := tree.Root
	for node.Right != nil && node.RTag == 0 {
		node = node.Right
	}
	for node != nil {
		fmt.Println(node.Data)
		if node.LTag == 1 {
			node = node.Left
		} else {
			node = node.Left
			for node.Right != nil && node.RTag == 0 {
				node = node.Right
			}
		}
	}
}
