package main

import (
	"fmt"
)

/*
要求:


尝试：


学习：
累乘,当n>2时,n>>1,随后区分奇偶

关键点：


*/

func myPow(x float64, n int) float64 { // faster 100% less 92.59%
	if n < 0 {
		return 1.0 / pow(x, -n)
	}
	return pow(x, n)
}

func pow(x float64, n int) float64 {
	if x == 0 {
		return 0
	}
	if n == 0 {
		return 1
	}
	res := pow(x, n>>1) // /2
	if n&1 == 0 { // 区分奇偶,n&1 == 0为偶
		return res * res
	}
	return res * res * x
}

func main() {
	// x := 2.10000
	n := 3
	fmt.Println(n >> 1)
	fmt.Println(n & 1)
	// result := myPow(x, n)
	// fmt.Println(result)
}
