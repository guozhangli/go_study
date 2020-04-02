package TestProject

type Queue interface {
	DeQueue() interface{}
	EnQueue(value interface{})
	IsEmpty() bool
	Length() int
	Clear()
	GetFront() interface{}
	Distroy(q *Queue)
	IsExist(value interface{}) bool
}
