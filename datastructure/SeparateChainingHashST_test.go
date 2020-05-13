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
	str, _ := json.Marshal(s)
	t.Log(string(str))
}

func TestSeparateChainingHashST_Get(t *testing.T) {

}
