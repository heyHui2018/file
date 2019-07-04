package main

import (
	"fmt"
)

/*
思路一：
先定符号,随后除数循环减被除数,直到除数比被除数小为止,最终检查溢出
func divide(dividend int, divisor int) int { // faster 20.33% less 95.37%
	if divisor == 0 {
		return math.MaxInt32
	}
	sign := 1
	newDividend := dividend
	newDivisor := divisor
	switch {
	case dividend < 0 && divisor < 0:
		newDividend = 0 - dividend
		newDivisor = 0 - divisor
		sign = 1
	case dividend < 0 && divisor > 0:
		newDividend = 0 - dividend
		sign = -1
	case dividend > 0 && divisor < 0:
		newDivisor = 0 - divisor
		sign = -1
	default:
		sign = 1
	}
	count := 0
	for {
		if newDividend < newDivisor {
			break
		}
		newDividend = newDividend - newDivisor
		count++
	}
	if sign < 0 {
		count =  0 - count
	}
	// 检查溢出
	if count < math.MinInt32 || count > math.MaxInt32 {
		return math.MaxInt32
	}
	return count
}
*/

func divide(dividend int, divisor int) int { // faster 80.00% less 69.44%
	// 防止有人把0当做除数
	if divisor == 0 {
		return math.MaxInt32
	}

	signM, absM := analysis(dividend)
	signN, absN := analysis(divisor)

	res, _ := d(absM, absN, 1)

	// 修改res的符号
	if signM != signN {
		res = res - res - res
	}

	// 检查溢出
	if res < math.MinInt32 || res > math.MaxInt32 {
		return math.MaxInt32
	}

	return res
}

func analysis(num int) (sign, abs int) {
	sign = 1
	abs = num
	if num < 0 {
		sign = -1
		abs = num - num - num
	}

	return
}

// d 计算m/n的值，返回结果和余数
// m >= 0
// n > 0
// count == 1, 代表初始n的个数，在递归过程中，count == 当前的n/初始的n
func d(m, n, count int) (res, remainder int) {
	switch {
	case m < n:
		return 0, m
	case n <= m && m < n+n:
		return count, m - n
	default:
		res, remainder = d(m, n+n, count+count)
		if remainder >= n {
			return res + count, remainder - n
		}

		return
	}
}

func main() {
	dividend := 10
	divisor := -1
	result := divide(dividend, divisor)
	fmt.Println(result)
}
