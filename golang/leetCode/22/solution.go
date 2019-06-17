package main

import (
	"fmt"
)

/*
思路一：
不应从有多少种组合入手,应从"生成字符串无非在末尾加上左括号或右括号,随后检查此字符串是否满足条件,然后重复上述步骤,直至长度为2n"
(未完成)
*/

/*
思路二：
满足的条件为：1、若左括号已用完,则不能再加左括号;2、若出现的右括号和左括号一样多,则不能再加又括号;3、左右括号均用完则结束.
根据上述条件完成的string一定是正解.
*/

func generateParenthesis(n int) []string { // faster 97.76% less 84.08%
	result := make([]string, 0, n*n)
	// bytes := make([]byte, n*2)
	s := ""
	// gen(n, n, 0, bytes, &result)
	gen1(n, n, s, &result)
	return result
}

// 在生成一个符合条件的string后会return到上一个能出现不同情况的阶段，如第一次生成((()))，会return到((,然后变成(()
func gen(left, right, idx int, bytes []byte, result *[]string) {
	fmt.Println("-------------1111 ", idx)
	fmt.Println("00000000000000000 ", result)
	fmt.Println("11111111111111111 ", string(bytes))
	if left == 0 && right == 0 {
		*result = append(*result, string(bytes))
		return
	}
	if left > 0 {
		bytes[idx] = '('
		gen(left-1, right, idx+1, bytes, result)
	}
	if right > 0 && right > left {
		bytes[idx] = ')'
		gen(left, right-1, idx+1, bytes, result)
	}
}

// todo 为什么return后会变成"((() 1 2"
func gen1(left, right int, s string, result *[]string) {
	fmt.Println("00000000000000000 ", result)
	fmt.Println("11111111111111111 ", s, left, right)
	if left == 0 && right == 0 {
		*result = append(*result, s)
		return
	}
	if left > 0 {
		s += "("
		gen1(left-1, right, s, result)
	}
	if right > 0 && right > left {
		s += ")"
		gen1(left, right-1, s, result)
	}
}

func main() {
	n := 3
	result := generateParenthesis(n)
	fmt.Println(result)
}
