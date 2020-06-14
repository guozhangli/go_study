package TestProject

//符号表接口
type SymbolTable interface {
	Get(key string) (interface{}, error)
	Put(key string, value interface{})
	Delete(key string)
	Size()
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
