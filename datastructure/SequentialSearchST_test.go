package TestProject

import (
	"encoding/json"
	"testing"
)

var s *SequentialSearchST

func TestNewSequentialSearchST(t *testing.T) {
	t.Log(s)
}

func Init() {
	s = NewSequentialSearchST()
	s.Put("S", "sssss")
	s.Put("T", "ttttt")
	s.Put("V", "vvvvv")
}

func TestSequentialSearchST_Put(t *testing.T) {
	Init()
	s.Put("M", "mmmmm")
	s.Put("N", "nnnnn")
	str, _ := json.Marshal(s)
	t.Log(string(str))
}

func TestSequentialSearchST_Get(t *testing.T) {
	Init()
	v, err := s.Get("T")
	if err != nil {
		t.Log(err)
	} else {
		t.Log(v.(string))
	}
}

func TestSequentialSearchST_Delete(t *testing.T) {
	Init()
	s.Delete("T")
	str, _ := json.Marshal(s)
	t.Log(string(str))
}

func TestSequentialSearchST_Size(t *testing.T) {
	Init()
	c := s.Size()
	t.Log(c)
}
