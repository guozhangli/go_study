package Concurrent

import (
	"fmt"
	"log"
	"sync"
	"testing"
	"time"
)

func TestCountDownLatchChan(t *testing.T) {
	c := make(chan rune, 2)
	go func() {
		fmt.Println("frist goroutine begin...")
		for i := 0; i < 10; i++ {
			for _, r := range `-\|/` {
				//fmt.Println("what happend")
				log.Printf("\r%c", r)
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

var wg1 sync.WaitGroup

func TestCountDownLatchWaitGroup(t *testing.T) {
	wg1.Add(2)
	go func() {
		defer wg1.Done()
		fmt.Println("frist goroutine begin...")
		for i := 0; i < 10; i++ {
			for _, r := range `-\|/` {
				//fmt.Println("what happend")
				log.Printf("\r%c", r)
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()

	go func() {
		defer wg1.Done()
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
	}()

	wg1.Wait()
	fmt.Println("main function end")
}

func TestWaitGroup(t *testing.T) {
	for i := 0; i < 11; i++ {
		go run2()
	}
	for i := 0; i < 10; i++ {
		add()
	}
	wg1.Wait()
	log.Println("main end")
}

func add() {
	wg1.Add(1)
	go run()
}
func run() {
	time.Sleep(2 * time.Second)
	log.Println("run exec")
}

func run2() {
	defer wg1.Done()
	time.Sleep(time.Second)
	log.Println("run2 exec")
}
