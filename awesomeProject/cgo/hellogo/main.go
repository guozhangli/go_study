package hellogo

//#include "hello.h"
import "C"

func Test1() {
	C.SayHello1(C.CString("hello world\n"))
}
