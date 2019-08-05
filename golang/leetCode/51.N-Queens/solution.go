package main

import (
	"fmt"
)

/*
要求：
行、列、对角线仅存在一个queen

解题思路：
取l,d1,d2 3个slice用于存储列及左右对角线是否已存在Q,从0开始循环n,判断l,d1,d2对应位置是否为true(即存在),若不存在,将此点设为Q并加入结果,开始递归(即此点设为Q后后续点在此基础上进行的求解,
当后续发现在此基础上无解时,会return到此处刚将此点设置为Q的时刻,随后取消Q的设置进行下一次迭代),若存在则直接下一次迭代.此处对于n的循环是判断第一行第n个点能否设为Q.

关键点：
d1对角线i+j为定值,d2对角线i-j+n-1为定值

*/

func solveNQueens(n int) [][]string { // faster 80.63% less 100%
	res := [][]string{}
	l := make([]bool, n)
	r := []int{}
	// 边长为n的正方形有2n-1条对角线
	d1 := make([]bool, 2*n-1) // d1上的点,i+j为定值
	d2 := make([]bool, 2*n-1) // d2上的点,i-j为定值,即i-j+n-1
	dfs(&res, &l, &d1, &d2, &r, n, 0)
	return res
}

func dfs(res *[][]string, l, d1, d2 *[]bool, r *[]int, n, index int) {
	if index == n {
		*res = append(*res, generateBoard(n, r))
		return
	}
	for i := 0; i < n; i++ {
		// 尝试将第index行的皇后摆放在第i列
		if !(*l)[i] && !(*d1)[index+i] && !(*d2)[index-i+n-1] {
			*r = append(*r, i) // 此处为append,故r不可再初始化时设置cap
			(*l)[i] = true
			(*d1)[index+i] = true
			(*d2)[index-i+n-1] = true
			dfs(res, l, d1, d2, r, n, index+1)
			(*l)[i] = false
			(*d1)[index+i] = false
			(*d2)[index-i+n-1] = false
			*r = (*r)[:len(*r)-1]
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
	result := solveNQueens(n)
	for k := range result {
		for j := range result[k] {
			fmt.Println(result[k][j])
		}
		fmt.Println("////////")
	}
}
