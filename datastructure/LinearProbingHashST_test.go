package TestProject

import "testing"

func TestNewLinearProbingHashST(t *testing.T) {
	lp := NewLinearProbingHashST(M)
	t.Log(lp)
}

func TestLinearProbingHashST_Put(t *testing.T) {
	lp := NewLinearProbingHashST(M)
	lp.Put("S", "sssss")
	lp.Put("T", "ttttt")
	lp.Put("A", "aaaaa")
	t.Log(lp)
}

func TestLinearProbingHashST_Get(t *testing.T) {
	lp := NewLinearProbingHashST(M)
	lp.Put("S", "sssss")
	lp.Put("T", "ttttt")
	lp.Put("A", "aaaaa")
	v, err := lp.Get("A")
	if err != nil {
		t.Log(err)
	} else {
		t.Log(v)
	}
}

func TestLinearProbingHashST_Delete(t *testing.T) {
	lp := NewLinearProbingHashST(M)
	lp.Put("S", "sssss")
	lp.Put("T", "ttttt")
	lp.Put("A", "aaaaa")
	lp.Put("C", "ccccc")
	lp.Put("B", "bbbbb")
	lp.Delete("B")
	t.Log(lp)
}
