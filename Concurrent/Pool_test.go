package Concurrent

import (
	"fmt"
	"log"
	"testing"
	"time"
)

type taskTest struct {
	pool *Pool
}

func (t taskTest) Run() error {
	fmt.Println(time.Now())
	fmt.Println(t.pool.WorkerSize())
	time.Sleep(time.Second)
	return nil
}
func TestNewRejectedHandlerPool(t *testing.T) {
	rejected := NewRejectedHandler(func() {
		log.Println("rejected execute goroutine")
	})
	pool := NewPoolRejectedHandler(500, rejected)
	task := taskTest{
		pool: pool,
	}
	for i := 0; i < 3000; i++ {
		panic("rejected execute goroutine")
	}
	pool.ShutDown()
	pool.Execute(task)
	time.Sleep(30 * time.Second)
}

func TestNewPool(t *testing.T) {
	pool := NewPool(10)
	task := taskTest{
		pool: pool,
	}
	for i := 0; i < 100; i++ {
		pool.Execute(task)
	}
	//pool.ShutDown()
	pool.Execute(task)
	for {
		time.Sleep(time.Second)
		fmt.Println("worker", pool.WorkerSize())
		fmt.Println("max", pool.maxNum)
	}
}

func TestChan(t *testing.T) {
	ch := make(chan int)

	start := time.Now().UnixNano()
	for i := 0; i < 10; i++ {
		func() {

			go func() {
				fmt.Println("send")
				ch <- i
				fmt.Println("xxxx")
			}()

			func() {
				time.Sleep(time.Second)
				fmt.Println("sleep 1 second")
			}()

			func() {
				fmt.Println(<-ch)
			}()
		}()
	}
	end := time.Now().UnixNano()
	fmt.Println(end - start)

	/*for c:=0;c<10;c++{
		//time.Sleep(time.Second)
		func() {
			fmt.Println(<-ch)
		}()
	}*/

	time.Sleep(5 * time.Second)
}
