package net

import (
	"bufio"
	"fmt"
	"go_study/Concurrent"
	"io"
	"log"
	"net"
	"testing"
)

//串行处理服务
func BenchmarkServer1(b *testing.B) {
	ServerSerial()
}

//每个连接一个协程处理
func BenchmarkServer2(b *testing.B) {
	ServerParallel()
}

//协程池处理
func BenchmarkServer3(b *testing.B) {
	ServerPool()
}

type taskServerTest struct {
	conn net.Conn
}

func (t *taskServerTest) Run() error {
	defer t.conn.Close()
	for {
		read := bufio.NewReader(t.conn)
		line, _, err := read.ReadLine()
		if err != nil {
			break
		}
		log.Println(string(line))
		io.WriteString(t.conn, "send to client\n")
	}
	return nil
}

func TestServer(t *testing.T) {
	listen, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		log.Fatal("listen is error")
	}
	pool := Concurrent.NewPool(100)
	for {
		c, err := listen.Accept()
		if err != nil {
			log.Fatal("connect is error")
			continue
		}
		task := &taskServerTest{conn: c}
		pool.Execute(task)
		fmt.Println(pool.WorkerSize())
	}
}
