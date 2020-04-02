package TestProject

type Stack interface {
	IsEmpty() bool
	Length() int
	Push(value interface{})
	Pop() interface{}
	Pop2() interface{}
	Clear()
	Distroy(s *Stack)
	GetTop() interface{}
}
