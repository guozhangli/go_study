package main

import "fmt"

type Data struct {
	int
	float32
	bool
}
type BasicColor struct {
	R,G,B float32
}
type Color struct {
	BasicColor
	Alpha float32
}
type Flying struct {

}
func (f *Flying) Fly(){
	fmt.Println("can fly")
}
type Walkable struct {

}
func (w *Walkable) Walk(){
	fmt.Println("can walk")
}

type Human struct {
	Walkable
}
type Bird struct {
	Flying
	Walkable
}
func main() {
	ins:=&Data{
		int:20,
		float32:20.3,
		bool:true,
	}
	fmt.Println(ins)

	var c Color
	c.BasicColor.R=1
	c.BasicColor.G=1
	c.B=0
	c.Alpha=1
	fmt.Printf("%+v\n",c)

	b:=new(Bird)
	fmt.Println("bird:")
	b.Walk()
	b.Fly()

	fmt.Println("human:")
	h:=new(Human)
	h.Walk()
}
