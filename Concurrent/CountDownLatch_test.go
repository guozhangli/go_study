package Concurrent

import (
	"fmt"
	"testing"
	"time"
)

func TestCountDownLatch(t *testing.T) {
	c := make(chan rune, 2)
	go func() {
		fmt.Println("frist goroutine begin...")
		for i := 0; i < 10; i++ {
			for _, r := range `-\|/` {
				//fmt.Println("what happend")
				fmt.Printf("\r%c", r)
				time.Sleep(100 * time.Millisecond)
			}
		}
		c <- 1
	}()

	go func() {
		fmt.Println("second goroutine begin...")
		var fib func(int) int
		fib = func(n int) int {
			if n < 2 {
				return n
			}
			return fib(n-1) + fib(n-2)
		}
		const n = 45
		fibN := fib(45)
		fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
		c <- 1
	}()

	for i := 0; i < 2; i++ {
		r := <-c
		fmt.Println(r)
	}

	fmt.Println("main function end")
}
