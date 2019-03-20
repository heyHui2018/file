package main

import (
	"fmt"
)

/*
第一次时的思路：
1、负数一定不是
2、通过%拆分放入int[]然后进行比较
*/

func isPalindrome(x int) bool { // faster 100% less 15.93%
	if x < 0 {
		return false
	}
	var list []int
	for {
		if x/10 == 0 {
			list = append(list, x)
			break
		}
		list = append(list, x%10)
		x = x / 10
	}
	for i, j := 0, len(list)-1; i < j; i, j = i+1, j-1 {
		if list[i] != list[j] {
			return false
		}
	}
	return true
}

func main() {
	x := 123321
	result := isPalindrome(x)
	fmt.Println(result)
}
