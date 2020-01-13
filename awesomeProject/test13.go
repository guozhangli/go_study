package main

import "fmt"

func main() {
	defer fmt.Println("宕机后要做1")
	defer fmt.Println("宕机后要做2")
	panic("crash")
}
