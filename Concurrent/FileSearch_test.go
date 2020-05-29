package Concurrent

import (
	"encoding/json"
	"testing"
	"time"
)

func BenchmarkFileSearchSerial(b *testing.B) {
	result := NewResult()
	start := time.Now().UnixNano()
	FileSearchSerial("C:\\Windows\\", "hosts", result)
	end := time.Now().UnixNano()
	b.Log((end - start) / 1000000)
	str, _ := json.Marshal(result)
	b.Log(string(str))
}

func BenchmarkFileSearchParallel(b *testing.B) {
	result := NewResult()
	start := time.Now().UnixNano()
	FileSearchParallel("C:\\Windows\\", "hosts", result)
	end := time.Now().UnixNano()
	b.Log((end - start) / 1000000)
	str, _ := json.Marshal(result)
	b.Log(string(str))
}

/**
速度非常快
BenchmarkFileSearchSerial-8             1000000000               0.192 ns/op
--- BENCH: BenchmarkFileSearchSerial-8
    FileSearch_test.go:14: 212
    FileSearch_test.go:16: {"Path":"C:\\Windows\\System32\\drivers\\etc\\hosts","Found":true}
    FileSearch_test.go:14: 218
    FileSearch_test.go:16: {"Path":"C:\\Windows\\System32\\drivers\\etc\\hosts","Found":true}
    FileSearch_test.go:14: 196
    FileSearch_test.go:16: {"Path":"C:\\Windows\\System32\\drivers\\etc\\hosts","Found":true}
    FileSearch_test.go:14: 190
    FileSearch_test.go:16: {"Path":"C:\\Windows\\System32\\drivers\\etc\\hosts","Found":true}
    FileSearch_test.go:14: 195
    FileSearch_test.go:16: {"Path":"C:\\Windows\\System32\\drivers\\etc\\hosts","Found":true}
        ... [output truncated]
BenchmarkFileSearchParallel-8           1000000000               0.196 ns/op
--- BENCH: BenchmarkFileSearchParallel-8
    FileSearch_test.go:24: 195
    FileSearch_test.go:26: {"Path":"C:\\Windows\\System32\\drivers\\etc\\hosts","Found":true}
    FileSearch_test.go:24: 202
    FileSearch_test.go:26: {"Path":"C:\\Windows\\System32\\drivers\\etc\\hosts","Found":true}
    FileSearch_test.go:24: 200
    FileSearch_test.go:26: {"Path":"C:\\Windows\\System32\\drivers\\etc\\hosts","Found":true}
    FileSearch_test.go:24: 204
    FileSearch_test.go:26: {"Path":"C:\\Windows\\System32\\drivers\\etc\\hosts","Found":true}
    FileSearch_test.go:24: 201
    FileSearch_test.go:26: {"Path":"C:\\Windows\\System32\\drivers\\etc\\hosts","Found":true}

*/
