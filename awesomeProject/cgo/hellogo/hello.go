package hellogo

import "C"
import "fmt"

//export SayHello1
func SayHello1(s *C.char) {
	fmt.Print(C.GoString(s))
}
