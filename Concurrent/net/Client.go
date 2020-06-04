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
	l := dao.getData()
	for i := 0; i < 1000; i++ {
		go func(list []*WDI) {
			conn, err := net.Dial("tcp", "localhost:6666")
			if err != nil {
				log.Println(err)
			}
			if conn == nil {
				log.Println("conn error")
				return
			}
			log.Println("client connect")
			sendQuery(conn, list)
		}(l)

	}
	for i := 0; i < 1000; i++ {
		go func(list []*WDI) {
			conn, err := net.Dial("tcp", "localhost:6666")
			if err != nil {
				log.Println(err)
			}
			if conn == nil {
				log.Println("conn error")
				return
			}
			log.Println("client connect")
			sendReport(conn, list)
		}(l)
	}
	for i := 0; i < 1; i++ {
		func() {
			conn, err := net.Dial("tcp", "localhost:6666")
			if err != nil {
				log.Println(err)
			}
			if conn == nil {
				log.Println("conn error")
				return
			}
			log.Println("client connect")
			sendStop(conn)
		}()
	}
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Println(err)
	}
}

func sendQuery(conn net.Conn, list []*WDI) {
	if conn != nil {
		defer conn.Close()
	}
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(len(list))
	wdi := list[num]
	var commandData string
	commandData = "q;"
	commandData += wdi.countryCode
	commandData += ";"
	commandData += wdi.indicatorCode
	commandData += "\n"
	_, err := io.WriteString(conn, commandData)
	if err != nil {
		log.Println(err)
	}
	rd := bufio.NewReader(conn)
	line, _, err := rd.ReadLine()
	if err != nil {
		log.Println(err)
	}
	log.Println(string(line))
}

func sendReport(conn net.Conn, list []*WDI) {
	if conn != nil {
		defer conn.Close()
	}
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(len(list))
	wdi := list[num]
	var commandData string
	commandData = "r;"
	commandData += wdi.indicatorCode
	commandData += "\n"
	_, err := io.WriteString(conn, commandData)
	if err != nil {
		log.Println(err)
	}
	rd := bufio.NewReader(conn)
	line, _, err := rd.ReadLine()
	if err != nil {
		log.Println(err)
	}
	log.Println(string(line))
}

func sendStop(conn net.Conn) {
	if conn != nil {
		defer conn.Close()
	}
	rand.Seed(time.Now().UnixNano())
	var commandData string
	commandData = "z;"
	commandData += "\n"
	_, err := io.WriteString(conn, commandData)
	if err != nil {
		log.Println(err)
	}
	rd := bufio.NewReader(conn)
	line, _, err := rd.ReadLine()
	if err != nil {
		log.Println(err)
	}
	log.Println(string(line))
}
