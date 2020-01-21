package TestProject

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestBinaryTree_Create_Pre(t *testing.T) {
	var tree = new(BinaryTree)
	var str = []string{"A", "B", "#", "#", "C", "#", "D", "#", "#"}
	tree.CreatePreBinaryTree(nil, str)
	str1, _ := json.Marshal(tree)
	t.Log(string(str1))
}

func TestBinaryTree_Create_Mid(t *testing.T) {
	var tree = new(BinaryTree)
	var str_p = []string{"A", "B", "#", "#", "C", "#", "D", "#", "#"}
	var str_m = []string{"#", "B", "#", "A", "#", "C", "#", "D", "#"}
	//var str = []string{"B", "A", "C"}
	tree.CreateMidBinaryTree(str_p, str_m)
	str1, _ := json.Marshal(tree)
	t.Log(string(str1))
}

func TestPointer(t *testing.T) {
	var index *int
	var a = 1
	index = &(a)
	fmt.Println(*index)
}
