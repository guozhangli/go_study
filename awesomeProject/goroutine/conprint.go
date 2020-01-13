package main

import (
	"fmt"
	"time"
)

func printer(c chan int){
	for{
		data:=<-c
		if data==0{
			break
		}
		fmt.Println(data)
	}
	time.Sleep(time.Second*10)
	fmt.Println("搞定了")
	c<-0
}

func main(){
	c:=make(chan int)
	go printer(c)

	for i:=1;i<=10;i++{
		c<-i
	}
	fmt.Println("没数据了")
	c<-0
   	<-c
   	fmt.Println("搞定喊我")
}