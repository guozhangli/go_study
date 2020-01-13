package main

import (
	"bytes"
	"fmt"
	"strings"
)

type Invoker interface {
	Call(interface{})
}

type Struct struct {

}

func (s *Struct)Call(p interface{}){
	fmt.Println("from struct",p)
}
//函数体实现接口
type FuncCaller func(interface{})

func (f FuncCaller) Call(p interface{}){
	f(p)
}


func main() {
	var invoker Invoker
	s:=new(Struct) //s:=&Struct
	invoker=s
	invoker.Call("hello")

	var invoker2 Invoker
	invoker2=FuncCaller(func(v interface{}){
		fmt.Println("from function",v)
	})
	invoker2.Call("world")

	accumulate:=Accumulate(1)
	fmt.Println(accumulate())
	fmt.Println(accumulate())
	fmt.Printf("%p\n",accumulate)

	accumulate2:=Accumulate(10)
	fmt.Println(accumulate2())
	fmt.Printf("%p\n",accumulate2)

	generator:=playerGen("high noon")
	name,hp:=generator()
	fmt.Println(name,hp)

	fmt.Println(5,"hello",&struct {
		a int
	}{1},true)

	str:=joinString("pig","and","sam")
	fmt.Println(str)
	str=joinString("dog","and","tom","and","lucy")
	fmt.Println(str)
	fmt.Println(printtypeValue("abc",true,123,false,"def",890))

	print(1,"2",3)
}

func Accumulate(value int)func()int{
	return func()int{
		value++
		return value
	}
}

func playerGen(name string)func()(string,int){
	hp:=150//封装性
	return func()(string,int){
		return name,hp
	}
}

func joinString(slist ... string)string{
	var b bytes.Buffer
	for _,s:=range slist{
		b.WriteString(s+" ")
	}
	return strings.TrimRight(b.String()," ")
}

func printtypeValue(slist ...interface{})string{
	var b bytes.Buffer
	for _,s:=range slist{
		str:=fmt.Sprintf("%v",s)
		var typeString string
		switch s.(type) {
		case bool:
			typeString="bool"
		case string:
			typeString="string"
		case int:
			typeString="int"
		}
		b.WriteString("value:")
		b.WriteString(str)
		b.WriteString(" type:")
		b.WriteString(typeString)
		b.WriteString("\n")
	}
	return b.String()
}

func rawPrint(rawList... interface{}){
	for _,a:=range rawList{
		fmt.Println(a)
	}
}

func print(slist... interface{}){
	rawPrint(slist...)
}
