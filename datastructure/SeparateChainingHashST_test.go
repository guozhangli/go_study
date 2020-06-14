package TestProject

import (
	"encoding/json"
	"testing"
)

func TestNewSeparateChainingHashST(t *testing.T) {
	s := NewSeparateChainingHashST(5)
	str, _ := json.Marshal(s)
	t.Log(string(str))
}

func TestSeparateChainingHashST_Put(t *testing.T) {
	s := NewSeparateChainingHashST(5)
	s.Put("S", "sssss")
	s.Put("T", "ttttt")
	s.Put("V", "vvvvv")
	s.Put("A", "aaaaa")
	s.Put("T", "aaaaa")
	str, _ := json.Marshal(s)
	t.Log(string(str))
}

func TestSeparateChainingHashST_Get(t *testing.T) {
	s := NewSeparateChainingHashST(5)
	s.Put("S", "sssss")
	s.Put("T", "ttttt")
	s.Put("V", "vvvvv")
	s.Put("A", "aaaaa")
	s.Put("T", "aaaaa")
	v, err := s.Get("T")
	if err == nil {
		t.Log(v.(string))
	}
}

func TestSeparateChainingHashST_Delete(t *testing.T) {
	s := NewSeparateChainingHashST(5)
	s.Put("S", "sssss")
	s.Put("T", "ttttt")
	s.Put("V", "vvvvv")
	s.Put("A", "aaaaa")
	s.Put("T", "aaaaa")
	s.Put("M", "mmmmm")
	s.Put("N", "nnnnn")
	str, _ := json.Marshal(s)
	t.Log(string(str))
	s.Delete("N")
	str1, _ := json.Marshal(s)
	t.Log(string(str1))
}

func TestSeparateChainingHashST_Contain(t *testing.T) {
	s := NewSeparateChainingHashST(5)
	s.Put("S", "sssss")
	s.Put("T", "ttttt")
	s.Put("V", "vvvvv")
	s.Put("A", "aaaaa")
	s.Put("T", "aaaaa")
	s.Put("M", "mmmmm")
	s.Put("N", "nnnnn")
	str, _ := json.Marshal(s)
	t.Log(string(str))
	if s.Contain("C") {
		t.Log("true")
	} else {
		t.Log("false")
	}
}
