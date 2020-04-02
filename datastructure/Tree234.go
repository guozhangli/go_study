package TestProject

import "reflect"

type Tree234 struct {
	Root *Node234
}

const NUM4 = 4

type Node234 struct {
	Data      []Map
	ChildNode []*Node234
	Parent    *Node234
	ItemNum   int
}

func NewTree234() *Tree234 {
	return new(Tree234)
}

func (tree *Tree234) CreateNode234(_map Map) *Node234 {
	node := &Node234{
		Data:      make([]Map, NUM4-1),
		ChildNode: make([]*Node234, NUM4),
		Parent:    nil,
		ItemNum:   0,
	}
	node.Data[0] = _map
	node.ItemNum++
	return node
}

func (node *Node234) IsLeaf() bool {
	return node.ChildNode[0] == nil
}

func (node *Node234) IsFull() bool {
	return node.ItemNum == NUM4-1
}

func (node *Node234) GetChildNode(index int) *Node234 {
	return node.ChildNode[index]
}

func (node *Node234) GetDataNum() int {
	return node.ItemNum
}

func (node *Node234) GetParent() *Node234 {
	return node.Parent
}

func (node *Node234) GetNodeKey(index int) reflect.Value {
	k0 := reflect.ValueOf(node.Data[index]).MapKeys()[0]
	return k0
}
func (node *Node234) GetNodeValue(index int) interface{} {
	k0 := reflect.ValueOf(node.Data[index]).MapKeys()[0]
	return reflect.ValueOf(node.Data[index]).MapIndex(k0).Interface()
}

func (node *Node234) AddChildNode(child *Node234, index int) {
	if index != 0 && index != 1 && index != 2 && index != 3 {
		panic("input index is error")
	}
	node.ChildNode[index] = child
	if child != nil {
		child.Parent = node
	}
}

func (node *Node234) RemoveChildNode(index int) *Node234 {
	if index != 0 && index != 1 && index != 2 && index != 3 {
		panic("input index is error")
	}
	var temp = node.ChildNode[index]
	node.ChildNode[index] = nil
	return temp
}

func (node *Node234) GetChildNodeIndex(key string) int {
	var postion int
	for postion = 0; postion < node.GetDataNum(); postion++ {
		k1 := node.GetNodeKey(postion)
		if CompareTo(k1.Interface(), key) > 0 {
			break
		}
	}
	return postion
}

func (tree *Tree234) Find(key string) interface{} {
	node := tree.Root
	var curIndex int
	for {
		if curIndex = node.FindData(key); curIndex != -1 {
			return node.Data[curIndex][key]
		} else if node.IsLeaf() {
			break
		} else {
			node = node.GetNextChild(key)
		}
	}
	return nil
}

func (node *Node234) FindData(key string) int {
	for i, v := range node.Data {
		if v != nil {
			if _, ok := v[key]; ok {
				return i
			}
		}
	}
	return -1
}

func (node *Node234) GetNextChild(key string) *Node234 {
	for i := 0; i < node.ItemNum; i++ {
		ki := node.Data[i].GetMapKey()
		if CompareTo(ki.Interface(), key) > 0 {
			return node.GetChildNode(i)
		} else if CompareTo(ki.Interface(), key) == 0 {
			return node
		}
	}
	return node.GetChildNode(node.ItemNum)
}

func (node *Node234) InsertData(_map Map) {
	for i := NUM4 - 2; i >= 0; i-- {
		if node.Data[i] == nil {
			continue
		} else {
			ki := node.Data[i].GetMapKey()
			key := _map.GetMapKey()
			if CompareTo(ki.Interface(), key.Interface()) > 0 {
				node.Data[i+1] = node.Data[i]
			} else {
				node.Data[i+1] = _map
				node.ItemNum++
				return
			}
		}
	}
	node.Data[0] = _map
	node.ItemNum++
}

func (node *Node234) RemoveData(_map Map) {
	for i := NUM4 - 2; i >= 0; i-- {
		if node.Data[i] == nil {
			continue
		} else {
			ki := node.Data[i].GetMapKey()
			key := _map.GetMapKey()
			if CompareTo(ki.Interface(), key.Interface()) == 0 {
				if i == NUM4-2 {
					node.Data[i] = nil
				} else {
					node.Data[i] = node.Data[i+1]
					node.Data[i+1] = nil
				}
				node.ItemNum--
			}
		}
	}
}

func (tree *Tree234) Insert(key string, value interface{}) {
	node := tree.Root
	if node == nil {
		tree.Root = tree.CreateNode234(CreateMap(key, value))
	} else {
		for {
			if node.IsFull() {
				tree.Split(node)
				node = node.Parent
			}
			if node.IsLeaf() {
				break
			}
			node = node.GetNextChild(key)
		}
		node.InsertData(CreateMap(key, value))
	}
}

func (tree *Tree234) Split(node *Node234) {
	newBorther := tree.CreateNode234(node.Data[2])
	newBorther.AddChildNode(node.GetChildNode(2), 0)
	newBorther.AddChildNode(node.GetChildNode(3), 1)
	parent := &node.Parent
	if *parent == nil {
		newRoot := tree.CreateNode234(node.Data[1])
		newRoot.AddChildNode(node, 0)
		newRoot.AddChildNode(newBorther, 1)
		tree.Root = newRoot
	} else {
		(*parent).InsertData(node.Data[1])
		index := (*parent).GetChildNodeIndex(node.Data[2].GetMapKey().String())
		for i := (*parent).GetDataNum(); i >= index; i-- {
			(*parent).ChildNode[i] = (*parent).ChildNode[i-1]
		}
		(*parent).AddChildNode(newBorther, index)

	}
	node.RemoveData(node.Data[2])
	node.RemoveData(node.Data[1])
	node.RemoveChildNode(2)
	node.RemoveChildNode(3)
}

func (tree *Tree234) MidOrder(node *Node234, queue *QueueList) {
	if node != nil {
		tree.MidOrder(node.GetChildNode(0), queue)
		queue.EnQueue(node.Data[0])
		tree.MidOrder(node.GetChildNode(1), queue)
		if node.Data[1] != nil {
			queue.EnQueue(node.Data[1])
		}
		tree.MidOrder(node.GetChildNode(2), queue)
		if node.Data[2] != nil {
			queue.EnQueue(node.Data[2])
		}
		tree.MidOrder(node.GetChildNode(3), queue)
	}
}

func (tree *Tree234) MidOrderAndPrint() {
	queue := NewQueueList().(*QueueList)
	tree.MidOrder(tree.Root, queue)
	queue.Print()
}
