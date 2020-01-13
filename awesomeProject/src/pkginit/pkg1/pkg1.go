package pkg1

import (
	"awesomeProject/src/pkginit/pkg2"
	"fmt"
)

func ExecPkg1() {

	fmt.Println("ExecPkg1")

	pkg2.ExecPkg2()
}

func init() {
	fmt.Println("pkg1 init")
}
