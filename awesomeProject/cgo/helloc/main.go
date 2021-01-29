package main

//#include "hello.h"
import "C"
import "go_study/awesomeProject/cgo/hellogo"

func main() {
	C.SayHello(C.CString("hello world\n"))
	hellogo.Test1()
}
