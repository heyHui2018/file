package main

import (
	"fmt"
)

/*
要求：


解题思路：
51题的基础上取len

关键点：


*/

func totalNQueens(n int) int { // faster 100% less 50%
	res := [][]string{}
	r := []int{}
	l := make([]bool, n)
	d1 := make([]bool, 2*n-1)
	d2 := make([]bool, 2*n-1)
	dfs(&res, &r, &l, &d1, &d2, 0, n)
	return len(res)
}

func dfs(res *[][]string, r *[]int, l, d1, d2 *[]bool, index, n int) {
	if index == n {
		*res = append(*res, generateBoard(n, r))
		return
	}

	for i := 0; i < n; i++ {
		if !(*l)[i] && !(*d1)[index+i] && !(*d2)[index-i+n-1] {
			*r = append(*r, i) // 加入结果集
			(*l)[i] = true
			(*d1)[index+i] = true
			(*d2)[index-i+n-1] = true
			dfs(res, r, l, d1, d2, index+1, n)
			(*l)[i] = false
			(*d1)[index+i] = false
			(*d2)[index-i+n-1] = false
			*r = (*r)[:len(*r)-1] // 移出结果集
		}
	}
	return
}

func generateBoard(n int, row *[]int) []string {
	board := []string{}
	res := ""
	for i := 0; i < n; i++ {
		res += "."
	}
	for i := 0; i < n; i++ {
		board = append(board, res)
	}
	for i := 0; i < n; i++ {
		tmp := []byte(board[i])
		tmp[(*row)[i]] = 'Q'
		board[i] = string(tmp)
	}
	return board
}

func main() {
	n := 4
	result := totalNQueens(n)
	fmt.Println(result)
}
