package TestProject

import "fmt"

type CyclicLinkedList struct {
	data     interface{}
	nextNode *CyclicLinkedList
}

func NewCyclicLinkedList(value interface{}) *CyclicLinkedList {
	cycLinkedList := &CyclicLinkedList{
		data:     value,
		nextNode: nil,
	}
	return cycLinkedList
}

func checkCyclicLinkedList(list *CyclicLinkedList) {
	if list == nil {
		panic("未创建循环链表")
	}
}

func (list *CyclicLinkedList) HeadAdd(value interface{}) {
	checkCyclicLinkedList(list)
	headNode := list
	node := NewCyclicLinkedList(value)
	if headNode.nextNode != nil {
		node.nextNode = headNode.nextNode
		headNode.nextNode = node
	} else {
		headNode.nextNode = node
		node.nextNode = headNode
	}
}

func (list *CyclicLinkedList) TailAdd(value interface{}) {
	checkCyclicLinkedList(list)
	headNode := list
	node := NewCyclicLinkedList(value)
	var currNode = headNode
	for currNode != nil && currNode.nextNode != headNode {
		currNode = currNode.nextNode
	}
	if currNode == nil {
		headNode.nextNode = node
		node.nextNode = headNode
	} else {
		currNode.nextNode = node
		node.nextNode = headNode
	}
}

func (list *CyclicLinkedList) Print(head *CyclicLinkedList) {
	checkCyclicLinkedList(list)
	cycLinkedList := list
	for cycLinkedList != nil && cycLinkedList.nextNode != head {
		fmt.Println(cycLinkedList)
		cycLinkedList = cycLinkedList.nextNode
	}
	if cycLinkedList != nil {
		fmt.Println(cycLinkedList)
	}
}

func (list *CyclicLinkedList) Insert(index int, value interface{}) bool {
	checkCyclicLinkedList(list)
	head := list
	cycLinkedList := list
	node := NewCyclicLinkedList(value)
	if index < 0 || index > list.Length() {
		return false
	}
	var count int
	var currNode *CyclicLinkedList
	for {
		if count == index {
			currNode = cycLinkedList
			break
		}
		cycLinkedList = cycLinkedList.nextNode
		count++
		if cycLinkedList == nil || cycLinkedList == head {
			break
		}
	}
	if currNode != nil && currNode != head {
		node.nextNode = currNode.nextNode
		currNode.nextNode = node
	} else {
		node.nextNode = head.nextNode
		head.nextNode = node

	}
	return true
}

func (list *CyclicLinkedList) Length() int {
	checkCyclicLinkedList(list)
	cycLinkedList := list
	head := list
	var count int
	for {
		cycLinkedList = cycLinkedList.nextNode
		if cycLinkedList == nil || cycLinkedList == head {
			break
		}
		count++
	}
	return count
}

func (list *CyclicLinkedList) Delete(value interface{}) {
	checkCyclicLinkedList(list)
	cycLinkedList := list
	head := list
	var currNode, perNode *CyclicLinkedList
	for {
		if cycLinkedList.data == value {
			currNode = cycLinkedList
			break
		}
		perNode = cycLinkedList
		cycLinkedList = cycLinkedList.nextNode
		if cycLinkedList == nil || cycLinkedList == head {
			break
		}
	}
	if currNode != nil {
		perNode.nextNode = currNode.nextNode
	}
}

func (list *CyclicLinkedList) DeleteF() {
	checkCyclicLinkedList(list)
	head := list
	if head.nextNode != nil {
		head.nextNode = head.nextNode.nextNode
	}
}

func (list *CyclicLinkedList) DeleteT() {
	checkCyclicLinkedList(list)
	cycLinkedList := list
	head := list
	var tail_1 *CyclicLinkedList
	for cycLinkedList.nextNode != head {
		tail_1 = cycLinkedList
		cycLinkedList = cycLinkedList.nextNode
	}
	if tail_1 != nil {
		tail_1.nextNode = head
	}
}

func (list *CyclicLinkedList) Clear() {
	checkCyclicLinkedList(list)
	cycLinkedList := list
	head := list
	var tail *CyclicLinkedList
	for cycLinkedList.nextNode != head {
		tail = cycLinkedList.nextNode
		cycLinkedList = nil
		cycLinkedList = tail
	}
	if tail != nil {
		head.nextNode = nil
		tail.nextNode = nil
	}

}

func (list *CyclicLinkedList) IsEmpty() bool {
	checkCyclicLinkedList(list)
	count := list.Length()
	if count > 0 {
		return false
	}
	return true
}
