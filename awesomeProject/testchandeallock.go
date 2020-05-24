package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 5)
	for i := 0; i < 5; i++ {
		ch <- i
	}

	go func() {
		for {
			fmt.Println(<-ch)
		}
	}()
	time.Sleep(10 * time.Second)
}
