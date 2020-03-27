package TestProject

import (
	"reflect"
)

//2-3树
type Tree23 struct {
	Root *Node23
}

const NUM = 3

type Node23 struct {
	Data       []map[interface{}]interface{}
	ChirdNode  []*Node23
	ParentNode *Node23
}

func InitTree23() *Tree23 {
	tree := new(Tree23)
	tree.Root = tree.CreateNode23()
	return tree
}

func CompareTo(val1 interface{}, val2 interface{}) int {
	if v1, ok := val1.(int); ok {
		if v2, ok := val2.(int); ok {
			if v1 > v2 {
				return 1
			} else if v1 < v2 {
				return -1
			} else {
				return 0
			}
		}
	}
	if v1, ok := val1.(string); ok {
		if v2, ok := val2.(string); ok {
			if v1 > v2 {
				return 1
			} else if v1 < v2 {
				return -1
			} else {
				return 0
			}
		}
	}
	panic("compare error")
}

func (tree *Tree23) CreateNode23() *Node23 {
	return &Node23{
		Data:       make([]map[interface{}]interface{}, NUM-1),
		ChirdNode:  make([]*Node23, NUM),
		ParentNode: nil,
	}
}

func (node *Node23) IsLeaf() bool {
	return node.ChirdNode[0] == nil
}

func (node *Node23) GetParent() *Node23 {
	return node.ParentNode
}

func (node *Node23) IsFull() bool {
	var count int
	for _, v := range node.Data {
		if v != nil {
			count++
		}
	}
	return count == NUM-1
}

func (node *Node23) IsEmpty() bool {
	var count int
	for _, v := range node.Data {
		if v == nil {
			count++
		}
	}
	return count == NUM-1
}

func (node *Node23) GetItemNum() int {
	var count int
	for _, v := range node.Data {
		if v != nil {
			count++
		}
	}
	return count
}

func (node *Node23) GetData(index int) interface{} {
	if index != 0 && index != 1 {
		panic("input index is error")
	}
	return node.Data[index]
}

//index =0 -->左, index=1 -->中, index=2 -->右
func (node *Node23) GetChildNode(index int) *Node23 {
	if index != 0 && index != 1 && index != 2 {
		panic("input index is error")
	}
	return node.ChirdNode[index]
}

//index =0 -->左, index=1 -->中, index=2 -->右
func (node *Node23) ConnectChildNode(childNode *Node23, index int) {
	if index != 0 && index != 1 && index != 2 {
		panic("input index is error")
	}
	node.ChirdNode[index] = childNode
	if childNode != nil {
		childNode.ParentNode = node
	}
}

//index =0 -->左, index=1 -->中, index=2 -->右
func (node *Node23) DisConnectChildNode(index int) *Node23 {
	if index != 0 && index != 1 && index != 2 {
		panic("input index is error")
	}
	var temp = node.ChirdNode[index]
	node.ChirdNode[index] = nil
	return temp
}

func (node *Node23) FindData(key interface{}) int {
	for i, v := range node.Data {
		if _, ok := v[key]; ok {
			return i
		}
	}
	return -1
}

func (node *Node23) InsertData(key interface{}, value interface{}) int {
	for i := NUM - 2; i >= 0; i-- {
		if node.Data[i] == nil {
			continue
		} else {
			v1 := reflect.ValueOf(node.Data[i]).MapKeys()[0]
			if CompareTo(v1, value) > 0 {
				node.Data[i+1] = node.Data[i]
			} else {
				m := make(map[interface{}]interface{})
				m[key] = value
				node.Data[i+1] = m
				return i + 1
			}
		}
	}
	m := make(map[interface{}]interface{})
	m[key] = value
	node.Data[0] = m
	return 0
}

func (node *Node23) RemoveData() interface{} {
	index := node.GetItemNum()
	temp := node.Data[index]
	node.Data[index] = nil
	return temp
}

func (tree *Tree23) Find(key interface{}) interface{} {
	node := tree.Root
	var curIndex int
	for {
		if curIndex = node.FindData(key); curIndex != -1 {
			return node.Data[curIndex][key]
		} else if node.IsLeaf() {
			return nil
		} else {
			node = GetNextChild(node, key)
		}
	}
}

func GetNextChild(node *Node23, key interface{}) *Node23 {
	for i := 0; i < node.GetItemNum(); i++ {
		k := reflect.ValueOf(node.Data[i]).MapKeys()[0]
		if CompareTo(k, key) > 0 {
			return node.GetChildNode(i)
		} else if CompareTo(k, key) == 0 {
			return node
		}
	}
	return node.GetChildNode(node.GetItemNum())
}

func (tree *Tree23) Insert(key interface{}, value interface{}) {
	node := tree.Root
	for {
		if node.IsLeaf() {
			break
		} else {
			node = GetNextChild(node, key)
			for _, v := range node.Data {
				k := reflect.ValueOf(v).MapKeys()[0]
				if CompareTo(k, key) == 0 {
					v[k] = value
					return
				}
			}
		}
	}
	if node.IsFull() {
		node.Split(key, value)
	} else {
		node.InsertData(key, value)
	}
}

func (node *Node23) Split(key interface{}, value interface{}) {
	parent := node.ParentNode
	if parent == nil {

	} else {
		if parent.GetItemNum() == 1 {

		} else if parent.GetItemNum() == 2 {

		}
	}
}
