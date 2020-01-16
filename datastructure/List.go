package TestProject

type List interface {
	HeadAdd(value interface{})
	TailAdd(value interface{})
	Insert(index int, value interface{}) bool
	Length() int
	Delete(value interface{})
	DeleteF()
	DeleteT()
	Clear()
	IsEmpty() bool
}
