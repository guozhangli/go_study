package TestProject

import (
	"encoding/json"
	"testing"
)

func TestBinaryTree_Create(t *testing.T) {
	var node = NewBinaryTree("A")
	var str = []string{"B", "#", "#", "C", "#", "D", "#", "#"}
	CreateBinaryTree(node, str, 0)
	str1, _ := json.Marshal(node)
	t.Log(string(str1))
}
