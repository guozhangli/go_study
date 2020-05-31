package TestProject

import "fmt"

/**
循环链表
	使用循环链表的例子：第二次机会页面置换算法
		第二次机会置换算法式实现LRU页面置换法的一种方式。它的工作方式是维护一个当前存在于物理内存中的页面的循环链表。
		为了简化说明，假设链表中的每个元素只存储一个页码和一个引用值，引用值要么是1要么是0.在实践中，每一个元素还包括
		其他信息。所有的页面初始引用值都设置为0，每当系统访问页面时，（比如，某个进程开始读或写某个页面），该页面的
		引用值就设置为1.
		当需要某个页帧时，操作系统就使用它维护的循环链表以及引用值来判断哪些页面应该释放其页帧。为了确定这一点，开始遍历
		链表，直到找到一个引用值为0的元素，当遍历每一个页面时，操作系统将页面的引用值从1重设回0.一旦它遇到引用值为0的元素，
		它就找到了一个自从上次遍历链表以来都没有被系统访问过的页面，因此这个页面就时最近最少使用的页面。那么这个页面
		就在物理内存中和新的页面置换，新的页面被插入道链表中原来的位置。如果自从算法上次运行以来，所有的页面都被访问过了，
		那么操作系统就完整的遍历了一次链表，此时就置换它开始的页面。

*/
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
