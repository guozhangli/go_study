package Concurrent

import "testing"

func TestGenerate(t *testing.T) {
	m := generate(10, 20)
	t.Log(m)
}

func TestMulti(t *testing.T) {
	a := generate(2, 3)
	t.Log(a)
	b := generate(3, 2)
	t.Log(b)
	c := multi(a, b)
	t.Log(c)
}
