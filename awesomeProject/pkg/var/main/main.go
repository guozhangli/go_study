package main

import (
	"fmt"
	"go_study/awesomeProject/pkg/var"
)

func main() {
	fmt.Println(_var.Id)
	fmt.Println(_var.Name)
	_var.NameData[0] = '?'
	fmt.Println(_var.Name)
	fmt.Println(_var.Name1)
}
