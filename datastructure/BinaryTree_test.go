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



func TestPointer(t *testing.T) {
	var index *int
	var a = 1
	index = &(a)
	fmt.Println(*index)
}
