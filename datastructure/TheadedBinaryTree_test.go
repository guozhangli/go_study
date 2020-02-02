package TestProject

import (
	"encoding/json"
	"testing"
)

func TestThreadedBinaryTree_Created(t *testing.T) {
	var tbt = NewThreadedBinaryTree()
	t.Log(tbt)
}

func TestThreadedBinaryTree_CreatedMidOrder(t *testing.T) {
	var tbt = NewThreadedBinaryTree()
	var str_p = []string{"A", "B", "#", "#", "C", "#", "D", "#", "#"}
	var str_m = []string{"#", "B", "#", "A", "#", "C", "#", "D", "#"}
	tbt.CreateMidThreadedBinaryTree(str_p, str_m)
	str1, _ := json.Marshal(tbt)
	t.Log(string(str1))
}

func TestThreadedBinaryTreeThreading(t *testing.T) {
	var tbt = NewThreadedBinaryTree()
	var str_p = []string{"A", "B", "#", "#", "C", "#", "D", "#", "#"}
	var str_m = []string{"#", "B", "#", "A", "#", "C", "#", "D", "#"}
	tbt.CreateMidThreadedBinaryTree(str_p, str_m)
	str0, _ := json.Marshal(tbt)
	t.Log(string(str0))
	tbt.ThreadedBinaryTreeThreading(tbt.Root)

}

func TestThreadedBinaryTreeThreadingMidForEach(t *testing.T) {
	var tbt = NewThreadedBinaryTree()
	var str_p = []string{"A", "B", "#", "#", "C", "#", "D", "#", "#"}
	var str_m = []string{"#", "B", "#", "A", "#", "C", "#", "D", "#"}
	tbt.CreateMidThreadedBinaryTree(str_p, str_m)
	str0, _ := json.Marshal(tbt)
	t.Log(string(str0))
	tbt.ThreadedBinaryTreeThreading(tbt.Root)
	tbt.ThreadedBinaryTreeMidForEach()
}

func TestThreadedBinaryTreeThreadingPreForEach(t *testing.T) {
	var tbt = NewThreadedBinaryTree()
	var str_p = []string{"A", "B", "#", "#", "C", "#", "D", "#", "#"}
	var str_m = []string{"#", "B", "#", "A", "#", "C", "#", "D", "#"}
	tbt.CreateMidThreadedBinaryTree(str_p, str_m)
	str0, _ := json.Marshal(tbt)
	t.Log(string(str0))
	tbt.ThreadedBinaryTreeThreading(tbt.Root)
	tbt.ThreadedBinaryTreePreForEach()
}
