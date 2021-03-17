package TestProject

import (
	"testing"
)

func TestNewSet(t *testing.T) {
	set := NewSet(func(o, n interface{}) bool {
		return o == n
	})
	t.Log(set)
}

func TestSet_Insert(t *testing.T) {
	set := NewSet(func(o, n interface{}) bool {
		return o == n
	})
	set.Insert("aaaa")
	t.Log(set)
}

func TestSet_IsMember(t *testing.T) {
	set := NewSet(func(o, n interface{}) bool {
		return o == n
	})
	set.Insert("aaaa")
	set.Insert("bbbb")
	f := set.Insert("bbbb")
	if !f {
		t.Log("repeate")
	}
	flag := set.IsMember("cccc")
	t.Log(flag)
}

func TestSet_Delete(t *testing.T) {
	set := initSet()
	t.Log(set.Size())
	set.Delete("cccc")
	t.Log(set.Size())
}

func initSet() *Set {
	set := NewSet(func(o, n interface{}) bool {
		return o == n
	})
	set.Insert("aaaa")
	set.Insert("bbbb")
	set.Insert("cccc")
	return set
}

func initSet2() *Set {
	set := NewSet(func(o, n interface{}) bool {
		return o == n
	})
	set.Insert("cccc")
	set.Insert("eeee")
	set.Insert("ffff")
	return set
}

func TestUnion(t *testing.T) {
	set1 := initSet()
	set2 := initSet2()
	set3 := NewSet(func(o, n interface{}) bool {
		return o == n
	})
	Union(set1, set2, set3)
	t.Log(set3.Size())
}

func TestIntersection(t *testing.T) {
	set1 := initSet()
	set2 := initSet2()
	set3 := NewSet(func(o, n interface{}) bool {
		return o == n
	})
	Intersection(set1, set2, set3)
	t.Log(set3.Size())
}

func TestDifference(t *testing.T) {
	set1 := initSet()
	set2 := initSet2()
	set3 := NewSet(func(o, n interface{}) bool {
		return o == n
	})
	Difference(set1, set2, set3)
	t.Log(set3.Size())
}

func initSet3() *Set {
	set := NewSet(func(o, n interface{}) bool {
		return o == n
	})
	set.Insert("cccc")
	set.Insert("aaaa")
	set.Insert("bbbb")

	return set
}
func TestIsSubSet(t *testing.T) {
	set1 := initSet()
	set2 := initSet3()
	flag := IsSubSet(set1, set2)
	t.Log(flag)
}

func TestIsEqual(t *testing.T) {
	set1 := initSet()
	set2 := initSet3()
	flag := IsEqual(set1, set2)
	t.Log(flag)
}

func TestConveringExample(t *testing.T) {
	ConveringExample()
}

//容斥原理
func TestInclusion_exclusion(t *testing.T) {
	set1 := initSet()
	set2 := initSet2()
	set3 := NewSet(func(o, n interface{}) bool {
		return o == n
	})
	Union(set1, set2, set3)
	set4 := NewSet(func(o, n interface{}) bool {
		return o == n
	})
	Intersection(set1, set2, set4)
	size := set1.Size() + set2.Size() - set4.Size()
	flag := set3.Size() == size
	t.Log(flag)
}
