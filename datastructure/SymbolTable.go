package TestProject

//符号表接口
type SymbolTable interface {
	Get(key string) (interface{}, error)
	Pub(key string, value interface{})
	Delete(key string)
	Size()
}
