package main

import (
	"fmt"
)

/*
尝试：(未完成)
func longestValidParentheses(s string) int {
	temp := 0
	left, index, l, r := 0, 0, 0, 0
	for {
		if index == len(s) {
			return temp
		}
		if s[index] == '(' {
			// 当前为左,则继续加入,index++
			index++
			l++
		} else {
			// 当前为右,1、现有左多则继续加入;2、现有一样多则刷新temp/left,index++
			switch {
			case l > r:
				index++
				r++
			case l == r:
				if l+r > temp {
					temp = l + r
				}
				l, r = 0, 0
				index++
				left = index
			}
		}
		if index == len(s) {
			if r == l && 2*r > temp {
				temp = 2 * r
			}
			if r < l {

			}
		}
	}
}
在遍历结束,右括号数量小于左括号数量时,无法获取最长有效长度

学习：
见注释
*/

func longestValidParentheses(s string) int { // faster 38.38% less 82.73%
	var left, max, temp int
	record := make([]int, len(s))

	// 统计Record,记录每个符号的状态,若能和前面的配上对,则为2,否则为0
	for i, b := range s {
		if b == '(' {
			left++
		} else if left > 0 {
			left--
			record[i] = 2
		}
	}
	fmt.Println(record)

	// 修改record,从左往右检查,若record[i]==2,record[j]==0,且record[j+1:i]中没有0,则record[i]=1,record[j]=1
	for i := 0; i < len(record); i++ {
		if record[i] == 2 {
			j := i - 1
			fmt.Println("j = ", j, " i = ", i, " record[j] = ", record[j])
			for record[j] != 0 {
				j--
			}
			record[i], record[j] = 1, 1
		}
		fmt.Println(record)
	}

	// 统计record中,最多的连续为1的次数
	for _, r := range record {
		if r == 0 {
			temp = 0
			continue
		}
		temp++
		if temp > max {
			max = temp
		}
	}
	return max
}

func main() {
	s := ")(()()))((((())))("
	result := longestValidParentheses(s)
	fmt.Println(result)
}
