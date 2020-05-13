package TestProject

import (
	"errors"
	"fmt"
)

//二分查找符号表

type BinarySearchST struct {
	Key   []string
	Value []interface{}
	N     int
}

func NewBinarySearchST(n int) *BinarySearchST {
	return &BinarySearchST{
		Key:   make([]string, n),
		Value: make([]interface{}, n),
		N:     0,
	}
}

func (bs *BinarySearchST) Rank(key string) int {
	io, ih := 0, bs.N-1
	for io <= ih {
		mid := io + (ih-io)/2
		c := CompareTo(key, bs.Key[mid])
		if c > 0 {
			io = mid + 1
		} else if c < 0 {
			ih = mid - 1
		} else {
			return mid
		}
	}
	return io
}

func (bs *BinarySearchST) Get(key string) (interface{}, error) {
	if bs.Key == nil || len(bs.Key) == 0 {
		return nil, errors.New("Keys is nil")
	}
	ki := bs.Rank(key)
	if ki < bs.N && CompareTo(key, bs.Key[ki]) == 0 {
		return bs.Value[ki], nil
	} else {
		return nil, errors.New("Key not found")
	}
}

func (bs *BinarySearchST) Put(key string, value interface{}) {
	ki := bs.Rank(key)
	if ki < bs.N && CompareTo(key, bs.Key[ki]) == 0 {
		bs.Value[ki] = value
		return
	}
	for i := bs.N - 1; i >= ki; i-- {
		bs.Key[i+1] = bs.Key[i]
		bs.Value[i+1] = bs.Value[i]
	}
	bs.Key[ki] = key
	bs.Value[ki] = value
	bs.N++
}

func (bs *BinarySearchST) Delete(key string) {
	ki := bs.Rank(key)
	if ki < bs.N && CompareTo(key, bs.Key[ki]) == 0 {
		for i := ki; i < bs.N; i++ {
			bs.Key[i] = bs.Key[i+1]
			bs.Value[i] = bs.Value[i+1]
		}
		bs.N--
		return
	} else {
		fmt.Printf("delete is error,no exist key %s", key)
	}
}

func (bs *BinarySearchST) Size() int {
	return bs.N
}
