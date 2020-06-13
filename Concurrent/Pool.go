package Concurrent

import (
	TestProject "go_study/datastructure"
	"log"
	"sync"
	"sync/atomic"
)

type task struct {
	f func() error
}

const (
	RUNNABLE = 0 //就绪状态
	RUNNING  = 1 //运行状态
)

var (
	muWork sync.Mutex
	muStop sync.Mutex
	wg     sync.WaitGroup
)

type worker struct {
	f     func()
	state int
}

func newWorker(f interface{}) *worker {
	fn := f.(*task).f
	w := &worker{
		f: func() {
			fn()
		},
		state: RUNNABLE,
	}
	return w
}
func (w *worker) setTask(f interface{}) {
	fn := f.(*task).f
	w.f = func() {
		fn()
	}
}

func (w *worker) run() {
	defer wg.Done()
	w.state = RUNNING
	w.f()
	w.state = RUNNABLE
}

type RejectedHandler struct {
	reject func()
}

func NewRejectedHandler(f func()) *RejectedHandler {
	rejectedHandler := new(RejectedHandler)
	rejectedHandler.reject = f
	return rejectedHandler
}

/**
协程池
*/
type Pool struct {
	jobChan  chan interface{}
	maxNum   int32
	workers  *TestProject.Set
	rejected *RejectedHandler
	closed   bool
}

func NewPool(num int32) *Pool {
	pool := &Pool{
		jobChan: make(chan interface{}, num),
		maxNum:  num,
		workers: TestProject.NewSet(func(o, n interface{}) bool {
			return o == n
		}),
		rejected: nil,
		closed:   false,
	}
	return pool
}

func NewPoolRejectedHandler(num int32, rejectedHandler *RejectedHandler) *Pool {
	pool := &Pool{
		jobChan: make(chan interface{}, num),
		maxNum:  num,
		workers: TestProject.NewSet(func(o, n interface{}) bool {
			return o == n
		}),
		rejected: rejectedHandler,
		closed:   false,
	}
	return pool
}

type Runnable interface {
	Run() error
}

func (p *Pool) Execute(r Runnable) {
	if p.closed {
		if p.rejected != nil {
			p.rejected.reject()
		}
		return
	}
	task := &task{
		f: r.Run,
	}
	p.jobChan <- task
	go p.addTask()
}

func (p *Pool) ShutDown() {
	muStop.Lock()
	p.closed = true  //设置关闭标志位
	close(p.jobChan) //关闭工作任务队列
	muStop.Unlock()
	func() {
		wg.Wait() //等待所有的协程执行完毕
		log.Println("all goroutine execute finish")
		p.workers.Clear() //清空工作集合
		log.Printf("worker num:%d\n", p.WorkerSize())
		p.workers = nil
	}()
}

func (p *Pool) WorkerSize() int {
	return p.workers.Size()
}

var total = int32(0)

func (p *Pool) addWorker(task interface{}) {
	var wok *worker
	ws := p.workers
	num := p.maxNum
	for {
		var iterator = ws.Iterator()
		for iterator.HasNode() {
			it, flag := func(it *TestProject.LinkedList, task interface{}) (*TestProject.LinkedList, bool) {
				defer muWork.Unlock()
				muWork.Lock()
				wok = iterator.Data().(*worker)
				if wok.state == RUNNABLE {
					wok.setTask(task)
					wok.run()
					return iterator, true
				}
				return iterator.NextNode(), false
			}(iterator, task)
			if flag {
				return
			} else {
				iterator = it
			}
		}
		for {
			count := atomic.LoadInt32(&total)
			if atomic.CompareAndSwapInt32(&total, count, count+1) {
				break
			}
		}
		if atomic.LoadInt32(&total) <= num {
			wok = newWorker(task)
			ws.Insert(wok)
			wok.run()
			return
		}
	}
}

func (p *Pool) addTask() {
	for task := range p.jobChan {
		wg.Add(1)
		p.addWorker(task)
	}
}
