package TestProject

import "errors"

/**
基于数组的固定长度的List
*/
type ArrayList struct {
	length int
	data   []interface{}
}

var capacity = func(list *ArrayList) int {
	return cap(list.data)
}

func NewArray(cap int) *ArrayList {
	if cap == 0 {
		return new(ArrayList)
	}
	list := &ArrayList{
		length: 0,
		data:   make([]interface{}, cap, cap)}
	return list
}

func checkInit(list *ArrayList) {
	if list == nil {
		panic("未初始化列表")
	}
}

func (list *ArrayList) IsEmpty() bool {
	checkInit(list)
	if list.length > 0 {
		return false
	}
	return true
}

func (list *ArrayList) Add(value interface{}) bool {
	checkInit(list)
	if list.length+1 > capacity(list) {
		return false
	}
	list.data[list.length] = value
	list.length++
	return true
}

func (list *ArrayList) Length() int {
	checkInit(list)
	return list.length
}

func (list *ArrayList) Insert(index int, value interface{}) bool {
	checkInit(list)
	if list.length+1 > capacity(list) {
		return false
	}
	if index >= capacity(list) {
		return false
	}

	for i := list.length; i > index; i-- {
		list.data[i] = list.data[i-1]
	}
	list.data[index] = value
	list.length++
	return true
}

func (list *ArrayList) GetValue(index int) interface{} {
	checkInit(list)
	if list.length == 0 {
		return nil
	}
	if index < 0 || index >= capacity(list) {
		return nil
	}
	return list.data[index]
}

func (list *ArrayList) Delete(value interface{}) bool {
	checkInit(list)
	if list.length == 0 {
		return false
	}
	index, err := list.GetIndex(value)
	if err != nil {
		return false
	}
	for i := index + 1; i <= list.length; i++ {
		list.data[i-1] = list.data[i]
	}
	list.length--
	return true
}

func (list *ArrayList) GetIndex(value interface{}) (int, error) {
	checkInit(list)
	if list.length == 0 {
		return 0, errors.New("数组无数据")
	}
	for i, v := range list.data {
		if v == value {
			return i, nil
		}
	}
	return 0, errors.New("未查询出数据")
}

func (list *ArrayList) ToArray() []interface{} {
	return list.data
}
