package TestProject

import (
	"testing"
)

func TestNewSet(t *testing.T) {
	set := NewSet()
	t.Log(set)
}

func TestSet_Insert(t *testing.T) {
	set := NewSet()
	set.Insert("aaaa")
	t.Log(set)
}

func TestSet_IsMember(t *testing.T) {
	set := NewSet()
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
	set := NewSet()
	set.Insert("aaaa")
	set.Insert("bbbb")
	set.Insert("cccc")
	return set
}

func initSet2() *Set {
	set := NewSet()
	set.Insert("cccc")
	set.Insert("eeee")
	set.Insert("ffff")
	return set
}

func TestUnion(t *testing.T) {
	set1 := initSet()
	set2 := initSet2()
	set3 := Union(set1, set2)
	t.Log(set3.Size())
}

func TestIntersection(t *testing.T) {
	set1 := initSet()
	set2 := initSet2()
	set3 := Intersection(set1, set2)
	t.Log(set3.Size())
}

func TestDifference(t *testing.T) {
	set1 := initSet()
	set2 := initSet2()
	set3 := Difference(set1, set2)
	t.Log(set3.Size())
}

func initSet3() *Set {
	set := NewSet()
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
