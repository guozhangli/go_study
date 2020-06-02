package net

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func Server() {
	listen, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
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
	b := make([]byte, 256)
	for {
		go func(bytes []byte) {
			if _, err := conn.Read(bytes); err != nil {
				log.Fatal(err)
				return
			}
		}(b)

		_, err := io.WriteString(conn, time.Now().Format(fmt.Sprintf("15:04:05%s\n", string(b))))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
