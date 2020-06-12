package main

import (
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	for i := 0; i < 10; i++ {
		go func() {
			wg.Add(1)
			go func() {
				time.Sleep(time.Minute)
				wg.Done()
			}()
		}()
	}
	wg.Wait()
}
