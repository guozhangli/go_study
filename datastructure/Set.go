package TestProject

import (
	"fmt"
	"sync"
)

/**
集合
	特征：1，成员是无序的。2，每个成员在集合中只出现一次。
*/

type Set LinkedList

type ISet interface {
	Insert(value interface{}) bool
	Delete(value interface{}) bool
	Union(set1, set2 *Set) *Set
	Intersection(set1, set2 *Set) *Set
	Difference(set1, set2 *Set) *Set
	IsMember(value interface{}) bool
	IsSubSet(set1, set2 *Set) bool
	IsEqual(set1, set2 *Set) bool
	Size(set *Set) int
}

var IsMatch func(oldValue, newValue interface{}) bool

func NewSet(f func(o, n interface{}) bool) *Set {
	IsMatch = f
	return new(Set)
}

var lock sync.Mutex

//向集合中插入成员
func (set *Set) Insert(value interface{}) bool {
	defer lock.Unlock()
	lock.Lock()
	if set.IsMember(value) {
		return false
	}
	(*LinkedList)(set).TailAdd(value)
	return true
}

//判断集合中是否存在某成员
func (set *Set) IsMember(value interface{}) bool {
	ll := (*LinkedList)(set)
	node := ll.nextNode
	for node != nil {
		if IsMatch(node.data, value) {
			return true
		}
		node = node.nextNode
	}
	return false
}

//从集合中删除成员
func (set *Set) Delete(value interface{}) bool {
	defer lock.Unlock()
	lock.Lock()
	if !set.IsMember(value) {
		return false
	}
	(*LinkedList)(set).Delete(value)
	return true
}

func (set *Set) Clear() {
	iterator := set.Iterator()
	for iterator.HasNode() {
		set.Delete(iterator.data)
		iterator = iterator.NextNode()
	}
}

//计算集合的成员数量
func (set *Set) Size() int {
	defer lock.Unlock()
	lock.Lock()
	ll := (*LinkedList)(set)
	return ll.Length()
}

func (set *Set) Iterator() *LinkedList {
	return set.nextNode
}

func (ll *LinkedList) NextNode() *LinkedList {
	return ll.nextNode
}

func (ll *LinkedList) HasNode() bool {
	if ll != nil {
		return true
	}
	return false
}
func (ll *LinkedList) Data() interface{} {
	return ll.data
}

//并集
func Union(set1, set2, set3 *Set) {
	node := set1.nextNode
	for node != nil {
		set3.Insert(node.data)
		node = node.nextNode
	}
	node2 := set2.nextNode
	for node2 != nil {
		set3.Insert(node2.data)
		node2 = node2.nextNode
	}
}

//交集
func Intersection(set1, set2, set3 *Set) {
	node := set1.nextNode
	for node != nil {
		if set2.IsMember(node.data) {
			set3.Insert(node.data)
		}
		node = node.nextNode
	}
}

//差集：只包含在set1中出现过，且不属于set2的成员
func Difference(set1, set2, set3 *Set) {
	node := set1.nextNode
	for node != nil {
		if !set2.IsMember(node.data) {
			set3.Insert(node.data)
		}
		node = node.nextNode
	}
}

//是否是子集
//	如果集合set1包含集合set2的所有成员，则set2是set1的子集
//	注意：一个集合是它自身的子集
func IsSubSet(set1, set2 *Set) bool {
	node := set2.nextNode
	for node != nil {
		if set1.IsMember(node.data) {
			node = node.nextNode
		} else {
			return false
		}
	}
	return true
}

//集合是否相等
func IsEqual(set1, set2 *Set) bool {
	if set1.Size() != set2.Size() {
		return false
	}
	return IsSubSet(set1, set2)
}

/**set 示例：集合覆盖
	S={a,b,c,d,e,f,g,h,i,j,k,l}
	P={A1,...A7}
	A1={a,b,c,d},A2={e,f,g,h,i},A3={j,k,l},A4={a,e},A5={b,f,g},A6={c,d,g,h,k,l},A7={l}
    最优的覆盖集是C={A1,A2,A3},这里给出的算法选择的集合是C={A6,A2,A1,A3}
	针对集合覆盖的算法是一种近似算法。它并不总是获得最优解，但肯定能够达到对数级的复杂度。
	该算法的工作原理是：不断从P中选出一个集合，使其能够覆盖S中最多的成员数量。换句话说，
	该算法每次都尝试尽可能早覆盖S中更多的成员，因此该算法采用了贪心法的思路。
*/

type pSet struct {
	key  string
	pSet *Set
}

func createSet(str ...string) *Set {
	s := NewSet(func(o, n interface{}) bool {
		return o == n
	})
	s1 := str
	for _, v := range s1 {
		s.Insert(v)
	}
	return s
}
func ConveringExample() {
	s := createSet("a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l")
	p := NewSet(func(o, n interface{}) bool {
		return o.(*pSet).key == n.(*pSet).key
	})
	pArr := []*pSet{
		&pSet{"A1", createSet("a", "b", "c", "d")},
		&pSet{"A2", createSet("e", "f", "g", "h", "i")},
		&pSet{"A3", createSet("j", "k", "l")},
		&pSet{"A4", createSet("a", "e")},
		&pSet{"A5", createSet("b", "f", "g")},
		&pSet{"A6", createSet("c", "d", "g", "h", "k", "l")},
		&pSet{"A7", createSet("l")},
	}
	for _, v := range pArr {
		p.Insert(v)
	}
	Covering(s, p)
}
func Covering(s *Set, p *Set) {
	keys := NewArray(7)
	MaxCovering(s, p, keys)
	for i := 0; i < keys.length; i++ {
		fmt.Printf("%s\n", keys.GetValue(i))
	}
}

func MaxCovering(s *Set, p *Set, list *ArrayList) {
	node := p.nextNode
	var max int
	var sMax *Set
	var pMax *pSet
	var key string
	if s.Size() == 0 || p.Size() == 0 {
		return
	}
	for node != nil {
		pMax := node.data.(*pSet)
		s3 := NewSet(func(o, n interface{}) bool {
			return o == n
		})
		Intersection(s, pMax.pSet, s3)
		size := s3.Size()
		if max < size {
			max = size
			key = node.data.(*pSet).key
			sMax = pMax.pSet
		}
		node = node.nextNode
	}
	sNode := sMax.nextNode
	for sNode != nil {
		s.Delete(sNode.data)
		sNode = sNode.nextNode
	}
	p.Delete(pMax)
	list.Add(key)
	MaxCovering(s, p, list)
}
