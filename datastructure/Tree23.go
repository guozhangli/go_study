package TestProject

import (
	"reflect"
)

//2-3树
type Tree23 struct {
	Root *Node23
}

const NUM = 3

type Map map[string]interface{}

type Node23 struct {
	Data       []Map
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
		Data:       make([]Map, NUM-1),
		ChirdNode:  make([]*Node23, NUM),
		ParentNode: nil,
	}
}

func CreateNode23HasValue(key string, value interface{}) *Node23 {
	node := CreateNode23()
	m := make(map[string]interface{})
	m[key] = value
	node.Data[0] = m
	return node
}

func CreateNode23HasMap(m map[string]interface{}) *Node23 {
	node := CreateNode23()
	node.Data[0] = m
	return node
}

func CreateMap(key string, value interface{}) Map {
	m := make(map[string]interface{})
	m[key] = value
	return m
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

func (node *Node23) GetData(index int) Map {
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
func (node *Node23) AddChildNode(childNode *Node23, index int) {
	if index != 0 && index != 1 && index != 2 {
		panic("input index is error")
	}
	node.ChirdNode[index] = childNode
	if childNode != nil {
		childNode.ParentNode = node
	}
}

//index =0 -->左, index=1 -->中, index=2 -->右
func (node *Node23) RemoveChildNode(index int) *Node23 {
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
			if _, ok := v[key]; ok {
				return i
			}
		}
	}
	return -1
}

func (node *Node23) InsertData(m Map) int {
	for i := NUM - 2; i >= 0; i-- {
		if node.Data[i] == nil {
			continue
		} else {
			ki := node.GetNodeKey(i)
			k := m.GetMapKey()
			if CompareTo(ki.Interface(), k.Interface()) > 0 {
				node.Data[i+1] = node.Data[i]
			} else {
				node.Data[i+1] = m
				return i + 1
			}
		}
	}
	node.Data[0] = m
	return 0
}

func (node *Node23) RemoveData(m Map) {
	for i := NUM - 2; i >= 0; i-- {
		if node.Data[i] == nil {
			continue
		} else {
			ki := node.GetNodeKey(i)
			k := m.GetMapKey()
			if CompareTo(ki.Interface(), k.Interface()) == 0 {
				node.Data[i] = nil
				if i == 0 {
					node.Data[i] = node.Data[i+1]
					node.Data[i+1] = nil
				}
			}
		}
	}
}

func (node *Node23) RemoveTailData() Map {
	index := node.GetItemNum()
	temp := node.Data[index-1]
	node.Data[index-1] = nil
	return temp
}

func (node *Node23) GetNodeKey(index int) reflect.Value {
	k0 := reflect.ValueOf(node.Data[index]).MapKeys()[0]
	return k0
}

func (_map Map) GetMapKey() reflect.Value {
	k0 := reflect.ValueOf(_map).MapKeys()[0]
	return k0
}

func (node *Node23) GetNodeValue(index int) interface{} {
	k0 := reflect.ValueOf(node.Data[index]).MapKeys()[0]
	return reflect.ValueOf(node.Data[index]).MapIndex(k0).Interface()
}

func (tree *Tree23) Find(key string) interface{} {
	node := tree.Root
	var curIndex int
	for {
		if curIndex = node.FindData(key); curIndex != -1 {
			return node.Data[curIndex][key]
		} else if node.IsLeaf() {
			return nil
		} else {
			node = node.GetNextChild(key)
		}
	}
}

func (node *Node23) GetNextChild(key string) *Node23 {
	for i := 0; i < node.GetItemNum(); i++ {
		k := node.GetNodeKey(i)
		if CompareTo(k.Interface(), key) > 0 {
			return node.GetChildNode(i)
		} else if CompareTo(k.Interface(), key) == 0 {
			return node
		}
	}
	return node.GetChildNode(node.GetItemNum())
}

func (node *Node23) GetNextIndex(key string) int {
	var postion int
	for postion = 0; postion < node.GetItemNum(); postion++ {
		k1 := node.GetNodeKey(postion)
		if CompareTo(k1.Interface(), key) > 0 {
			break
		}
	}
	return postion
}

func (tree *Tree23) Insert(key string, value interface{}) {
	node := tree.Root
	for {
		if node.IsLeaf() {
			break
		} else {
			node = node.GetNextChild(key)
			for _, v := range node.Data {
				if v != nil {
					k := v.GetMapKey()
					if CompareTo(k.Interface(), key) == 0 {
						v[k.String()] = value
						return
					}
				}
			}
			if node.IsFull() {
				tree.Split(node, CreateMap(key, value))
			}
		}
	}
	node.InsertData(CreateMap(key, value))
}

func (tree *Tree23) Split(node *Node23, _map Map) *Node23 {
	k0 := node.GetNodeKey(0)
	k1 := node.GetNodeKey(1)
	key := _map.GetMapKey()
	var mid map[string]interface{}
	right := node.Data[1]
	if CompareTo(k0.Interface(), key.Interface()) > 0 {
		mid = node.Data[0]
		node.RemoveData(mid)
		node.InsertData(_map)
	} else if CompareTo(k1.Interface(), key.Interface()) < 0 {
		mid = node.Data[1]
		right = _map
	} else {
		mid = _map
	}
	node.RemoveTailData()
	newBorther := CreateNode23HasMap(right)
	parent := &node.ParentNode
	if *parent == nil {
		newParent := CreateNode23HasMap(mid)
		newParent.AddChildNode(node, 0)
		newParent.AddChildNode(newBorther, 1)
		tree.Root = newParent
	} else {
		index := (*parent).GetNextIndex(key.String())
		if (*parent).GetItemNum() == 1 {
			(*parent).InsertData(mid)
			if index == 0 {
				(*parent).AddChildNode((*parent).GetChildNode(1), 2)
				(*parent).AddChildNode(newBorther, 1)
			} else if index == 1 {
				(*parent).AddChildNode(newBorther, 2)
			}
		} else if (*parent).GetItemNum() == 2 {
			parentBorther := tree.Split(*parent, mid)
			if index == 0 { //当插入位置为0时,0位置原子节点和新增子节点交给原父节点，1位置和2位置子节点交给新增父节点
				parentBorther.AddChildNode((*parent).GetChildNode(1), 0)
				parentBorther.AddChildNode((*parent).GetChildNode(2), 1)
				(*parent).AddChildNode(newBorther, 1)
				(*parent).RemoveChildNode(2)
			} else if index == 1 { //当插入位置为1时,0位置和1位置原子节点交给原父节点，1位置新增子节点和2位置子节点交给新增父节点；
				parentBorther.AddChildNode(newBorther, 0)
				parentBorther.AddChildNode((*parent).GetChildNode(2), 1)
				(*parent).RemoveChildNode(2)
			} else { //当插入位置为2时,0位置和1位置子节点交给原父节点，2位置原子节点和新增子节点交给新增父节点
				parentBorther.AddChildNode((*parent).GetChildNode(2), 0)
				parentBorther.AddChildNode(newBorther, 1)
				(*parent).RemoveChildNode(2)
			}
		}
	}
	return newBorther
}

func (tree *Tree23) MidOrderAndPrint() {
	queue := NewQueueList().(*QueueList)
	tree.MidOrder(tree.Root, queue)
	queue.Print()
}

//2-3树中序遍历（递归）
func (tree *Tree23) MidOrder(node *Node23, queue *QueueList) {
	if node != nil {
		tree.MidOrder(node.GetChildNode(0), queue)
		queue.EnQueue(node.Data[0])
		tree.MidOrder(node.GetChildNode(1), queue)
		if node.Data[1] != nil {
			queue.EnQueue(node.Data[1])
		}
		tree.MidOrder(node.GetChildNode(2), queue)
	}
}
