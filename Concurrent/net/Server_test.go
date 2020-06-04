package net

import (
	"bufio"
	"io"
	"log"
	"net"
	"testing"
)

//串行处理服务
func BenchmarkServer1(b *testing.B) {
	Server1()
}

//每个连接一个协程处理
func BenchmarkServer2(b *testing.B) {
	Server2()
}

//协程池处理
func BenchmarkServer3(b *testing.B) {
	Server3()
}

func TestServer(t *testing.T) {
	listen, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		log.Fatal("listen is error")
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal("connect is error")
			continue
		}
		go handle(conn)
	}

}

func handle(conn net.Conn) {
	defer conn.Close()
	for {
		read := bufio.NewReader(conn)
		line, _, err := read.ReadLine()
		if err != nil {
			log.Println(err)
			break
		}
		log.Println(string(line))
		io.WriteString(conn, "send to client\n")
	}

}
