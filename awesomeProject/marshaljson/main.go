package main

import "fmt"

func main() {

	type Skill struct {
		Name string
		Level int
	}
	type Actor struct {
		Name string
		Age int
		Skills []Skill
	}

	a:=Actor{
		Name:"cow boy",
		Age:37,
		Skills:[]Skill{
			{Name:"roll and rool",Level:1},
			{Name:"flash your dog eye",Level:2},
			{Name:"time to have lunch",Level:3},
		},
	}
	if result,err:=MarshalJson(a);err==nil{
		fmt.Println(result)
	}else{
		fmt.Println(err)
	}

}
