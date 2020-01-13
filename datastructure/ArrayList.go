package TestProject

const capacity  = 8

type ArrayList struct {
	length int
	data [capacity]interface{}
}

func checkInit(list *ArrayList){
	if list==nil{
		panic("未初始化列表")
	}
}

func (list *ArrayList) isEmpty() bool{
	checkInit(list)
	if list.length > 0 {
		return false
	}
	return true
}

func (list *ArrayList) add(value interface{}) bool{
	checkInit(list)
	if list.length+1>capacity{
		return false
	}
	list.data[list.length]=value
	list.length++
	return true
}

func (list *ArrayList) size() int{
	checkInit(list)
	return list.length
}

func(list *ArrayList) insert(index int,value interface{}) bool{
	checkInit(list)
	if list.length+1>capacity{
		return false
	}
	if index<0 || index>=capacity{
		return false
	}

	for i:=list.length;i>index;i--{
		list.data[i]=list.data[i-1]
	}
	list.data[index]=value
	list.length++
	return true
}

func (list *ArrayList) getValue(index int) interface{}{
	checkInit(list)
	if list.length==0{
		return nil
	}
	if index<0||index>=capacity{
		return nil
	}
	return list.data[index]
}

func (list *ArrayList) delete(value interface{}) bool{
	checkInit(list)
	if list.length==0{
		return false
	}
	index,flag:=list.getIndex(value)
	if flag{
		for i:=index+1;i<=list.length;i++{
			list.data[i-1]=list.data[i]
		}
		list.length--
		return true
	}
	return flag
}

func (list *ArrayList) getIndex(value interface{}) (int,bool){
	checkInit(list)
	if list.length==0{
		return 0,false
	}
	for i,v:=range list.data{
		if v==value{
			return i,true
		}
	}
	return 0,false
}