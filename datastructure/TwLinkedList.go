package TestProject

type TwLinkedList struct {
	data interface{}
	preNode *TwLinkedList
	nextNode *TwLinkedList
}

func newTwLinkedList(vaule interface{}) *TwLinkedList{
	return &TwLinkedList{
		data:vaule,
		preNode:nil,
		nextNode:nil,
	}
}

func checkTwLinkedList(list *TwLinkedList){
	if list==nil{
		panic("未创建双向链表")
	}
}

func (list *TwLinkedList) addHead(value interface{}){
	checkTwLinkedList(list)
   	node:=newTwLinkedList(value)
	node.preNode=list
	node.nextNode=list.nextNode
	if(list.nextNode!=nil){
		list.nextNode.preNode=node
	}
    list.nextNode=node
}

func (list *TwLinkedList) addTail(value interface{}) {
	checkTwLinkedList(list)
	node:=newTwLinkedList(value)
	var p *TwLinkedList
	if list.nextNode==nil{
		p=list
	}else{
		for list.nextNode!=nil{
			list=list.nextNode
			p=list
		}
	}
	p.nextNode=node
	node.preNode=p
}

func (list *TwLinkedList) insert(index int ,value interface{}) bool{
	checkTwLinkedList(list)
	node:=newTwLinkedList(value)
	if index<0 || index>=list.length(){
		return false
	}
	var count int
	var p *TwLinkedList
	for list!=nil{
		if count==index{
			p=list
			break
		}
		list=list.nextNode
		count++
	}
	if p.nextNode!=nil{
		p.nextNode.preNode=node
	}
	node.nextNode=p.nextNode
	p.nextNode=node
	node.preNode=p
	return true
}

func (list *TwLinkedList) length() int {
	checkTwLinkedList(list)
	var count int
	for list!=nil{
		list=list.nextNode
		count++
	}
  	return count
}

func (list *TwLinkedList) isEmpty() bool {
	checkTwLinkedList(list)
	for list.nextNode!=nil{
		return false
	}
	return true
}

func (list *TwLinkedList) getValue(index int) interface{} {
	checkTwLinkedList(list)
	if index<0 || index>=list.length(){
		return nil
	}
	var count int
	for list!=nil{
		if count==index {
			return list.data
		}
		list=list.nextNode
		count++
	}
	return nil
}

func (list *TwLinkedList) delete(value interface{}) {
	checkTwLinkedList(list)
	var p *TwLinkedList
	for list!=nil{
		if list.data==value{
			p=list
		}
		list=list.nextNode
	}
	if p!=nil&&p.nextNode!=nil{
		p.nextNode.preNode=p.preNode
		p.preNode.nextNode=p.nextNode
	}
	if p!=nil&&p.nextNode==nil{
		p.preNode.nextNode=nil
	}
	p=nil
}

func (list *TwLinkedList) deleteF(){
	checkTwLinkedList(list)
	var p *TwLinkedList
	if list.nextNode!=nil{
		p=list.nextNode
		if p.nextNode!=nil{
			p.nextNode.preNode=p.preNode
			p.preNode.nextNode=p.nextNode
		}
		if p.nextNode==nil{
			p.preNode.nextNode=nil
		}
	}
}

func (list *TwLinkedList) deleteT(){
	checkTwLinkedList(list)
	var p *TwLinkedList
	for list.nextNode!=nil{
		list=list.nextNode
		p=list
	}
	if p!=nil&&p.nextNode==nil{
		p.preNode.nextNode=nil
	}
	p=nil
}

func (list *TwLinkedList) clear(){
	checkTwLinkedList(list)
	list.nextNode=nil
}

