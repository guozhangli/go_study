package Concurrent

import (
	"fmt"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	pool := NewPool(10)
	for i := 0; i < 4000; i++ {
		pool.Execute(func() error {
			fmt.Println(time.Now())
			time.Sleep(time.Second)
			return nil
		})
	}
	pool.ShutDown()

	for {
		time.Sleep(time.Second)
		fmt.Println("worker", pool.workerSize())
		fmt.Println("max", pool.maxNum)
	}
}
