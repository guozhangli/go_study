package net

import (
	"Concurrent"
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"sync"
)

var cache *ParallelCache

var sp sync.Mutex

//串行处理服务
func ServerSerial() {
	_ = GetDAO()
	cache = NewParallelCache()
	listen, err := net.Listen("tcp", "localhost:6666")
	if err != nil {
		log.Println(err)
	}
	log.Println("waiting client connect")
	for {
		conn, err := listen.Accept() //阻塞
		if err != nil {
			log.Println(err)
			continue
		}
		handleConn(conn)
	}
}

var stoped = false

//并行处理服务
func ServerParallel() {
	_ = GetDAO()
	cache = NewParallelCache()
	listen, err := net.Listen("tcp", "localhost:6666")
	if err != nil {
		log.Println(err)
	}
	log.Println("waiting client connect")
	for !stoped {
		conn, err := listen.Accept() //阻塞
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConn(conn) //每个连接提供一个协程去处理，连接数较多时系统资源消耗较大
	}
}

func ServerPool() {
	_ = GetDAO()
	cache = NewParallelCache()
	listen, err := net.Listen("tcp", "127.0.0.1:6666")
	if err != nil {
		log.Println(err)
	}
	log.Println("waiting client connect")
	rejected := Concurrent.NewRejectedHandler(func() {
		log.Println("rejected execute goroutine")
	})
	pool := Concurrent.NewPoolRejectedHandler(100, rejected)
	for !stoped {
		conn, err := listen.Accept() //阻塞
		if err != nil {
			log.Println(err)
			continue
		}
		pool.Execute(func() error {
			handleConn(conn)
			return nil
		})
	}
	if stoped {
		pool.ShutDown() //goroutine pool close
	}
	log.Println("pool worker size：", pool.WorkerSize(), "cache size:", cache.getItemCount())
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	rd := bufio.NewReader(conn)
	line, _, err := rd.ReadLine() //阻塞
	if err != nil {
		return
	}
	commandData := strings.Split(string(line), ";")
	fmt.Println("command:", string(line))
	var res string
	if v, ok := cache.get(string(line)); ok {
		res = v
	} else {
		var command ICommand
		switch commandData[0] {
		case "q":
			command = NewQueryCommand(commandData)
			break
		case "r":
			command = NewReportCommand(commandData)
			break
		case "z":
			command = NewStopCommand(commandData)
			break
		default:
			command = NewErrorCommand(commandData)
		}
		res = command.execute()
		cache.put(string(line), res)
	}
	res += "\n"
	io.WriteString(conn, res)
}

func SetStoped() {
	defer sp.Unlock()
	sp.Lock()
	stoped = true
	cache.shutDown()
}
