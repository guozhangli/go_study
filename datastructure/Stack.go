package TestProject

type Stack interface {
	IsEmpty() bool
	Length() int
	Push(value interface{})
	Pop() interface{}
	Clear()
	Distroy(s *Stack)
}
