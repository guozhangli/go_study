package TestProject

import (
	"errors"
	"fmt"
)

type NodeST struct {
	Key   string
	Value interface{}
	Next  *NodeST
}

func NewNodeST(key string, val interface{}, next *NodeST) *NodeST {
	return &NodeST{
		Key:   key,
		Value: val,
		Next:  next,
	}
}

//无序链表
type SequentialSearchST struct {
	First *NodeST
}

func NewSequentialSearchST() *SequentialSearchST {
	return &SequentialSearchST{
		First: nil,
	}
}

func (s *SequentialSearchST) Get(key string) (interface{}, error) {
	for node := s.First; node != nil; node = node.Next {
		if node.Key == key {
			return node.Value, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("Not found key='%s'", key))
}

//{"First":{"Key":"N","Value":"nnnnn","Next":{"Key":"M","Value":"mmmmm","Next":{"Key":"V","Value":"vvvvv","Next":{"Key":"T","Value":"ttttt","Next":{"Key":"S","Value":"sssss","Next":null}}}}}}
func (s *SequentialSearchST) Put(key string, value interface{}) {
	//查找给的键，找到则更新，否则新建节点插入到头
	for node := s.First; node != nil; node = node.Next {
		if node.Key == key {
			node.Value = value
		}
	}
	s.First = NewNodeST(key, value, s.First)
}

func (s *SequentialSearchST) Delete(key string) {
	var pre *NodeST
	for node := s.First; node != nil; node = node.Next {
		if node.Key == key {
			pre.Next = node.Next
			node = nil
			break
		}
		pre = node
	}
}

func (s *SequentialSearchST) Size() int {
	var count int
	for node := s.First; node != nil; node = node.Next {
		count++
	}
	return count
}
