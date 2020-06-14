package Concurrent

import (
	"fmt"
	"testing"
	"time"
)

type taskFuture struct {
}

func (t *taskFuture) Call() interface{} {
	time.Sleep(5 * time.Second)
	return "hello world"
}

func TestFutureService(t *testing.T) {
	future := NewFutureService()
	task := new(taskFuture)
	ft := future.submit(task)
	fmt.Printf("%s", ft.get())
}
