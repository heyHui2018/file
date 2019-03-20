package main

import (
	"fmt"
	"math"
	"strings"
)

/*
思路一：(未完成)
无思路
*/

/*
思路二：
通过 abs = abs*10 + int(b-'0') 将str转成int
*/
func myAtoi(str string) int { // faster 100% less 100%
	// 去除前后空格
	str = strings.TrimSpace(str)
	if str == "" {
		return 0
	}
	// 保存符号数字str
	var sign int
	var abs string
	switch str[0] {
	case '+':
		sign = 1
		abs = str[1:]
	case '-':
		sign = -1
		abs = str[1:]
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		sign = 1
		abs = str
	default:
		return 0
	}
	// 获取第一个非数字前的数字段
	for i, v := range abs {
		if v < '0' || v > '9' {
			abs = abs[:i]
			break
		}
	}
	var result int
	// 通过 abs = abs*10 + int(b-'0') 将str转成int
	for _, v := range abs {
		result = result*10 + int(v-'0')
		if sign == 1 && result >= math.MaxInt32 {
			return math.MaxInt32
		} else if sign == -1 && sign*result <= math.MinInt32 {
			return math.MinInt32
		}
	}
	return sign * result
}

func main() {
	str := "-12300"
	result := myAtoi(str)
	fmt.Println(result)
}
