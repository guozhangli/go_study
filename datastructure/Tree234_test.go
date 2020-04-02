package TestProject

import "testing"

func TestNewTree234(t *testing.T) {
	tree234 := NewTree234()
	t.Log(tree234)
}

func TestCreateNode234(t *testing.T) {
	tree234 := NewTree234()
	node := tree234.CreateNode234(CreateMap("A", "aaaaa"))
	t.Log(node)
}

func TestNode234_IsLeaf(t *testing.T) {
	tree234 := NewTree234()
	node := tree234.CreateNode234(CreateMap("A", "aaaaa"))
	flag := node.IsLeaf()
	t.Log(flag)
}

func TestNode234_IsFull(t *testing.T) {
	tree234 := NewTree234()
	node := tree234.CreateNode234(CreateMap("A", "aaaaa"))
	flag := node.IsFull()
	t.Log(flag)
}

func TestTree234_Find(t *testing.T) {
	tree234 := NewTree234()
	tree234.Insert("S", "sssss")
	tree234.Insert("E", "eeeee")
	tree234.Insert("A", "aaaaa")
	tree234.Insert("R", "rrrrr")
	tree234.Insert("C", "ccccc")
	tree234.Insert("H", "hhhhh")
	tree234.Insert("X", "xxxxx")
	tree234.Insert("M", "mmmmm")
	tree234.Insert("P", "ppppp")
	tree234.Insert("L", "lllll")
	tree234.Insert("J", "jjjjj")
	value := tree234.Find("J")
	t.Log(value)
}

func TestTree234_Insert(t *testing.T) {
	tree234 := NewTree234()
	tree234.Insert("S", "sssss")
	tree234.Insert("E", "eeeee")
	tree234.Insert("A", "aaaaa")
	tree234.Insert("R", "rrrrr")
	tree234.Insert("C", "ccccc")
	tree234.Insert("H", "hhhhh")
	tree234.Insert("X", "xxxxx")
	tree234.Insert("M", "mmmmm")
	tree234.Insert("P", "ppppp")
	tree234.Insert("L", "lllll")
	tree234.Insert("J", "jjjjj")
	t.Log(tree234)
}

func TestTree234_MidOrder(t *testing.T) {
	tree234 := NewTree234()
	tree234.Insert("S", "sssss")
	tree234.Insert("E", "eeeee")
	tree234.Insert("A", "aaaaa")
	tree234.Insert("R", "rrrrr")
	tree234.Insert("C", "ccccc")
	tree234.Insert("H", "hhhhh")
	tree234.Insert("X", "xxxxx")
	tree234.Insert("M", "mmmmm")
	tree234.Insert("P", "ppppp")
	tree234.Insert("L", "lllll")
	tree234.Insert("J", "jjjjj")
	tree234.MidOrderAndPrint()
}
