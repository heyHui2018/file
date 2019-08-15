package main

import (
	"fmt"
	"strings"
)

/*
要求：


解题思路：
从后遍历,从最后一个不是空格的index开始计数直到再遇到空格

关键点：


*/

func lengthOfLastWord(s string) int { // faster 100% less 33.33%
	list := strings.Split(s, " ")
	for k := range list {
		if len(list[len(list)-1-k]) == 0 {
			continue
		} else {
			return len(list[len(list)-1-k])
		}
	}
	return 0
}

func main() {
	s := "a   "
	result := lengthOfLastWord(s)
	fmt.Println(result)
}
