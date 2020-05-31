package TestProject

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

func NewSet() *Set {
	return new(Set)
}

//向集合中插入成员
func (set *Set) Insert(value interface{}) bool {
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
		if node.data == value {
			return true
		}
		node = node.nextNode
	}
	return false
}

//从集合中删除成员
func (set *Set) Delete(value interface{}) bool {
	if !set.IsMember(value) {
		return false
	}
	(*LinkedList)(set).Delete(value)
	return true
}

//计算集合的成员数量
func (set *Set) Size() int {
	ll := (*LinkedList)(set)
	return ll.Length()
}

//并集
func Union(set1, set2 *Set) *Set {
	set3 := NewSet()
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
	return set3
}

//交集
func Intersection(set1, set2 *Set) *Set {
	set3 := NewSet()
	node := set1.nextNode
	for node != nil {
		if set2.IsMember(node.data) {
			set3.Insert(node.data)
		}
		node = node.nextNode
	}
	return set3
}

//差集：只包含在set1中出现过，且不属于set2的成员
func Difference(set1, set2 *Set) *Set {
	set3 := NewSet()
	node := set1.nextNode
	for node != nil {
		if !set2.IsMember(node.data) {
			set3.Insert(node.data)
		}
		node = node.nextNode
	}
	return set3
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
