package Concurrent

import (
	"fmt"
	"testing"
)

func TestReadFile(t *testing.T) {
	train, _ := Load("data/bank.data")
	fmt.Printf("%d\n", train.Length())

	test, _ := Load("data/bank.test")
	fmt.Printf("%d\n", test.Length())

}
