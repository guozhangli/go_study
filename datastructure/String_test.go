package TestProject

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestString_pattern(t *testing.T) {
	index := BruteForceStringMatch("goodgoogle", "google")
	t.Log(index)
}

func TestString_pattern2(t *testing.T) {
	var str MyString = "abaabbabaab"
	var F []int
	F = BuildFArray(str)
	t.Log(F)
}

func TestMyString_IndexKMP(t *testing.T) {
	var b MyString = "abaabaabbabaaabaabbabaab"
	var str MyString = "abaabbabaab"
	index := b.IndexKMP(str)
	t.Log(index)
}

//1. 直接使用运算符
func BenchmarkAddStringWithOperator(b *testing.B) {
	hello := "hello"
	world := "world"
	for i := 0; i < 1000000; i++ {
		_ = hello + "," + world
	}
}

/**
goos: windows
goarch: amd64
pkg: go_study/datastructure
BenchmarkAddStringWithOperator-8        52173912                22.6 ns/op
PASS
ok      go_study/datastructure  2.079s

*/
//golang 里面的字符串都是不可变的，每次运算都会产生一个新的字符串，所以会产生很多临时的无用的字符串，不仅没有用，还会给 gc 带来额外的负担，所以性能比较差

//2. fmt.Sprintf()
func BenchmarkAddStringWithSprintf(b *testing.B) {
	hello := "hello"
	world := "world"
	for i := 0; i < 1000000; i++ {
		_ = fmt.Sprintf("%s,%s", hello, world)
	}
}

/**
goos: windows
goarch: amd64
pkg: go_study/datastructure
BenchmarkAddStringWithSprintf-8          8241756               146 ns/op
PASS

*/
//内部使用 []byte 实现，不像直接运算符这种会产生很多临时的字符串，但是内部的逻辑比较复杂，有很多额外的判断，还用到了 interface，所以性能也不是很好

//3. strings.Join()
func BenchmarkAddStringWithJoin(b *testing.B) {
	hello := "hello"
	world := "world"
	for i := 0; i < 1000000; i++ {
		_ = strings.Join([]string{hello, world}, ",")
	}
}

/**
goos: windows
goarch: amd64
pkg: go_study/datastructure
BenchmarkAddStringWithJoin-8    26086956                40.8 ns/op
PASS
ok      go_study/datastructure  1.497s

*/
//join会先根据字符串数组的内容，计算出一个拼接之后的长度，然后申请对应大小的内存，一个一个字符串填入，在已有一个数组的情况下，这种效率会很高，但是本来没有，去构造这个数据的代价也不小

//4. buffer.WriteString()
func BenchmarkAddStringWithBuffer(b *testing.B) {
	hello := "hello"
	world := "world"
	for i := 0; i < 1000000; i++ {
		var buffer bytes.Buffer
		buffer.WriteString(hello)
		buffer.WriteString(",")
		buffer.WriteString(world)
		_ = buffer.String()
	}
}

/**
goos: windows
goarch: amd64
pkg: go_study/datastructure
BenchmarkAddStringWithBuffer-8          20689654                54.1 ns/op
PASS
ok      go_study/datastructure  1.776s

*/
//这个比较理想，可以当成可变字符使用，对内存的增长也有优化，如果能预估字符串的长度，还可以用 buffer.Grow() 接口来设置 capacity

//>go test -run=none -bench=BenchmarkAddStringWith*
