package TestProject

import (
	"encoding/json"
	"testing"
)

func printRBTree(tree *RBTree, t *testing.T) {
	str, _ := json.Marshal(tree)
	t.Log(string(str))
}

func InitRBTree() *RBTree {
	rbTree := NewRBTree()
	rbNode := NewRBNode("A", "aaaaa", 1, RED)
	rbNodeL := NewRBNode("B", "bbbbb", 1, RED)
	rbNode.Left = rbNodeL
	rbTree.Root = rbNode
	return rbTree
}
func InitRBTree2() *RBTree {
	rbTree := NewRBTree()
	rbNode := NewRBNode("A", "aaaaa", 1, RED)
	rbNodeL := NewRBNode("B", "bbbbb", 1, RED)
	rbNode.Right = rbNodeL
	rbTree.Root = rbNode
	return rbTree
}
func TestNewRBTree(t *testing.T) {
	rbTree := InitRBTree()
	printRBTree(rbTree, t)
}

func TestRBNode_Right_Rotate(t *testing.T) {
	rbTree := InitRBTree()
	printRBTree(rbTree, t)
	Right_Rotate(&rbTree.Root)
	printRBTree(rbTree, t)
}

func TestRBNode_Left_Rotate(t *testing.T) {
	rbTree := InitRBTree2()
	printRBTree(rbTree, t)
	Left_Rotate(&rbTree.Root)
	printRBTree(rbTree, t)
}

func TestRBNode_Put(t *testing.T) {
	rbTree := NewRBTree()
	rbTree.Put("S", "sssss")
	rbTree.Put("E", "eeeee")
	rbTree.Put("A", "aaaaa")
	rbTree.Put("R", "rrrrr")
	rbTree.Put("C", "ccccc")
	rbTree.Put("H", "hhhhh")
	rbTree.Put("X", "xxxxx")
	rbTree.Put("M", "mmmmm")
	rbTree.Put("P", "ppppp")
	rbTree.Put("L", "lllll")
	printRBTree(rbTree, t)
}

func TestRBTree_DeleteMin(t *testing.T) {
	rbTree := NewRBTree()
	rbTree.Put("S", "sssss")
	rbTree.Put("E", "eeeee")
	rbTree.Put("A", "aaaaa")
	rbTree.Put("R", "rrrrr")
	rbTree.Put("C", "ccccc")
	rbTree.Put("H", "hhhhh")
	rbTree.Put("X", "xxxxx")
	rbTree.Put("M", "mmmmm")
	rbTree.Put("P", "ppppp")
	rbTree.Put("L", "lllll")
	rbTree.DeleteMin()
	printRBTree(rbTree, t)
}
