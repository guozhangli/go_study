package net

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"testing"
	"time"
)

func BenchmarkClient1(b *testing.B) {
	start := time.Now().UnixNano()
	Client()
	end := time.Now().UnixNano()
	time.Sleep(10 * time.Second)
	fmt.Println("串行处理请求的执行时间(ms)：", (end-start)/1000000-10*1000)
}

/**
多协程处理请求的执行时间(ms)： 5585
goos: windows
goarch: amd64
pkg: Concurrent/net
BenchmarkClient2-8   	       1	25585000000 ns/op
PASS
*/
func BenchmarkClient2(b *testing.B) {
	start := time.Now().UnixNano()
	Client()
	end := time.Now().UnixNano()
	time.Sleep(10 * time.Second)
	fmt.Println("多协程处理请求的执行时间(ms)：", (end-start)/1000000-10*1000)
}

func BenchmarkClient3(b *testing.B) {
	start := time.Now().UnixNano()
	Client()
	end := time.Now().UnixNano()
	time.Sleep(10 * time.Second)
	fmt.Println("协程池处理请求的执行时间(ms)：", (end-start)/1000000-10*1000)
}

func TestClient(t *testing.T) {
	for i := 0; i < 1000; i++ {
		go func() {
			conn, err := net.Dial("tcp", "localhost:8888")
			if err != nil {
				log.Println(err)
			}
			go func(conn net.Conn) {
				defer conn.Close()
				io.WriteString(conn, "send to server\n")
				read := bufio.NewReader(conn)
				line, _, err := read.ReadLine()
				if err != nil {
					log.Println(err)
				}
				log.Println(string(line))
			}(conn)
		}()
	}
	for i := 0; i < 1000; i++ {
		go func() {
			conn, err := net.Dial("tcp", "localhost:8888")
			if err != nil {
				log.Println(err)
			}
			go func(conn net.Conn) {
				defer conn.Close()
				io.WriteString(conn, "send to server\n")
				read := bufio.NewReader(conn)
				line, _, err := read.ReadLine()
				if err != nil {
					log.Println(err)
				}
				log.Println(string(line))
			}(conn)
		}()
	}
	for i := 0; i < 1000; i++ {
		go func() {
			conn, err := net.Dial("tcp", "localhost:8888")
			if err != nil {
				log.Println(err)
			}
			go func(conn net.Conn) {
				defer conn.Close()
				io.WriteString(conn, "send to server\n")
				read := bufio.NewReader(conn)
				line, _, err := read.ReadLine()
				if err != nil {
					log.Println(err)
				}
				log.Println(string(line))
			}(conn)
		}()
	}
	select {}
}
