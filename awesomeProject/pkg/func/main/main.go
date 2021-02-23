package main

import (
	"fmt"
	"go_study/awesomeProject/pkg/func"
)

func main() {
	v1, v2 := _func.Swap(1, 2)
	fmt.Printf("%d\n%d", v1, v2)
}
