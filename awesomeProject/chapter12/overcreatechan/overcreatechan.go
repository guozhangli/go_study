package main

import (
	"fmt"
	"runtime"
)

// 一段耗时的计算函数
func consumer(ch chan int) {

	// 无限获取数据的循环
	for {

		// 从通道获取数据
		data := <-ch

		if data==0{
			break
		}
		// 打印数据
		fmt.Println(data)
	}
	fmt.Println("goroutine exit")
}

func main() {

	// 创建一个传递数据用的通道
	ch := make(chan int)

	for {

		// 空变量，什么也不做
		var dummy string

		// 获取输入，模拟进程持续运行
		fmt.Scan(&dummy)

		if dummy=="quit"{
			for i:=0;i<runtime.NumGoroutine()-1;i++{
				ch<-0
			}
			continue
		}
		// 启动并发执行consumer函数u

		go consumer(ch)

		// 输出现在的goroutine数量
		fmt.Println("goroutines:", runtime.NumGoroutine())
	}

}
