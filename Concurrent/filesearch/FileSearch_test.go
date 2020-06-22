package filesearch

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
	b.Log((end - start) / 1000000) //ms
	str, _ := json.Marshal(result)
	b.Log(string(str))
}

func BenchmarkFileSearchParallel(b *testing.B) {
	result := NewResult()
	start := time.Now().UnixNano()
	FileSearchParallel("C:\\Windows\\", "hosts", result)
	end := time.Now().UnixNano()
	b.Log((end - start) / 1000000) //ms
	str, _ := json.Marshal(result)
	b.Log(string(str))
}

/**
速度非常快
BenchmarkFileSearchSerial-8             1000000000               0.190 ns/op
--- BENCH: BenchmarkFileSearchSerial-8
    FileSearch_test.go:14: 211
    FileSearch_test.go:16: {"Path":"C:\\Windows\\System32\\drivers\\etc\\hosts","Found":true}
    FileSearch_test.go:14: 193
    FileSearch_test.go:16: {"Path":"C:\\Windows\\System32\\drivers\\etc\\hosts","Found":true}
    FileSearch_test.go:14: 191
    FileSearch_test.go:16: {"Path":"C:\\Windows\\System32\\drivers\\etc\\hosts","Found":true}
    FileSearch_test.go:14: 195
    FileSearch_test.go:16: {"Path":"C:\\Windows\\System32\\drivers\\etc\\hosts","Found":true}
    FileSearch_test.go:14: 205
    FileSearch_test.go:16: {"Path":"C:\\Windows\\System32\\drivers\\etc\\hosts","Found":true}
        ... [output truncated]
BenchmarkFileSearchParallel-8           1000000000               0.0970 ns/op
--- BENCH: BenchmarkFileSearchParallel-8
    FileSearch_test.go:24: 124
    FileSearch_test.go:26: {"Path":"C:\\Windows\\System32\\drivers\\etc\\hosts","Found":true}
    FileSearch_test.go:24: 105
    FileSearch_test.go:26: {"Path":"C:\\Windows\\System32\\drivers\\etc\\hosts","Found":true}
    FileSearch_test.go:24: 104
    FileSearch_test.go:26: {"Path":"C:\\Windows\\System32\\drivers\\etc\\hosts","Found":true}
    FileSearch_test.go:24: 99
    FileSearch_test.go:26: {"Path":"C:\\Windows\\System32\\drivers\\etc\\hosts","Found":true}
    FileSearch_test.go:24: 101
    FileSearch_test.go:26: {"Path":"C:\\Windows\\System32\\drivers\\etc\\hosts","Found":true}

*/
