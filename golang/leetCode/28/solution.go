package main

import (
	"fmt"
)

/*
思路一：
遍历haystack,当当前字母与needle中的第一个字母匹配时,比较后续字母,完全相等时返回当前遍历进度的index,否则继续遍历,匹配失败则返回-1
func strStr1(haystack string, needle string) int { // faster 100% less 65.52%
	if len(needle) == 0 {
		return 0
	} else {
		if len(haystack) == 0 {
			return -1
		}
	}
	left, index := 0, 0
	for {
		if left == len(haystack) || left+index == len(haystack) {
			return -1
		}
		if haystack[left+index] == needle[index] {
			if index == len(needle)-1 {
				return left
			}
			index++
		} else {
			index = 0
			left++
		}
	}
}
*/

/*
思路二：
比思路一更简便,当第一个字母匹配时,直接截取haystack[index:index+len(needle)]与needle比较,无需逐个字母比较
*/

func strStr(haystack string, needle string) int { // faster 100% less 65.52%
	hlen, nlen := len(haystack), len(needle)
	// 当hlen等于nlen的时候，需要i == 0
	for i := 0; i <= hlen-nlen; i++ {
		if haystack[i:i+nlen] == needle {
			return i
		}
	}

	return -1
}
func main() {
	haystack := "mississippi"
	needle := "issipi"
	result := strStr(haystack, needle)
	fmt.Println(result)
}
