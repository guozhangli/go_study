package TestProject

import (
	"encoding/json"
	"reflect"
	"testing"
)

func printTree23(tree *Tree23, t *testing.T) {
	str, err := json.Marshal(tree)
	if err != nil {
		t.Log(err)
	}
	t.Log(string(str))
}

func TestInitTree23(t *testing.T) {
	tree := InitTree23()
	printTree23(tree, t)
}

func TestCompareTo(t *testing.T) {
	v := CompareTo("a", "b")
	t.Log(v)
	v1 := CompareTo(10, 20)
	t.Log(v1)
}

func TestMap(t *testing.T) {
	m := make(map[interface{}]interface{})
	m["ss"] = "wwww"
	m["bb"] = "ssss"

	v := reflect.ValueOf(m).MapKeys()
	t.Log(v[0])
	for _, v1 := range v {
		t.Log(v1)
	}
	for k, v := range m {
		if k == "ss" {
			t.Log(v)
		}
	}
}

func TestFindNode(t *testing.T) {
	tree := InitTree23()
	tree.Insert("A", "aaaaa")
	tree.Insert("B", "bbbbb")
	tree.Insert("C", "ccccc")
	tree.Insert("D", "ddddd")
	tree.Insert("E", "eeeee")
	node := tree.Find("E")
	t.Log(node)
}

func TestTree23Insert(t *testing.T) {
	tree := InitTree23()
	tree.Insert("A", "aaaaa")
	tree.Insert("B", "bbbbb")
	tree.Insert("C", "ccccc")
	tree.Insert("D", "ddddd")
	tree.Insert("E", "eeeee")
	tree.Insert("F", "fffff")
	tree.Insert("G", "ggggg")
	t.Log(tree)
	//printTree23(tree,t)  //嵌套调用，死循环
}
