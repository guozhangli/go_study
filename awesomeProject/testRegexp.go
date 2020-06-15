package main

import (
	"fmt"
	"regexp"
)

func main() {
	reg := regexp.MustCompile("\\P{L}+")
	flag := reg.Split("abc defe", -1)
	fmt.Printf("%s", flag)
}
