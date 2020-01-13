package main

import (
	"flag"
	"fmt"
)
var skillParam=flag.String("skill","","skill to perform")
func main() {
	//匿名函数
	func(data int){
		fmt.Println("hello",data)
	}(100)

	f:=func(data int){
		fmt.Println("world",data)
	}
	f(100)
	visit([]int{1,2,3,4,5}, func(i int) {
		fmt.Println(i)
	})

	flag.Parse()
	var skill=map[string]func(){
		"fire": func() {
			fmt.Println("chicken fire")
		},
		"run": func() {
			fmt.Println("soldier run")
		},
		"fly": func() {
			fmt.Println("angel fly")
		},
	}

	if f,ok:=skill[*skillParam];ok{
		f()
	}else{
		fmt.Println("skill not found",ok)
	}

}

func visit(list []int,f func(int)){
	for _,v:=range list{
		f(v)
	}
}