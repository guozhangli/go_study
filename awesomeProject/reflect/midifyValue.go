package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a int=1024
	valueOfA:=reflect.ValueOf(&a)
	valueOfA=valueOfA.Elem()
	valueOfA.SetInt(1)
	fmt.Println(valueOfA.Int())

	type dog struct {
		LegCount int
	}
	valueOfDog:=reflect.ValueOf(&dog{})
	vLegCount:=valueOfDog.Elem().FieldByName("LegCount")
	vLegCount.SetInt(4)
	fmt.Println(vLegCount.Int())

	var a1 int
	typeOfA:=reflect.TypeOf(a1)
	aIns:=reflect.New(typeOfA)
	fmt.Println(aIns.Type(),aIns.Kind())
}