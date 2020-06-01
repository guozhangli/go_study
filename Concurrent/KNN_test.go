package Concurrent

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

func TestReadFile(t *testing.T) {
	train, _ := Load("data/bank.data")
	fmt.Printf("%d\n", train.Length())
	test, _ := Load("data/bank.test")
	fmt.Printf("%d\n", test.Length())

}

func TestCreateFile(t *testing.T) {
	data, _ := ioutil.ReadFile("data/test.data")
	for _, line := range strings.Split(string(data), "\n") {
		s := strings.TrimSpace(line)
		fmt.Println("bytes2Float64(" + strings.TrimSpace(strings.Split(s, "=")[0]) + "),")
	}
}

func TestKnn(t *testing.T) {
	Knn()
}
