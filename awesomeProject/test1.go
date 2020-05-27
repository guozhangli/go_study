package main

import (
	"fmt"
	"math/big"
	"time"
)

func dummy(b int) int {
	var c int
	c = b
	return c
}

func void() {

}

func main() {
	var a int
	void()
	fmt.Println(a, dummy(0))
	t := time.Now().Unix()
	fmt.Printf("%32b\n", (1<<24)-1)
	fmt.Printf("%32b\n", t)
	fmt.Printf("%32b\n", (t)&((1<<24)-1))
	fmt.Printf("%d\n", 1<<16)
	var i = big.NewInt(2)
	i.Exp(i, big.NewInt(32), nil)
	fmt.Printf("%d\n", i)
	var b = i.Int64() / (4 * 1024)
	fmt.Printf("%d\n", b)

	fmt.Printf("%d\n", b*4)
	fmt.Printf("%d\n", 4*1024*1024)
}
