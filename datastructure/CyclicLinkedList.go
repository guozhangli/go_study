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
	cycLinkedList := list
	for cycLinkedList != nil && cycLinkedList.nextNode != head {
		fmt.Println(cycLinkedList)
		cycLinkedList = cycLinkedList.nextNode
	}
	fmt.Println(cycLinkedList)
}

func (list *CyclicLinkedList) insert(index int, value interface{}) bool {
	head := list
	cycLinkedList := list
	node := NewCyclicLinkedList(value)
	if index < 0 || index > list.length() {
		return false
	}
	var count int
	var currNode *CyclicLinkedList
	for cycLinkedList != nil && cycLinkedList.nextNode != head {
		if count == index {
			currNode = cycLinkedList
			break
		}
		cycLinkedList = cycLinkedList.nextNode
		count++
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

func (list *CyclicLinkedList) length() int {
	cycLinkedList := list
	head := list
	var count int
	for cycLinkedList != nil && cycLinkedList.nextNode != head {
		cycLinkedList = cycLinkedList.nextNode
		count++
	}
	return count
}
