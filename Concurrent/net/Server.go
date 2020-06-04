package net

import (
	"Concurrent"
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

//串行处理服务
func Server1() {
	_ = GetDAO()
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

//并行处理服务
func Server2() {
	_ = GetDAO()
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
		go handleConn(conn) //每个连接提供一个协程去处理，连接数较多时系统资源消耗较大
	}
}

func Server3() {
	_ = GetDAO()
	listen, err := net.Listen("tcp", "localhost:6666")
	if err != nil {
		log.Println(err)
	}
	log.Println("waiting client connect")
	pool := Concurrent.NewPool(100)
	for {
		conn, err := listen.Accept() //阻塞
		if err != nil {
			log.Println(err)
			continue
		}
		pool.Execute(func() error {
			handleConn(conn)
			return nil
		})
		log.Println("pool worker size：", pool.WorkerSize())
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	rd := bufio.NewReader(conn)
	line, _, err := rd.ReadLine() //阻塞
	if err != nil {
		log.Println(err)
	}
	commandData := strings.Split(string(line), ";")
	fmt.Println("command:", string(line))
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
	res := command.execute()
	res += "\n"
	io.WriteString(conn, res)
}
