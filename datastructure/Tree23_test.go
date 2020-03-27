package TestProject

import (
	"encoding/json"
	"testing"
)

func print(tree *Tree23, t *testing.T) {
	str, _ := json.Marshal(tree)
	t.Log(string(str))
}

func TestInitTree23(t *testing.T) {
	tree := InitTree23()
	print(tree, t)
}

func TestCompareTo(t *testing.T) {
	v := CompareTo("a", "b")
	t.Log(v)
	v1 := CompareTo(10, 20)
	t.Log(v1)
}

func TestFindNode(t *testing.T) {
	/*tree:=InitTree23()
	node:=tree.Find("A")
	t.Log(node)*/
}

func TestMap(t *testing.T) {
	m := make(map[interface{}]interface{})
	m["ss"] = "wwww"
	m["bb"] = "ssss"
	for k, v := range m {
		if k == "ss" {
			t.Log(v)
		}
	}

}
