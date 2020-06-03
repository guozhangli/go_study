package net

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"time"
)

func Client() {
	dao := GetDAO()
	list := dao.getData()
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("client connect")
	defer conn.Close()
	sendData(conn, list)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

func sendData(conn net.Conn, list []*WDI) {
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(len(list))
	wdi := list[num]
	var commandData string
	commandData = "q;"
	commandData += wdi.countryCode
	commandData += ";"
	commandData += wdi.indicatorCode
	io.WriteString(conn, commandData)
	rd := bufio.NewReader(conn)
	line, _, err := rd.ReadLine()
	if err != nil {
		log.Fatal("read conn io error")
	}
	fmt.Println(line)
}
