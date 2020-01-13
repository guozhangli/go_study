package main

import (
	"container/list"
	"fmt"
)

func main() {
	l:=list.New()
	l.PushBack("candon")
	l.PushFront(67)
	element:=l.PushBack("fist")
	l.InsertAfter("high",element)
	l.InsertBefore("noon",element)
	for i:=l.Front();i!=nil;i=i.Next(){
		fmt.Println(i.Value)
	}
	l.Remove(element)
	for i:=l.Front();i!=nil;i=i.Next(){
		if i.Value==67 {
			fmt.Println("ok")
		}
		fmt.Println(i.Value)
	}

}
