package main

import (
	TestProject "go_study/datastructure"
)
// 定义任务类型
type Task struct {
	f func() error
}
var queue =TestProject.NewQueueLinked()

func Take(arg_f func() error){
	task:=&Task{
		f:arg_f,
	}
	queue.EnQueue(task)
}

func (task *Task) Execute(){
	task.f()
}


