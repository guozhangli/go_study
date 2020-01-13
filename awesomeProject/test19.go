package main

import "fmt"

type Wheel struct {
	Size int
}
type Engine struct {
	Power int
	Type string
}
type Car struct {
	Wheel
	Engine
}
type Car2 struct {
	Wheel
	Engine struct{
		Power int
		Type string
	}
}
func main() {
	c:=Car{
		Wheel{Size:18,},
		Engine{Type:"1.5T",Power:182},
	}
	fmt.Printf("%+v\n",c)

	c2:=Car2{
		Wheel:Wheel{
			Size:19,
		},
		Engine:struct{
			Power int
			Type string
		}{
			Power:190,
			Type:"1.9T",
		},
	}
	fmt.Printf("%+v\n",c2)
}
