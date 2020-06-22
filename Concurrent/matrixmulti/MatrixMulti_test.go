package matrixmulti

import (
	"testing"
)

func TestGenerate(t *testing.T) {
	m := generate(10, 20)
	t.Log(m)
}

func BenchmarkSerialMultiply(b *testing.B) {
	a1 := generate(2000, 2000)
	a2 := generate(2000, 2000)
	serial_multiply(a1, a2)
}

func BenchmarkParallelRowMultiply(b *testing.B) {
	a1 := generate(2000, 2000)
	a2 := generate(2000, 2000)
	parallel_row_multiply(a1, a2)
}

func BenchmarkParallelIndividualMultiply(b *testing.B) {
	a1 := generate(2000, 2000)
	a2 := generate(2000, 2000)
	parallel_individual_multiply(a1, a2)
}

func BenchmarkParallelGroupMultiply(b *testing.B) {
	a1 := generate(2000, 2000)
	a2 := generate(2000, 2000)
	parallel_group_multiply(a1, a2)
}

/**
两次基准测试结果：
BenchmarkSerialMultiply-8                              1        155548000000 ns/op
BenchmarkParallelRowMultiply-8                         1        67171000000 ns/op
BenchmarkParallelIndividualMultiply-8                  1        98711000000 ns/op
BenchmarkParallelGroupMultiply-8                       1        67294500000 ns/op

BenchmarkSerialMultiply-8                              1        154727500000 ns/op
BenchmarkParallelRowMultiply-8                         1        67295500000 ns/op
BenchmarkParallelIndividualMultiply-8                  1        97023500000 ns/op
BenchmarkParallelGroupMultiply-8                       1        67085000000 ns/op
*/
