package main

import (
	"fmt"
)

/*
要求：


解题思路：
定义4个变量存储上下左右边界,取dx,dy分别表示增量,用x+dx、y+dy与边界的比较来确定拐点

关键点：
取dx,dy分别表示增量,用x+dx、y+dy与边界的比较来确定拐点

*/

func spiralOrder(matrix [][]int) []int { // faster 100% less 100%
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}
	res := make([]int, len(matrix)*len(matrix[0]))
	top, down := 0, len(matrix)-1
	left, right := 0, len(matrix[0])-1
	x, y := 0, -1
	dx, dy := 0, 1
	for i := range res {
		x += dx
		y += dy
		switch { // 如果撞墙了，需要修改 dx, dy 和相应的边界值
		case y+dy > right:
			top++
			dx, dy = 1, 0
		case x+dx > down:
			right--
			dx, dy = 0, -1
		case y+dy < left:
			down--
			dx, dy = -1, 0
		case x+dx < top:
			left++
			dx, dy = 0, 1
		}
		res[i] = matrix[x][y]
	}

	return res
}

func main() {
	nums := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	result := spiralOrder(nums)
	fmt.Println(result)
}
