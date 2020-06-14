package reader

import (
	"fmt"
	"go_study/Concurrent"
	"testing"
	"time"
)

type taskScheduled struct {
}

func (task *taskScheduled) Run() error {
	ticker := time.NewTicker(time.Second * 5)
	go func() {
		for _ = range ticker.C {
			println("test")
		}
	}()
	return nil
}

//周期性执行任务
func TestScheduled(t *testing.T) {
	rejected := Concurrent.NewRejectedHandler(func() {
		panic("rejected execute goroutine")
	})
	pool := Concurrent.NewPoolRejectedHandler(2, rejected)
	for i := 0; i < 10; i++ {
		pool.Execute(new(taskScheduled))
	}
	fmt.Println(pool.WorkerSize())
	time.Sleep(time.Minute)
	fmt.Println(pool.WorkerSize())
}

//倒计时,5秒后执行
func TestScheduled2(t *testing.T) {
	//主线程阻塞
	timer1 := time.NewTimer(time.Second * 5)
	<-timer1.C
	println("test")

}

//golang 定时器，启动的时候执行一次，以后每天晚上12点执行
func startTimer(f func()) {
	go func() {
		for {
			f()
			now := time.Now()
			// 计算下一个零点
			next := now.Add(time.Hour * 24)
			next = time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, next.Location())
			t := time.NewTimer(next.Sub(now))
			<-t.C
			t.Reset(5)
		}
	}()
}

func TestNewsReaderLoad(t *testing.T) {
	rss := new(RSSDataCapturer)
	rss.load("http://www.matrix67.com/blog/feed.asp")
}

func TestMyNewsReader_NewsReader(t *testing.T) {
	MyNewsReaderStart()
}
