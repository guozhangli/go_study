package TestProject

import "fmt"

type SeparateChainingHashST struct {
	M  int //散列表大小
	N  int //键值对总数
	ST []*SequentialSearchST
}

func NewSeparateChainingHashST(m int) *SeparateChainingHashST {
	st := &SeparateChainingHashST{
		M:  m,
		N:  0,
		ST: make([]*SequentialSearchST, m),
	}
	for i, _ := range st.ST {
		st.ST[i] = NewSequentialSearchST()
	}
	return st
}

type Str string

func (s Str) hashCode() int {
	h := 0
	bs := []byte(s)
	if h == 0 && len(bs) > 0 {
		for _, b := range bs {
			h = 31*h + int(b)
		}
	}
	return h
}

func (s *SeparateChainingHashST) hashCode(key Str) int {
	return (key.hashCode() & 0x7fffffff) % s.M
}

func (s *SeparateChainingHashST) Put(key string, value interface{}) {
	fmt.Printf("key:%s hashcode:%d\n", key, s.hashCode(Str(key)))
	s.ST[s.hashCode(Str(key))].Put(key, value)
}
func (s *SeparateChainingHashST) Get(key string) (interface{}, error) {
	return s.ST[s.hashCode(Str(key))].Get(key)
}
