package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		fmt.Println("start goroutine")
		ch <- 0
		time.Sleep(time.Second)
		fmt.Println("exit goroutine")
	}()
	fmt.Println("wait goroutine")
	data := <-ch
	fmt.Println(data)
}
