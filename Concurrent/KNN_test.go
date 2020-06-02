package Concurrent

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

func TestReadFile(t *testing.T) {
	train, _ := Load("data/bank.data")
	fmt.Printf("%d\n", len(train))
	test, _ := Load("data/bank.test")
	fmt.Printf("%d\n", len(test))

}

func TestCreateFile(t *testing.T) {
	data, _ := ioutil.ReadFile("data/test.data")
	for _, line := range strings.Split(string(data), "\n") {
		s := strings.TrimSpace(line)
		fmt.Println("bytes2Float64(" + strings.TrimSpace(strings.Split(s, "=")[0]) + "),")
	}
}

func BenchmarkKnnSerial(b *testing.B) {
	KnnSerial()
}

/**
Train: 39129
Test: 2059
******************************************
Serial Classifier - K: 10
Success: 1869
Mistakes: 190
Execution Time: 105 seconds.
******************************************
goos: windows
goarch: amd64
pkg: Concurrent
BenchmarkKnnSerial-8           1        105387500000 ns/op      94845236456 B/op        241788567 allocs/op
PASS
ok      Concurrent      106.793s

*/
