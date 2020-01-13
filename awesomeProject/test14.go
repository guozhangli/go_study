package main

import "fmt"

type Point struct{
	X int
	Y int
	R,G,B byte
}
type Player1 struct {
	Name string
	HealthPoint int
	MagicPoint int
}
type Command struct{
	Name string
	Var *int
	Comment string
}
func main() {
	var p Point
	p.X=10
	p.Y=20
	fmt.Println(p)

	tank:=new(Player1)
	(*tank).Name="canon"
	tank.HealthPoint=10
	fmt.Println(tank)

	var version int=2
	cmd:=&Command{}
	(*cmd).Name="version"
	cmd.Var=&version
	cmd.Comment="show version"
	fmt.Println(cmd)

	msg:=&struct {
		id int
		data string
	}{
		1024,
		"hello",
	}
	printMsgType(msg)
}
func printMsgType(msg *struct{
	id int
	data string
}){
	fmt.Printf("%T\n",msg)
	fmt.Println(msg)
}