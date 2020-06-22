package knn

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

/**
k-邻近算法串行版本
Train: 39129
Test: 2059
******************************************
Serial Classifier - K: 10
Success: 1873
Mistakes: 186
Execution Time: 105 seconds.
******************************************
goos: windows
goarch: amd64
pkg: Concurrent
BenchmarkKnnSerial-8           1        105387500000 ns/op      94845236456 B/op        241788567 allocs/op
PASS
ok      Concurrent      106.793s

*/
func BenchmarkKnnSerial(b *testing.B) {
	KnnSerial()
}

/**
k-邻近算法细粒度并行版本
Train: 39129
Test: 2059
******************************************
Serial Classifier - K: 10
Success: 1873
Mistakes: 186
Execution Time: 48 seconds.
******************************************
goos: windows
goarch: amd64
pkg: Concurrent
BenchmarkKnnParallelIndividual-8               1        48612000000 ns/op
PASS
ok      Concurrent      49.090s

*/
func BenchmarkKnnParallelIndividual(b *testing.B) {
	KnnParallelIndividual()
}

/**
k-邻近算法粗粒度并行
Train: 39129
Test: 2059
******************************************
Serial Classifier - K: 10
Success: 1872
Mistakes: 187
Execution Time: 37 seconds.
******************************************
goos: windows
goarch: amd64
pkg: Concurrent
BenchmarkKnnClassifierParallelGroup-8                  1        37419000000 ns/op
PASS
ok      Concurrent      37.839s
*/
func BenchmarkKnnClassifierParallelGroup(b *testing.B) {
	KnnParallelGroup()
}
