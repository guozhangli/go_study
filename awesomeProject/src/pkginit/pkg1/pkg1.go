package pkg1

import (
	"fmt"
	"go_study/awesomeProject/src/pkginit/pkg2"
)

func ExecPkg1() {

	fmt.Println("ExecPkg1")

	pkg2.ExecPkg2()
}

func init() {
	fmt.Println("pkg1 init")
}
