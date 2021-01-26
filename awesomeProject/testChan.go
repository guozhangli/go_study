package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)
	func() {
		for i := 0; i < 15; i++ {
			go sent(ch, i)
		}
	}()
	var wg sync.WaitGroup
	wg.Add(15)
	func() {

		for i := 0; i < 15; i++ {
			go rec(ch, &wg)
		}

	}()
	wg.Wait()
}

func sent(ch chan int, i int) {
	ch <- i
}

func rec(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	j := <-ch
	fmt.Println(j)
}
