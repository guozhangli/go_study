package main

import "fmt"

type MyInt int

type class struct {

}
func (c *class) Do(v int){
	fmt.Println("call method do:",v)
}
func funcDo(v int){
	fmt.Println("call function do:",v)
}
func main() {
	var b MyInt
	fmt.Println(b.IsZero())
	b=10
	fmt.Println(b.add(5))

	/////////delegate
	var delegate func(int)
	c:=new(class)
	delegate=c.Do
	delegate(50)

	delegate=funcDo
	delegate(100)
}

func (m MyInt) IsZero()bool{
	return m==0
}

func (m MyInt) add(n int)int{
	return n+int(m)
}