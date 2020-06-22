package TestProject

import (
	json2 "encoding/json"
	"fmt"
	"testing"
	"time"
)

func TestLinkedBlockingQueue(t *testing.T) {
	lbq := NewLinkedBlockingQueue(10)
	str, _ := json2.Marshal(lbq)
	t.Log(string(str))
}

func printLBQ(lbq *LinkedBlockingQueue, t *testing.T) {
	str, _ := json2.Marshal(lbq)
	t.Log(string(str))
}
func TestLinkedBlockingQueue_enqueue(t *testing.T) {
	lbq := NewLinkedBlockingQueue(10)
	lbq.enqueue("aaaaa")
	lbq.enqueue("bbbbb")
	lbq.enqueue("ccccc")
	lbq.enqueue("ddddd")
	printLBQ(lbq, t)
}

func TestLinkedBlockingQueue_dequeue(t *testing.T) {
	lbq := NewLinkedBlockingQueue(10)
	lbq.enqueue("aaaaa")
	lbq.enqueue("bbbbb")
	lbq.enqueue("ccccc")
	lbq.enqueue("ddddd")
	printLBQ(lbq, t)
	lbq.dequeue()
	printLBQ(lbq, t)
}

func TestLinkedBlockingQueue_take(t *testing.T) {
	lbq := NewLinkedBlockingQueue(10)
	go lbq.Take()
	time.Sleep(time.Minute)
}

func TestLinkedBlockingQueue_put(t *testing.T) {
	lbq := NewLinkedBlockingQueue(10)
	for i := 0; i < 5; i++ {
		go lbq.Put("aaaaa")
	}
	printLBQ(lbq, t)
	time.Sleep(time.Minute)
}

func TestParallelPutAndTake(t *testing.T) {
	lbq := NewLinkedBlockingQueue(1000)
	go func(l *LinkedBlockingQueue) {
		for {
			d := l.Take()
			fmt.Println(d.(int))
		}

	}(lbq)
	go func(l *LinkedBlockingQueue) {
		for {
			d := l.Take()
			fmt.Println(d.(int))
		}

	}(lbq)
	for i := 0; i < 100; i++ {
		go func(l *LinkedBlockingQueue) {
			for i := 0; i < 100; i++ {
				time.Sleep(1 * time.Second)
				l.Put(i)
			}
		}(lbq)
	}
	for {
		time.Sleep(10 * time.Second)
	}

}
