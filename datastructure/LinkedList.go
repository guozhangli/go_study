package TestProject

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
