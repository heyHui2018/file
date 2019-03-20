package main

import (
	"fmt"
)

/*
思路一：
转换成string,逆序输出
sign := 1
	// 去0
	for {
		if x%10 != 0 {
			break
		}
		x = x / 10
	}
	xStr := strconv.Itoa(x)
	// 去符号
	if x < 0 {
		sign = -1
		xStr = xStr[1:]
	}
	var resultByte []byte
	for i := len(xStr); i > 0; i-- {
		resultByte = append(resultByte, xStr[i-1])
	}
	result, _ := strconv.Atoi(string(resultByte))
	return result * sign
*/

/*
思路二：
通过 %10 和 *10+余数 来代替转string
*/
func reverse(x int) int { // faster 100% less 100%
	sign := 1

	// 处理负数
	if x < 0 {
		sign = -1
		x = -1 * x
	}

	res := 0
	for x > 0 { // 这一步就直接将x倒置
		// 取出x的末尾
		temp := x % 10
		// 放入 res 的开头
		res = res*10 + temp
		// x 去除末尾
		x = x / 10
	}

	// 还原 x 的符号到 res
	res = sign * res

	// 处理 res 的溢出问题
	if res > math.MaxInt32 || res < math.MinInt32 {
		res = 0
	}

	return res
}

func main() {
	x := -12300
	result := reverse(x)
	fmt.Println(result)
}
