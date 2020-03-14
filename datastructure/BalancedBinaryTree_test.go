package TestProject

import (
	"encoding/json"
	"testing"
)

func printTree(tree *BalancedBinaryTree, t *testing.T) {
	str, _ := json.Marshal(tree)
	t.Log(string(str))
}
func InitAVL() *BalancedBinaryTree {
	node := [10]int{3, 2, 1, 4, 5, 6, 7, 10, 9, 8}
	tree := InitBalancedBinaryTree()
	for _, v := range node {
		tree.InsertNodeInBalancedBinaryTree(&tree.Root, v)
	}
	return tree
}
func TestInsertNodeInBalancedBinaryTree(t *testing.T) {
	tree := InitAVL()
	printTree(tree, t)
}

func TestDeleteNodeInBalancedBinaryTree(t *testing.T) {
	node := [10]int{3, 2, 1, 4, 5, 6, 7, 10, 9, 8}
	tree := InitBalancedBinaryTree()
	for _, v := range node {
		tree.InsertNodeInBalancedBinaryTree(&tree.Root, v)
		printTree(tree, t)
	}
	tree.DeleteNodeInBalancedBinaryTree(&tree.Root, 4)
	printTree(tree, t)
}

func TestSearchMaxElementInBalancedBinaryTree(t *testing.T) {
	tree := InitAVL()
	printTree(tree, t)
	node := tree.SearchMaxElementInBalancedBinaryTree()
	t.Log(node)
}

func TestSearchMaxElementInBalancedBinaryTree_1(t *testing.T) {
	tree := InitAVL()
	printTree(tree, t)
	node := tree.SearchMaxElementInBalancedBinaryTree_1(tree.Root)
	t.Log(node)
}

func TestSearchMinElementInBalancedBinaryTree(t *testing.T) {
	tree := InitAVL()
	printTree(tree, t)
	node := tree.SearchMinElementInBalancedBinaryTree()
	t.Log(node)
}

func TestSearchMinElementInBalancedBinaryTree_1(t *testing.T) {
	tree := InitAVL()
	printTree(tree, t)
	node := tree.SearchMinElementInBalancedBinaryTree_1(tree.Root)
	t.Log(node)
}
