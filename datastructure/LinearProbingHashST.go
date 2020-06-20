package TestProject

import (
	"errors"
	"fmt"
)

//基于线性探测的符号表(开地址哈希表)
/**
 假设进行均匀散列，我们能够在一个开地址哈希表中探查的槽位个数是：1/（1-a）,a是负载因子。
当哈希表的负载因子趋近于1时，探查槽位的数量将显著增大。在一个对时间特别敏感的应用中，就可以增加哈希表的空间来提高探查的效率。


*/
type LinearProbingHashST struct {
	N      int //符号表中键值对的总数
	M      int //线性探测表的大小
	Keys   []string
	Values []interface{}
}

const M = 16

func NewLinearProbingHashST(cap int) *LinearProbingHashST {
	return &LinearProbingHashST{
		N:      0,
		M:      cap,
		Keys:   make([]string, cap),
		Values: make([]interface{}, cap),
	}
}

func (lp *LinearProbingHashST) hashCode(key Str) int {
	return (key.hashCode() & 0x7fffffff) % lp.M
}

func (lp *LinearProbingHashST) Resize(cap int) {
	lpRedo := NewLinearProbingHashST(cap)
	for i := 0; i < lp.M; i++ {
		if lp.Keys[i] != "" {
			lpRedo.Put(lp.Keys[i], lp.Values[i])
		}
	}
	lp.Keys = lpRedo.Keys
	lp.Values = lpRedo.Values
	lp.M = lpRedo.M
}

func (lp *LinearProbingHashST) Put(key string, value interface{}) {
	if lp.N >= lp.M/2 {
		lp.Resize(M * 2)
	}
	var i int
	for i = lp.hashCode(Str(key)); lp.Keys[i] != ""; i = (i + 1) % lp.M {
		if CompareTo(lp.Keys[i], key) == 0 {
			lp.Values[i] = value
			return
		}
	}
	lp.Keys[i] = key
	lp.Values[i] = value
	lp.N++
}

func (lp *LinearProbingHashST) Get(key string) (interface{}, error) {
	var i int
	for i = lp.hashCode(Str(key)); lp.Keys[i] != ""; i = (i + 1) % lp.M {
		if CompareTo(lp.Keys[i], key) == 0 {
			return lp.Values[i], nil
		}
	}
	return nil, errors.New(fmt.Sprintf("Not found key :%s", key))
}

func (lp *LinearProbingHashST) Delete(key string) {
	if !lp.Contains(key) {
		fmt.Printf("key : %s is not exist", key)
		return
	}
	var i int
	i = lp.hashCode(Str(key))
	for CompareTo(lp.Keys[i], key) != 0 {
		i = (i + 1) % lp.M
	}
	lp.Keys[i] = ""
	lp.Values[i] = nil
	i = (i + 1) % lp.M
	for lp.Keys[i] != "" {
		keyToRedo := lp.Keys[i]
		valueToRedo := lp.Values[i]
		lp.Keys[i] = ""
		lp.Values[i] = nil
		lp.N--
		lp.Put(keyToRedo, valueToRedo)
		i = (i + 1) % lp.M
	}
	lp.N--
	if lp.N > 0 && lp.N == lp.M/8 {
		lp.Resize(lp.M / 2)
	}
}

func (lp *LinearProbingHashST) Contains(key string) bool {
	if _, err := lp.Get(key); err != nil {
		return false
	}
	return true
}
