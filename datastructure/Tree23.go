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
	Data       []*map[string]interface{}
	ChirdNode  []*Node23
	ParentNode *Node23
}

func InitTree23() *Tree23 {
	tree := new(Tree23)
	tree.Root = CreateNode23()
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

func CreateNode23() *Node23 {
	return &Node23{
		Data:       make([]*map[string]interface{}, NUM-1),
		ChirdNode:  make([]*Node23, NUM),
		ParentNode: nil,
	}
}

func CreateNode23HasValue(key string, value interface{}) *Node23 {
	node := CreateNode23()
	m := make(map[string]interface{})
	m[key] = value
	node.Data[0] = &m
	return node
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

func (node *Node23) GetData(index int) *map[string]interface{} {
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

func (node *Node23) FindData(key string) int {
	for i, v := range node.Data {
		if v != nil {
			if _, ok := (*v)[key]; ok {
				return i
			}
		}
	}
	return -1
}

func (node *Node23) InsertData(key string, value interface{}) int {
	for i := NUM - 2; i >= 0; i-- {
		if node.Data[i] == nil {
			continue
		} else {
			v1 := reflect.ValueOf(*node.Data[i]).MapKeys()[0]
			if CompareTo(v1.Interface(), key) > 0 {
				node.Data[i+1] = node.Data[i]
			} else {
				m := make(map[string]interface{})
				m[key] = value
				node.Data[i+1] = &m
				return i + 1
			}
		}
	}
	m := make(map[string]interface{})
	m[key] = value
	node.Data[0] = &m
	return 0
}

func (node *Node23) RemoveData() *map[string]interface{} {
	index := node.GetItemNum()
	temp := node.Data[index-1]
	node.Data[index-1] = nil
	return temp
}

func (tree *Tree23) Find(key string) interface{} {
	node := tree.Root
	var curIndex int
	for {
		if curIndex = node.FindData(key); curIndex != -1 {
			return (*(node.Data[curIndex]))[key]
		} else if node.IsLeaf() {
			return nil
		} else {
			node = GetNextChild(node, key)
		}
	}
}

func GetNextChild(node *Node23, key string) *Node23 {
	for i := 0; i < node.GetItemNum(); i++ {
		k := reflect.ValueOf(*node.Data[i]).MapKeys()[0]
		if CompareTo(k.Interface(), key) > 0 {
			return node.GetChildNode(i)
		} else if CompareTo(k.Interface(), key) == 0 {
			return node
		}
	}
	return node.GetChildNode(node.GetItemNum())
}

func (tree *Tree23) Insert(key string, value interface{}) {
	node := tree.Root
	for {
		if node.IsLeaf() {
			break
		} else {
			node = GetNextChild(node, key)
			for _, v := range node.Data {
				if v != nil {
					k := reflect.ValueOf(*v).MapKeys()[0]
					if CompareTo(k.Interface(), key) == 0 {
						(*v)[k.String()] = value
						return
					}
				}
			}
		}
	}
	if node.IsFull() {
		tree.Split(&node, key, value)
	} else {
		node.InsertData(key, value)
	}
}

func (tree *Tree23) Split(node **Node23, key string, value interface{}) {
	parent := (*node).ParentNode
	if parent == nil {
		tree.NoParent(node, key, value)
		tree.Root = *node
	} else {
		if parent.GetItemNum() == 1 {
			tree.OneParent(node, key, value)
		} else if parent.GetItemNum() == 2 {
			tree.TwoParent(node, key, value)
		}
	}
}

func (tree *Tree23) NoParent(node **Node23, key string, value interface{}) {
	var leftNode, midNode, rightNode *Node23
	k0 := reflect.ValueOf(*(*node).Data[0]).MapKeys()[0]
	k1 := reflect.ValueOf(*(*node).Data[1]).MapKeys()[0]
	if CompareTo(k0.Interface(), key) > 0 {
		midNode = CreateNode23HasValue(k0.String(), (*(*node).Data[0])[k0.String()])
		leftNode = CreateNode23HasValue(key, value)
		rightNode = CreateNode23HasValue(k1.String(), (*(*node).Data[1])[k1.String()])
		leftNode.ChirdNode[0] = (*node).ChirdNode[0]
		rightNode.ChirdNode[0] = (*node).ChirdNode[1]
		rightNode.ChirdNode[1] = (*node).ChirdNode[2]
		midNode.ChirdNode[0] = leftNode
		midNode.ChirdNode[1] = rightNode
		leftNode.ParentNode = midNode
		rightNode.ParentNode = midNode
	} else if CompareTo(k1.Interface(), key) < 0 {
		midNode = CreateNode23HasValue(k1.String(), (*(*node).Data[1])[k1.String()])
		leftNode = CreateNode23HasValue(k0.String(), (*(*node).Data[0])[k0.String()])
		rightNode = CreateNode23HasValue(key, value)
		leftNode.ChirdNode[0] = (*node).ChirdNode[0]
		rightNode.ChirdNode[0] = (*node).ChirdNode[1]
		rightNode.ChirdNode[1] = (*node).ChirdNode[2]
		midNode.ChirdNode[0] = leftNode
		midNode.ChirdNode[1] = rightNode
		leftNode.ParentNode = midNode
		rightNode.ParentNode = midNode
	} else {
		midNode = CreateNode23HasValue(key, value)
		leftNode = CreateNode23HasValue(k0.String(), (*(*node).Data[0])[k0.String()])
		rightNode = CreateNode23HasValue(k1.String(), (*(*node).Data[1])[k1.String()])
		leftNode.ChirdNode[0] = (*node).ChirdNode[0]
		rightNode.ChirdNode[0] = (*node).ChirdNode[1]
		rightNode.ChirdNode[1] = (*node).ChirdNode[2]
		midNode.ChirdNode[0] = leftNode
		midNode.ChirdNode[1] = rightNode
		leftNode.ParentNode = midNode
		rightNode.ParentNode = midNode
	}
	*node = midNode
}

func (tree *Tree23) OneParent(node **Node23, key string, value interface{}) {
	var leftNode, rightNode *Node23
	parent := &(*node).ParentNode
	k0 := reflect.ValueOf(*(*node).Data[0]).MapKeys()[0]
	k1 := reflect.ValueOf(*(*node).Data[1]).MapKeys()[0]
	if CompareTo(k0.Interface(), key) > 0 {
		leftNode = CreateNode23HasValue(key, value)
		rightNode = CreateNode23HasValue(k1.String(), (*(*node).Data[1])[k1.String()])
		(*parent).Data[1] = (*node).Data[0]
		(*parent).ChirdNode[1] = leftNode
		(*parent).ChirdNode[2] = rightNode
		leftNode.ChirdNode[0] = (*node).ChirdNode[0]
		rightNode.ChirdNode[0] = (*node).ChirdNode[1]
		rightNode.ChirdNode[1] = (*node).ChirdNode[2]
		leftNode.ParentNode = *parent
		rightNode.ParentNode = *parent
	} else if CompareTo(k1.Interface(), key) < 0 {
		leftNode = CreateNode23HasValue(k0.String(), (*(*node).Data[0])[k0.String()])
		rightNode = CreateNode23HasValue(key, value)
		(*parent).Data[1] = (*node).Data[1]
		(*parent).ChirdNode[1] = leftNode
		(*parent).ChirdNode[2] = rightNode
		leftNode.ChirdNode[0] = (*node).ChirdNode[0]
		rightNode.ChirdNode[0] = (*node).ChirdNode[1]
		rightNode.ChirdNode[1] = (*node).ChirdNode[2]
		leftNode.ParentNode = *parent
		rightNode.ParentNode = *parent
	} else {
		leftNode = CreateNode23HasValue(k0.String(), (*(*node).Data[0])[k0.String()])
		rightNode = CreateNode23HasValue(k1.String(), (*(*node).Data[1])[k1.String()])
		m := make(map[string]interface{})
		m[key] = value
		(*parent).Data[1] = &m
		(*parent).ChirdNode[1] = leftNode
		(*parent).ChirdNode[2] = rightNode
		leftNode.ChirdNode[0] = (*node).ChirdNode[0]
		rightNode.ChirdNode[0] = (*node).ChirdNode[1]
		rightNode.ChirdNode[1] = (*node).ChirdNode[2]
		leftNode.ParentNode = *parent
		rightNode.ParentNode = *parent
	}
}

func (tree *Tree23) TwoParent(node **Node23, key string, value interface{}) {
	var leftNode, rightNode *Node23
	parent := &(*node).ParentNode
	k0 := reflect.ValueOf(*(*node).Data[0]).MapKeys()[0]
	k1 := reflect.ValueOf(*(*node).Data[1]).MapKeys()[0]
	if CompareTo(k0.Interface(), key) > 0 {
		leftNode = CreateNode23HasValue(key, value)
		rightNode = CreateNode23HasValue(k1.String(), (*(*node).Data[1])[k1.String()])

		(*parent).ChirdNode[0] = leftNode
		(*parent).ChirdNode[1] = rightNode
		leftNode.ChirdNode[0] = (*node).ChirdNode[0]
		rightNode.ChirdNode[0] = (*node).ChirdNode[1]
		rightNode.ChirdNode[1] = (*node).ChirdNode[2]
		leftNode.ParentNode = *parent
		rightNode.ParentNode = *parent
		tree.Split(parent, k0.String(), (*(*node).Data[0])[k0.String()])
	} else if CompareTo(k1.Interface(), key) < 0 {
		leftNode = CreateNode23HasValue(k0.String(), (*(*node).Data[0])[k0.String()])
		rightNode = CreateNode23HasValue(key, value)
		(*parent).ChirdNode[0] = leftNode
		(*parent).ChirdNode[1] = rightNode
		leftNode.ChirdNode[0] = (*node).ChirdNode[0]
		rightNode.ChirdNode[0] = (*node).ChirdNode[1]
		rightNode.ChirdNode[1] = (*node).ChirdNode[2]
		leftNode.ParentNode = *parent
		rightNode.ParentNode = *parent
		tree.Split(parent, k1.String(), (*(*node).Data[1])[k1.String()])
	} else {
		leftNode = CreateNode23HasValue(k0.String(), (*(*node).Data[0])[k0.String()])
		rightNode = CreateNode23HasValue(k1.String(), (*(*node).Data[1])[k1.String()])
		(*parent).ChirdNode[0] = leftNode
		(*parent).ChirdNode[1] = rightNode
		leftNode.ChirdNode[0] = (*node).ChirdNode[0]
		rightNode.ChirdNode[0] = (*node).ChirdNode[1]
		rightNode.ChirdNode[1] = (*node).ChirdNode[2]
		leftNode.ParentNode = *parent
		rightNode.ParentNode = *parent
		tree.Split(parent, key, value)

	}
}
