package Concurrent

import (
	"math/rand"
)

func generate(row, col int) [][]int {
	matrix := make([][]int, row)
	for i := 0; i < row; i++ {
		v := make([]int, col)
		matrix[i] = v
		for j := 0; j < col; j++ {
			num := rand.Intn(10)
			matrix[i][j] = num
		}
	}
	return matrix
}

func multi(a, b [][]int) [][]int {
	row_a := len(a)
	col_b := len(b[0])
	if row_a != col_b {
		panic("the row of a must equals the col of b")
	}
	col_a := len(a[0])
	c := make([][]int, row_a)
	for i := 0; i < row_a; i++ {
		c[i] = make([]int, col_b)
		for j := 0; j < col_b; j++ {
			c[i][j] = 0
			for k := 0; k < col_a; k++ {
				c[i][j] += a[i][k] * b[k][j]
			}
		}
	}
	return c
}
