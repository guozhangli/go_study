package main

import (
	"fmt"
	"reflect"
)

type dummy struct {
	a int
	b string
	float32
	bool
	next *dummy
}
func main() {
	d:=reflect.ValueOf(dummy{
		next:&dummy{},
	})
	fmt.Println("NumField",d.NumField())
	floatField:=d.Field(2)
	fmt.Println("field",floatField.Type())

	fmt.Println("fieldByName(\"b\").Type",d.FieldByName("b").Type())
	fmt.Println("fieldByIndex([]int{4,0}).Type()",d.FieldByIndex([]int{4,0}).Type())
}
