package Concurrent

import (
	TestProject "go_study/datastructure"
	"sync"
)

type task struct {
	f func() error
}

const (
	RUNNING    = 1
	TERMINATED = 2
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
	}
	return w
}
func (w *worker) setTask(f interface{}) {
	fn := f.(*task).f
	w.f = func() {
		fn()
	}
	w.run()
}

func (w *worker) run() {
	defer wg.Done()
	w.state = RUNNING
	w.f()
	w.state = TERMINATED
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
	maxNum   int
	workers  *TestProject.Set
	rejected *RejectedHandler
	closed   bool
}

func NewPool(num int) *Pool {
	return &Pool{
		jobChan: make(chan interface{}, num),
		maxNum:  num,
		workers: TestProject.NewSet(func(o, n interface{}) bool {
			return o == n
		}),
		rejected: nil,
		closed:   false,
	}
}

func NewPoolRejectedHandler(num int, rejectedHandler *RejectedHandler) *Pool {
	return &Pool{
		jobChan: make(chan interface{}, num),
		maxNum:  num,
		workers: TestProject.NewSet(func(o, n interface{}) bool {
			return o == n
		}),
		rejected: rejectedHandler,
		closed:   false,
	}
}

var wg sync.WaitGroup

func (p *Pool) Execute(f func() error) {
	if p.closed {
		if p.rejected != nil {
			p.rejected.reject()
		}
		return
	}
	task := &task{
		f: f,
	}
	p.jobChan <- task
	go p.run()
}

var mu sync.Mutex

func (p *Pool) ShutDown() {
	defer mu.Unlock()
	mu.Lock()
	wg.Wait()
	p.closed = true
	p.workers.Clear()
	close(p.jobChan)
}
func (p *Pool) WorkerSize() int {
	return p.workers.Size()
}

var total int

func (p *Pool) addWorker(task interface{}) {
	ch := make(chan int)
	for {
		ws := p.workers
		num := p.maxNum
		go func() {
			ch <- ws.Size()
		}()
		var wok *worker
		var iterator = ws.Iterator()
		for iterator.HasNode() {
			wok = iterator.Data().(*worker)
			if wok.state == TERMINATED {
				go wok.setTask(task)
				return
			}
			iterator = iterator.NextNode()
		}

		select {
		case s := <-ch:
			mu.Lock()
			if s < num {
				mu.Unlock()
				wok = newWorker(task)
				ws.Insert(wok)
				go wok.run()
				return
			}

		}

	}
}

func (p *Pool) run() {
	for task := range p.jobChan {
		wg.Add(1)
		p.addWorker(task)
	}
}
