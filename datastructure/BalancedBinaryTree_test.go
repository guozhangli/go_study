package TestProject

import "testing"

func TestInsertNodeInBalancedBinaryTree(t *testing.T) {
	node := [10]int{3, 2, 1, 4, 5, 6, 7, 10, 9, 8}
	tree := InitBalancedBinaryTree()
	for _, v := range node {
		tree.InsertNodeInBalancedBinaryTree(v)
	}
	t.Log(tree)
}
