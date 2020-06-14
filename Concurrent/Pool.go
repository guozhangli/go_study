package Concurrent

import (
	TestProject "go_study/datastructure"
	"log"
	"sync"
)

type task struct {
	f func() error
}

var (
	muStop sync.Mutex
	wg     sync.WaitGroup
)

type worker struct {
	f func()
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
	go pool.start()
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
	go pool.start()
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
	wg.Add(1)
	p.jobChan <- task
}

func (p *Pool) Submit(f Future, future *FutureService) *FutureTask {
	if p.closed {
		if p.rejected != nil {
			p.rejected.reject()
		}
		return nil
	}
	task := &task{
		f: func() error {
			go func(ft *FutureTask) {
				result := f.Call()
				ft.finish(result)
			}(future.futureTask)
			return nil
		},
	}
	wg.Add(1)
	go func() {
		p.jobChan <- task
	}()
	return future.futureTask
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

func (p *Pool) start() {
	//初始化，添加工作者
	p.addWorker()
}

func (p *Pool) addWorker() {
	ws := p.workers
	num := p.maxNum
	for ws.Size() < int(num) {
		wok := new(worker)
		ws.Insert(wok)
		go wok.run(p)
	}
}

func (w *worker) setTask(f interface{}) {
	fn := f.(*task).f
	w.f = func() {
		fn()
	}
}

func (w *worker) run(p *Pool) {
	for task := range p.jobChan {
		w.setTask(task)
		w.f() //执行任务
		wg.Done()
	}
}
