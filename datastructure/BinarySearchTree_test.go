package TestProject

import (
	"encoding/json"
	"testing"
)

func TestCreateBinarySearchTree(t *testing.T) {
	bst := NewBinarySearchTree()
	str_p := []string{"7", "4", "2", "5", "9", "10"}
	str_m := []string{"2", "4", "5", "7", "9", "10"}
	bst.InitBinarySearchTree(str_p, str_m)
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