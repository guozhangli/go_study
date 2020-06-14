package Concurrent

type FutureTask struct {
	result interface{}
	isDone bool
	ch     chan interface{}
}

func (task *FutureTask) get() interface{} {
	if !task.isDone {
		var r interface{}
		for r == nil {
			select {
			case r = <-task.ch:
				return r
			default:
				r = nil
			}
		}
	}
	return task.result
}

func (task *FutureTask) finish(result interface{}) {
	if task.isDone {
		return
	}
	task.result = result
	task.isDone = true
	task.ch <- result
}

type FutureService struct {
	futureTask *FutureTask
}

func NewFutureService() *FutureService {
	ft := &FutureTask{
		ch: make(chan interface{}),
	}
	fs := &FutureService{futureTask: ft}
	return fs
}

type Future interface {
	Call() interface{}
}

func (future *FutureService) submit(f Future) *FutureTask {
	go func(ft *FutureTask) {
		result := f.Call()
		ft.finish(result)
	}(future.futureTask)
	return future.futureTask
}
