package main

import "fmt"

type Stringer interface {
	String() string
}

type MyImplement struct {

}

func (m *MyImplement) String() string{
	return "hi"
}

func GetStringer() Stringer{
	var s *MyImplement=nil
	return s
}

func main() {
	if GetStringer()==nil{
		fmt.Println("GetStringer()==nil")
	}else {
		fmt.Println("GetStringer()!=nil")
	}
}