package main

import (
	"fmt"
	"reflect"
)

type Eunm int

const Zero Eunm  =0
func main() {
	var a int
	typeOfA:=reflect.TypeOf(a)
	fmt.Println(typeOfA.Name(),typeOfA.Kind())
	var p *int
	typeOfp:=reflect.TypeOf(p)
	fmt.Println(typeOfp.Name(),typeOfp.Kind())

	type cat struct {

	}
	typeOfcat:=reflect.TypeOf(cat{})
	fmt.Println(typeOfcat.Name(),typeOfcat.Kind())

	typeOfB:=reflect.TypeOf(Zero)
	fmt.Println(typeOfB.Name(),typeOfB.Kind())

	ins:=&cat{}
	typeOfcat2:=reflect.TypeOf(ins)
	fmt.Println(typeOfcat2.Name(),typeOfcat2.Kind())
	typeOfcat2=typeOfcat2.Elem()
	fmt.Println(typeOfcat2.Name(),typeOfcat2.Kind())
}
