package main

import (
	"errors"
	"fmt"
	"time"
)

func RPCClient(ch chan string,req string)(string,error){
	ch<-req
	select {
	case ack:=<-ch:
		return ack,nil
	case <-time.After(time.Second):
		return "",errors.New("Time out")
	}
}

func RPCSever(ch chan string){
	for {
		data:=<-ch
		fmt.Println("server received",data)
		time.Sleep(time.Second*2)
		ch<-"roger"
	}
}

func main(){
	ch:=make(chan string)
	go RPCSever(ch)
	recv,err:=RPCClient(ch,"hi")
	if err!=nil{
		fmt.Println(err)
	}else {
		fmt.Println("client received",recv)
	}
}