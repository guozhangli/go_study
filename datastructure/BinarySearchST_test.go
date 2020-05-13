package TestProject

import "testing"

func TestNewBinarySearchST(t *testing.T) {
	b := NewBinarySearchST(5)
	t.Log(b)
}

func TestBinarySearchST_Put(t *testing.T) {
	b := NewBinarySearchST(10)
	b.Put("S", "sssss")
	b.Put("T", "ttttt")
	b.Put("A", "aaaaa")
	b.Put("B", "bbbbb")
	t.Log(b)
}

func TestBinarySearchST_Get(t *testing.T) {
	b := NewBinarySearchST(10)
	b.Put("S", "sssss")
	b.Put("T", "ttttt")
	b.Put("A", "aaaaa")
	b.Put("B", "bbbbb")
	t.Log(b)
	v, err := b.Get("T")
	if err != nil {
		t.Log(err)
	} else {
		t.Log(v)
	}
}

func TestBinarySearchST_Size(t *testing.T) {
	b := NewBinarySearchST(10)
	b.Put("S", "sssss")
	b.Put("T", "ttttt")
	b.Put("A", "aaaaa")
	b.Put("B", "bbbbb")
	t.Log(b.Size())
}

func TestBinarySearchST_Delete(t *testing.T) {
	b := NewBinarySearchST(10)
	b.Put("S", "sssss")
	b.Put("T", "ttttt")
	b.Put("A", "aaaaa")
	b.Put("B", "bbbbb")
	b.Delete("S")
	t.Log(b)
}
