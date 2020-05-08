package TestProject

//链表

/*链表和数组的区别：
1)、空间：数组被分配一块连续的内存空间(大小固定)，
  动态数组在数组存储满了时，动态的调整大小（扩容为两倍的大小）当删除到数数据量小于一半的数组大小时，则把数组大小减小到一半。
  链表被分配不连续的内存空间，不需要做内存复制和重新分配的操作。
2)、时间：数组是随机存取的，当访问一个元素的时间开销为O(1),而链表在最差的情况下访问一个元素的开销为O(n).
 |___________参数___________|_____链表_____|_______________数组_______________|_____________________动态数组______________________|
 |___________索引___________|_____O(n)_____|_______________O(1)_______________|_____________________O(1)__________________________|
 |____在最前端插入或删除____|_____O(1)_____|_____如果数组空间没有填满O(n)_____|_____________________O(n)__________________________|
 |____在最末端插入__________|_____O(n)_____|_____如果数组空间没有填满O(1)_____|_如果数组空间没有填满O(1),如果数组空间已经填满O(n)_|
 |____在最末端删除__________|_____O(n)_____|_______________O(1)_______________|_____________________O(n)__________________________|
 |____在中间插入____________|_____O(n)_____|_____如果数组空间没有填满O(n)_____|_____________________O(n)__________________________|
 |____在中间删除____________|_____O(n)_____|_____如果数组空间没有填满O(n)_____|_____________________O(n)__________________________|
 |____空间浪费______________|_____O(n)_____|________________0_________________|_____________________O(n)__________________________|
*/
type LinkedList struct {
	data     interface{}
	nextNode *LinkedList
}

func newLinkedList() *LinkedList {
	var node = &LinkedList{
		data:     nil,
		nextNode: nil,
	}
	return node
}

func checkInitLinkedList(list *LinkedList) {
	if list == nil {
		panic("未创建单链表")
	}
}
func (list *LinkedList) IsEmpty() bool {
	checkInitLinkedList(list)
	r := list.nextNode
	if r != nil {
		return true
	}
	return false
}

func (list *LinkedList) TailAdd(value interface{}) {
	checkInitLinkedList(list)
	r := list
	var node = &LinkedList{
		data:     value,
		nextNode: nil,
	}
	if r != nil {
		for r.nextNode != nil {
			r = r.nextNode
		}
	}
	r.nextNode = node
}

func (list *LinkedList) HeadAdd(value interface{}) {
	checkInitLinkedList(list)
	f := list
	var node = &LinkedList{
		data:     value,
		nextNode: nil,
	}
	n := f.nextNode
	node.nextNode = n
	f.nextNode = node
}

func (list *LinkedList) Insert(index int, value interface{}) bool {
	checkInitLinkedList(list)
	f := list
	if index >= 0 && index <= f.Length() {
		var node = &LinkedList{
			data:     value,
			nextNode: nil,
		}
		var count int
		for f != nil {
			if count == index {
				node.nextNode = f.nextNode
				f.nextNode = node
				return true
			}
			f = f.nextNode
			count++
		}
	}
	return false
}

func (list *LinkedList) Length() int {
	checkInitLinkedList(list)
	f := list
	var count int
	if f != nil {
		for f.nextNode != nil {
			f = f.nextNode
			count++
		}
	}
	return count
}

func (list *LinkedList) GetVaule(index int) interface{} {
	checkInitLinkedList(list)
	f := list
	if index < 0 || index > f.Length() {
		return nil
	}
	var count int
	for {
		if count == index {
			return f.data
		}
		count++
		if f = f.nextNode; f == nil {
			break
		}
	}
	return nil
}

func (list *LinkedList) Delete(value interface{}) {
	checkInitLinkedList(list)
	p := list
	if value != nil {
		var q LinkedList
		for {
			if p.nextNode != nil && p.nextNode.data == value {
				q = *p.nextNode
				p.nextNode = q.nextNode
			}
			if p = p.nextNode; p == nil {
				break
			}
		}
	}
}

func (list *LinkedList) DeleteF() {
	checkInitLinkedList(list)
	f := list
	var p LinkedList
	if f.nextNode != nil {
		p = *f.nextNode
		f.nextNode = p.nextNode
	}
}

func (list *LinkedList) DeleteT() {
	checkInitLinkedList(list)
	f := list
	var p *LinkedList
	for f != nil {
		p = f.nextNode
		if p != nil && p.nextNode == nil {
			f.nextNode = nil
		}
		f = f.nextNode
		p = nil
	}
}

func (list *LinkedList) Clear() {
	checkInitLinkedList(list)
	p := list
	var it *LinkedList
	for p != nil {
		it = p.nextNode
		p = nil
		p = it
	}
	list.nextNode = nil
}
