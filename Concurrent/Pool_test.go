package Concurrent

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
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
	pool := NewPoolRejectedHandler(100, rejected)
	task := taskTest{
		pool: pool,
	}
	for i := 0; i < 1000; i++ {
		pool.Execute(task)
	}
	pool.ShutDown()
	pool.Execute(task)
	pool.WaitTermination()
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

type taskPoolFuture struct {
}

func (task *taskPoolFuture) Call() interface{} {
	//time.Sleep(2*time.Second)
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(1000)
	return "hello world " + strconv.Itoa(r)
}

func TestPool_Submit(t *testing.T) {
	rejected := NewRejectedHandler(func() {
		log.Println("rejected execute goroutine")
	})
	pool := NewPoolRejectedHandler(10, rejected)
	var fts []*FutureTask
	for i := 0; i < 100; i++ {
		future := NewFutureService()
		taskPool := new(taskPoolFuture)
		ft := pool.Submit(taskPool, future)
		fts = append(fts, ft)
	}
	for _, v := range fts {
		fmt.Printf("%s\n", v.get())
	}
	pool.ShutDown()
	pool.WaitTermination()
	//time.Sleep(20 * time.Second)
}
