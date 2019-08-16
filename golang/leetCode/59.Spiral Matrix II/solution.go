package main

import (
	"fmt"
	"time"
)

/*
要求：


解题思路：
定义4个变量存储上下左右边界,取dx,dy分别表示增量,用x+dx、y+dy与边界的比较来确定拐点

关键点：


*/

func generateMatrix(n int) [][]int { // faster 100% less 100%.
	res := make([][]int, n)
	if n == 0 {
		return res
	}
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}
	top, down := 0, n-1
	left, right := 0, n-1
	x, y := 0, -1 // x-行数 y-列数
	dx, dy := 0, 1
	count := 1
	for {
		if top == down && left == right {
			res[top][left] = count
			return res
		}
		x += dx
		y += dy
		switch {
		case x+dx > down:
			right--
			dx = 0
			dy = -1
		case x+dx < top:
			left++
			dx = 0
			dy = 1
		case y+dy > right:
			top++
			dx = 1
			dy = 0
		case y+dy < left:
			down--
			dx = -1
			dy = 0
		}
		res[x][y] = count
		count++
	}
	return res
}

func main() {
	result := generateMatrix(4)
	for k := range result {
		fmt.Println(result[k])
	}
}
