package TestProject

import "fmt"

//基于分离链接法的符号表
/**

链式哈希表描述

链式哈希表从根本上来说是由一组链表构成。每个链表都可以看作是一个”桶“，我们将所有的元素通过散列的方式
放到具体的不同的桶中。插入元素时，首先将其键传入一个哈希函数，函数通过散列的方式告知元素属于哪个“桶”，
然后在相应的链表头插入元素。查找或删除元素时，用同样的方式先找到元素的桶，然后遍历相应的链表，直到发现
我们想要的元素。因为每个桶都是一个链表，所以链式哈希表并不限制包含元素的个数。然而，如果表变得太大，它的
性能会降低。

解决冲突

当哈希表中两个键散列到相同的槽位时这两个键将会发生冲突。链式哈希表解决冲突的方式非常简单，当冲突发生时，它
将元素放入已经准备好的桶中。但这洋会带来一个问题，当过多的冲突发生在同一个槽位时，此位置的桶将变得越来越深，
从而造成访问在这个位置上的元素所需要的时间越来越多。
在理想情况下，我们希望所有的桶以几乎相同的速度增长，这样它们就可以尽可能的保持小的容量和相同的大小。换句话说，
我们的目标就是尽可能均匀地随机的分配表中的元素，这种情况在理论上称为均匀散列。而在实际中，我们只能尽可能近似达到
这种状态。
如果想插入表中的元素数量远大于表中桶的数量，那么即使时在一个均匀散列的过称中，表的性能也会迅速降低.在这种
情况下，表中所有的桶都变得越来越深。因此，我们必须要特别注意一个哈希表的负载因子，其定义为：a=n/m。
其中n是元素的个数，m是桶的个数，在均匀散列的情况下，链式哈系表的负载因子告诉我们表中的桶能装下的元素的最大个数。
当有一个表的负载因子小于1时，它告诉我们这个表的每个桶装载的元素不超过1个。当然，由于均匀散列时理想的近似情况，
因此在实际情况中我们往往多少会检索负载因子建议的数值。如何达到更接近于均匀散列的情况，最终取决于如何选择哈希函数。

*/
type SeparateChainingHashST struct {
	M  int //散列表大小
	N  int //键值对总数
	ST []*SequentialSearchST
}

func NewSeparateChainingHashST(m int) *SeparateChainingHashST {
	st := &SeparateChainingHashST{
		M:  m,
		N:  0,
		ST: make([]*SequentialSearchST, m),
	}
	for i, _ := range st.ST {
		st.ST[i] = NewSequentialSearchST()
	}
	return st
}

func (s *SeparateChainingHashST) hashCode(key Str) int {
	return (key.hashCode() & 0x7fffffff) % s.M
}

/**
时间复杂度O（1）
*/
func (s *SeparateChainingHashST) Put(key string, value interface{}) {
	fmt.Printf("key:%s hashcode:%d\n", key, s.hashCode(Str(key)))
	//如果存在键则更新，不存在键则新增
	if s.ST[s.hashCode(Str(key))].Put(key, value) {
		s.N++
	}
}

/**
时间复杂度O（1）
*/
func (s *SeparateChainingHashST) Get(key string) (interface{}, error) {
	return s.ST[s.hashCode(Str(key))].Get(key)
}

/**
时间复杂度O（1）
*/
func (s *SeparateChainingHashST) Delete(key string) {
	if s.ST[s.hashCode(Str(key))].Delete(key) {
		s.N--
	}
}

/**
时间复杂度O（1）
*/
func (s *SeparateChainingHashST) Size() int {
	return s.N
}

/**
时间复杂度O（1）
*/
func (s *SeparateChainingHashST) Contain(key string) bool {
	if v, err := s.Get(key); err == nil {
		if v != nil {
			return true
		}
	}
	return false
}
