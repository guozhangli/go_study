package main

import (
	"errors"
	"fmt"
)

var errDivisionByZero=errors.New("division by zero")

type ParseError struct {
	Filename string
	Line int
}
func main() {
	fmt.Println(div(1,0))
	result,err:=div(1,1)
	if err!=nil{
		fmt.Println(err)
	}else{
		fmt.Println(result)
	}

	var e error
	e=newParseError("test12.go",1)
	fmt.Println(e.Error())
	switch detail:=e.(type) {
	case *ParseError:
		fmt.Printf("filename:%s line:%d\n",detail.Filename,detail.Line)
	default:
		fmt.Println("other error")
	}
}

func div(dividend,divisor int)(int,error){
	if divisor==0{
		return 0,errDivisionByZero
	}
	return dividend/divisor,nil
}

func (e *ParseError) Error()string{
	return fmt.Sprintf("%s:%d",e.Filename,e.Line)
}

func newParseError(filename string,line int)error{
	return &ParseError{filename,line}
}
