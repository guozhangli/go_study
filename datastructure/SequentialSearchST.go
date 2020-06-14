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
func (s *SequentialSearchST) Put(key string, value interface{}) bool {
	//查找给的键，找到则更新，否则新建节点插入到头
	for node := s.First; node != nil; node = node.Next {
		if node.Key == key {
			node.Value = value
			return false
		}
	}
	s.First = NewNodeST(key, value, s.First)
	return true
}

func (s *SequentialSearchST) Delete(key string) bool {
	var pre *NodeST
	for node := s.First; node != nil; node = node.Next {
		if node.Key == key {
			if pre != nil {
				pre.Next = node.Next
			} else {
				s.First = node.Next
			}
			node = nil
			return true
		}
		pre = node
	}
	return false
}

func (s *SequentialSearchST) Size() int {
	var count int
	for node := s.First; node != nil; node = node.Next {
		count++
	}
	return count
}
