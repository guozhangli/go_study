package TestProject

type LinkedList struct {
	data interface{}
	nextNode *LinkedList
}

func newLinkedList() *LinkedList{
	var node=&LinkedList{
		data:nil,
		nextNode:nil,
	}
	return node
}

func checkInitLinkedList(list *LinkedList){
	if list==nil{
		panic("未创建单链表")
	}
}
func (list *LinkedList) isEmpty() bool{
	checkInitLinkedList(list)
	r:=list.nextNode
	if r!=nil{
		return true
	}
	return false
}

func (list *LinkedList) tailAdd(value interface{}){
	checkInitLinkedList(list)
	r:=list
	var node=&LinkedList{
		data:value,
		nextNode:nil,
	}
	if r!=nil{
		for ;r.nextNode!=nil;{
			r=r.nextNode
		}
	}
	r.nextNode=node
}

func (list *LinkedList) headAdd(value interface{}) {
	checkInitLinkedList(list)
	f:=list
	var node=&LinkedList{
		data:value,
		nextNode:nil,
	}
	n:=f.nextNode
	node.nextNode=n
	f.nextNode=node
}

func (list *LinkedList) length() int {
	checkInitLinkedList(list)
	f := list
	var count int
	if f != nil {
		for ;f.nextNode != nil; {
			f=f.nextNode
			count++
		}
	}
	return count
}

func (list *LinkedList) getVaule(index int) interface{}{
	checkInitLinkedList(list)
    return nil
}