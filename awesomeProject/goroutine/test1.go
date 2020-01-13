package main

import (
	"fmt"
	"runtime"
	"time"
)

func running(){
	var times int
	for {
		times++
		fmt.Println("tick",times)
		time.Sleep(time.Second*10)
	}
}
func main(){
	fmt.Println(runtime.NumCPU())
	go running()
	var input string
	fmt.Scanln(&input)
	fmt.Println("haha")
	go func() {
		for {
			var times1 int
			times1++
			fmt.Println("tick", times1)
			time.Sleep(time.Second)
		}
	}()
	var input1 string
	fmt.Scanln(&input1)
}

