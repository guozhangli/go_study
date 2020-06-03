package net

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func Server() {
	_ = GetDAO()
	listen, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("waiting client connect")
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}
		handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	for {
		rd := bufio.NewReader(conn)
		line, _, err := rd.ReadLine()
		if err != nil {
			log.Fatal("read net io error")
		}
		commandData := strings.Split(string(line), ";")
		fmt.Println("command:", commandData[0])
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
		io.WriteString(conn, res)
	}
}
