package TestProject

import (
	"encoding/json"
	"testing"
)

func TestCreateBinarySearchTree(t *testing.T) {
	bst := NewBinarySearchTree()
	str_p := []string{"7", "4", "2", "5", "9", "10"} //先序
	str_m := []string{"2", "4", "5", "7", "9", "10"} //中序
	bst.InitBinarySearchTree(str_p, str_m)           //创建二叉搜索树跟元素顺序有关
	str, _ := json.Marshal(bst)
	t.Log(string(str))
}

func TestSearchMinElementInBinarySearchTree(t *testing.T) {
	bst := NewBinarySearchTree()
	str_p := []string{"7", "4", "2", "5", "9", "10"}
	str_m := []string{"2", "4", "5", "7", "9", "10"}
	bst.InitBinarySearchTree(str_p, str_m)
	str, _ := json.Marshal(bst)
	t.Log(string(str))
	min := bst.SearchMinElementInBinarySearchTree(bst.Root)
	t.Log(min)
}

func TestBinarySearchTree_SearchMinElementInBinarySearchTreeForeach(t *testing.T) {
	bst := NewBinarySearchTree()
	str_p := []string{"7", "4", "2", "5", "9", "10"}
	str_m := []string{"2", "4", "5", "7", "9", "10"}
	bst.InitBinarySearchTree(str_p, str_m)
	str, _ := json.Marshal(bst)
	t.Log(string(str))
	min := bst.SearchMinElementInBinarySearchTreeForeach(bst.Root)
	t.Log(min)
}

func TestSearchMaxElementInBinarySearchTree(t *testing.T) {
	bst := NewBinarySearchTree()
	str_p := []string{"7", "4", "2", "5", "9", "10"}
	str_m := []string{"2", "4", "5", "7", "9", "10"}
	bst.InitBinarySearchTree(str_p, str_m)
	str, _ := json.Marshal(bst)
	t.Log(string(str))
	max := bst.SearchMaxElementInBinarySearchTree(bst.Root)
	t.Log(max)
}

func TestBinarySearchTree_SearchMaxElementInBinarySearchTreeForeach(t *testing.T) {
	bst := NewBinarySearchTree()
	str_p := []string{"7", "4", "2", "5", "9", "10"}
	str_m := []string{"2", "4", "5", "7", "9", "10"}
	bst.InitBinarySearchTree(str_p, str_m)
	str, _ := json.Marshal(bst)
	t.Log(string(str))
	max := bst.SearchMaxElementInBinarySearchTreeForeach(bst.Root)
	t.Log(max)
}

func TestSearchElementInBinarySearchTree(t *testing.T) {
	bst := NewBinarySearchTree()
	str_p := []string{"7", "4", "2", "5", "9", "10"}
	str_m := []string{"2", "4", "5", "7", "9", "10"}
	bst.InitBinarySearchTree(str_p, str_m)
	str, _ := json.Marshal(bst)
	t.Log(string(str))
	ele := bst.SearchElementInBinarySearchTree(bst.Root, 9)
	t.Log(ele)
}

func TestBinarySearchTree_SearchElementInBinarySearchTreeForeach(t *testing.T) {
	bst := NewBinarySearchTree()
	str_p := []string{"7", "4", "2", "5", "9", "10"}
	str_m := []string{"2", "4", "5", "7", "9", "10"}
	bst.InitBinarySearchTree(str_p, str_m)
	str, _ := json.Marshal(bst)
	t.Log(string(str))
	ele := bst.SearchElementInBinarySearchTreeForeach(bst.Root, 5)
	t.Log(ele)
}

func TestBinarySearchTree_InsertElementInBinarySearchTree(t *testing.T) {
	bst := NewBinarySearchTree()
	str_p := []string{"7", "4", "2", "5", "9", "10"}
	str_m := []string{"2", "4", "5", "7", "9", "10"}
	bst.InitBinarySearchTree(str_p, str_m)
	str, _ := json.Marshal(bst)
	t.Log(string(str))
	bst.InsertElementInBinarySearchTree(bst.Root, 3)
	str1, _ := json.Marshal(bst)
	t.Log(string(str1))
}

func TestBinarySearchTree_DeleteElementInBinarySearchTree(t *testing.T) {
	bst := NewBinarySearchTree()
	str_p := []string{"7", "4", "2", "5", "9", "10"}
	str_m := []string{"2", "4", "5", "7", "9", "10"}
	bst.InitBinarySearchTree(str_p, str_m)
	bst.InsertElementInBinarySearchTree(bst.Root, 6)
	str, _ := json.Marshal(bst)
	t.Log(string(str))
	bst.DeleteElementInBinarySearchTree(&bst.Root, 4)
	str1, _ := json.Marshal(&bst)
	t.Log(string(str1))
}

func TestBinarySearchTree_GetHight(t *testing.T) {
	bst := NewBinarySearchTree()
	str_p := []string{"7", "4", "2", "5", "9", "10"}
	str_m := []string{"2", "4", "5", "7", "9", "10"}
	bst.InitBinarySearchTree(str_p, str_m)
	str, _ := json.Marshal(bst)
	t.Log(string(str))
	h := bst.GetHight()
	t.Log(h)
}
