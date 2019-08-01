package main

import (
	"fmt"
)

/*
要求:
原地变换,不可申请额外变量

尝试：
第一列变成第一行,第一行变成最后一列,最后一列变成最后一行,最后一行变成第一列

学习：


关键点：


*/

func rotate(matrix [][]int) { // faster 100% less 66.67%
	for i := 0; i < len(matrix)/2; i++ {
		for j := i; j < len(matrix)-i-1; j++ {
			temp := matrix[i][j]
			matrix[i][j] = matrix[len(matrix)-j-1][i]
			matrix[len(matrix)-j-1][i] = matrix[len(matrix)-i-1][len(matrix)-j-1]
			matrix[len(matrix)-i-1][len(matrix)-j-1] = matrix[j][len(matrix)-i-1]
			matrix[j][len(matrix)-i-1] = temp
		}
	}
}

func main() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotate(matrix)
	for k := range matrix {
		fmt.Println(matrix[k])
	}
}
