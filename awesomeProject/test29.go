package main

import (
	"fmt"
)

type Flyer interface {
	Fly()
}
type Walker interface {
	Walk()
}

type bird struct {
}

func (*bird) Fly() {
	fmt.Println("bird:fly")
}
func (*bird) Walk() {
	fmt.Println("bird:walk")
}

type pig struct {
}

func (*pig) Walk() {
	fmt.Println("pig:walk")
}

func main() {
	animals := map[string]interface{}{
		"bird": new(bird),
		"pig":  new(pig),
	}
	for name, obj := range animals {
		f, isFlyer := obj.(Flyer)
		w, isWalker := obj.(Walker)
		fmt.Printf("name:%s isFlyer:%v isWalker:%v\n", name, isFlyer, isWalker)
		if isFlyer {
			f.Fly()
		}

		if isWalker {
			w.Walk()
		}
	}

	p1 := new(pig)
	var p Walker = p1
	p2 := p.(*pig)
	fmt.Printf("%+v\n", p2)

	var a interface{}
	a = 1
	fmt.Printf("%v,%T\n", a, a)
	a = "hello"
	fmt.Println(a)
	a = false
	fmt.Println(a)

	var b int = 1
	var i interface{} = b
	fmt.Println(i)
	var c int = i.(int)
	fmt.Println(c)

	var e interface{} = 1
	var f interface{} = "hi"
	fmt.Println(e == f)
}
