package kMeansClustering

import (
	"fmt"
	"os"
	"strconv"
	"testing"
)

func TestLoad(t *testing.T) {
	m := loadfile2Map()
	for k, v := range m {
		fmt.Printf("key:%s,value:%d\n", k, v)
	}
	docs := loadfile2Slice(m)
	for _, doc := range docs {
		fmt.Printf("%v\n", doc)
	}
}

func TestSerialKMeansClustering(t *testing.T) {
	K := os.Args[4]
	seed := os.Args[5]
	Ki, _ := strconv.Atoi(K)
	seedi, _ := strconv.Atoi(seed)
	t.Log(Ki, seedi)
	SerialKMeansClustering(Ki, seedi)
}
