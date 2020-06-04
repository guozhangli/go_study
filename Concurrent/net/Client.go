package net

import (
	"bufio"
	"io"
	"log"
	"math/rand"
	"net"
	"time"
)

func Client() {
	dao := GetDAO()
	list := dao.getData()
	for i := 0; i < 10; i++ {
		conn, err := net.Dial("tcp", "localhost:8000")
		if err != nil {
			log.Println(err)
		}
		log.Println("client connect")
		sendQuery(conn, list)
	}
	for i := 0; i < 10; i++ {
		conn, err := net.Dial("tcp", "localhost:8000")
		if err != nil {
			log.Println(err)
		}
		log.Println("client connect")
		sendReport(conn, list)
	}
	for i := 0; i < 10; i++ {
		conn, err := net.Dial("tcp", "localhost:8000")
		if err != nil {
			log.Println(err)
		}
		log.Println("client connect")
		sendStop(conn)
	}
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Println(err)
	}
}

func sendQuery(conn net.Conn, list []*WDI) {
	defer conn.Close()
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(len(list))
	wdi := list[num]
	var commandData string
	commandData = "q;"
	commandData += wdi.countryCode
	commandData += ";"
	commandData += wdi.indicatorCode
	commandData += "\n"
	io.WriteString(conn, commandData)
	rd := bufio.NewReader(conn)
	line, _, err := rd.ReadLine()
	if err != nil {
		log.Println("read conn io error")
	}
	log.Println(string(line))
}

func sendReport(conn net.Conn, list []*WDI) {
	defer conn.Close()
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(len(list))
	wdi := list[num]
	var commandData string
	commandData = "r;"
	commandData += wdi.indicatorCode
	commandData += "\n"
	io.WriteString(conn, commandData)
	rd := bufio.NewReader(conn)
	line, _, err := rd.ReadLine()
	if err != nil {
		log.Println("read conn io error")
	}
	io.WriteString(conn, "EOF\n")
	/*太快的关闭客户端连接，服务器端报
	An existing connection was forcibly closed by the remote host.
	远程主机强制关闭现有的连接。
	*/
	time.Sleep(time.Second)
	log.Println(string(line))
}

func sendStop(conn net.Conn) {
	defer conn.Close()
	rand.Seed(time.Now().UnixNano())
	var commandData string
	commandData = "z;"
	commandData += "\n"
	io.WriteString(conn, commandData)
	rd := bufio.NewReader(conn)
	line, _, err := rd.ReadLine()
	if err != nil {
		log.Println("read conn io error")
	}
	log.Println(string(line))
}
