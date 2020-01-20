package TestProject

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestBinaryTree_Create(t *testing.T) {
	var tree = new(BinaryTree)
	var str = []string{"A", "B", "#", "#", "C", "#", "D", "#", "#"}
	tree.CreatePerBinaryTree(nil, str)
	str1, _ := json.Marshal(tree)
	t.Log(string(str1))
}

func TestBinaryTree_CreateMid(t *testing.T) {
	var tree = new(BinaryTree)
	var str = []string{"#", "B", "#", "A", "#", "C", "#", "D", "#"}
	n := tree.CreateMidBinaryTree(nil, str)
	str2, _ := json.Marshal(n)
	t.Log(string(str2))
	str1, _ := json.Marshal(tree)
	t.Log(string(str1))
}

func TestPointer(t *testing.T) {
	var index *int
	var a = 1
	index = &(a)
	fmt.Println(*index)
}
