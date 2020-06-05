package Concurrent

import (
	TestProject "datastructure"
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
	w.state = RUNNING
	w.f()
	w.state = TERMINATED
}

/**
协程池
*/
type Pool struct {
	jobChan chan interface{}
	maxNum  int
	workers *TestProject.Set
}

func NewPool(num int) *Pool {
	return &Pool{
		jobChan: make(chan interface{}, num),
		maxNum:  num,
		workers: TestProject.NewSet(func(o, n interface{}) bool {
			return o == n
		}),
	}

}
func (p *Pool) Execute(f func() error) {
	task := &task{
		f: f,
	}
	p.jobChan <- task
	go p.run()
}

func (p *Pool) ShutDown() {
	close(p.jobChan)
}
func (p *Pool) WorkerSize() int {
	return p.workers.Size()
}

func (p *Pool) addWorker(task interface{}) {
	ch := make(chan int)
	for {
		ws := p.workers
		go func(w *TestProject.Set, c chan int) {
			c <- w.Size()
		}(ws, ch)
		num := p.maxNum
		var wok *worker
		var iterator = ws.Iterator()
		for iterator.HasNode() {
			wok = iterator.Data().(*worker)
			if wok.state == TERMINATED {
				wok.setTask(task)
				return
			}
			iterator = iterator.NextNode()
		}
		//同步集合的大小
		if <-ch < num {
			wok = newWorker(task)
			ws.Insert(wok)
			wok.run()
			return
		}
	}
}

func (p *Pool) run() {
	for task := range p.jobChan {
		p.addWorker(task)
	}
}
