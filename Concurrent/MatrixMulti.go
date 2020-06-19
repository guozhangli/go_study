package Concurrent

import (
	"math/rand"
	"time"
)

/**
矩阵乘法
*/
func generate(row, col int) [][]int {
	matrix := make([][]int, row)
	for i := 0; i < row; i++ {
		v := make([]int, col)
		matrix[i] = v
		for j := 0; j < col; j++ {
			rand.Seed(time.Now().Unix())
			num := rand.Intn(10)
			matrix[i][j] = num
		}
	}
	return matrix
}

func serial_multiply(a, b [][]int) [][]int {
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

func parallel_row_multiply(a, b [][]int) [][]int {
	row_a := len(a)
	col_b := len(b[0])
	if row_a != col_b {
		panic("the row of a must equals the col of b")
	}
	col_a := len(a[0])
	c := make([][]int, row_a)
	ch := make(chan int, 10)
	for i := 0; i < row_a; i++ {
		go parallel_row(a, b, c, i, col_a, col_b, ch)
		go read_chan(i, ch)
	}
	return c
}

func parallel_row(a, b, c [][]int, i, col_a, col_b int, ch chan int) {
	c[i] = make([]int, col_b)
	for j := 0; j < col_b; j++ {
		c[i][j] = 0
		for k := 0; k < col_a; k++ {
			c[i][j] += a[i][k] * b[k][j]
		}
	}
	ch <- 1
}

func read_chan(i int, ch chan int) {
	if i%10 == 0 {
		for c := 0; c < 10; c++ {
			<-ch
		}
	}
}

func parallel_individual_multiply(a, b [][]int) [][]int {
	row_a := len(a)
	col_b := len(b[0])
	if row_a != col_b {
		panic("the row of a must equals the col of b")
	}
	col_a := len(a[0])
	c := make([][]int, row_a)
	ch := make(chan int, 10)
	for i := 0; i < row_a; i++ {
		c[i] = make([]int, col_b)
		for j := 0; j < col_b; j++ {
			go parallel_individual(a, b, c, i, j, col_a, ch)
			go read_chan(j, ch)
		}
	}
	return c
}

func parallel_individual(a, b, c [][]int, i, j, col_a int, ch chan int) {
	c[i][j] = 0
	for k := 0; k < col_a; k++ {
		c[i][j] += a[i][k] * b[k][j]
	}
	ch <- 1
}

func parallel_group_multiply(a, b [][]int) [][]int {
	row_a := len(a)
	col_b := len(b[0])
	if row_a != col_b {
		panic("the row of a must equals the col of b")
	}
	col_a := len(a[0])
	c := make([][]int, row_a)
	ch := make(chan int, 10)
	var startIndex, endIndex int
	var step = row_a / 10
	endIndex = step
	for i := 0; i < 10; i++ {
		go parallel_group(a, b, c, col_a, col_b, startIndex, endIndex, ch)
		startIndex = endIndex
		if i < 8 {
			endIndex = endIndex + step
		} else {
			endIndex = row_a
		}
	}
	go read_chan(10, ch)
	return c
}

func parallel_group(a, b, c [][]int, col_a, col_b, start_index, end_index int, ch chan int) {
	for i := start_index; i < end_index; i++ {
		c[i] = make([]int, col_b)
		for j := 0; j < col_b; j++ {
			c[i][j] = 0
			for k := 0; k < col_a; k++ {
				c[i][j] += a[i][k] * b[k][j]
			}
		}
	}
	ch <- 1
}
