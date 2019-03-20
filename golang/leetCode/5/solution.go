package main

import (
	"fmt"
)

/*
思路一：(未完成)
利用回文字的中间为xx或xyx的特性逐一遍历,利用map存储字符上一次出现的位置(无法解决"babadada"中最长的是adada的情况)
if len(s) < 2 {
		return s
	}
	sList := strings.Split(s, "")
	maxLength := 0
	result := sList[0:1]
	fmt.Println("result = ", result)
	sMap := make(map[string]int) // map[v]index
	for index, v := range sList {
		flag := true
		if _, ok := sMap[v]; ok {
			tempList := sList[sMap[v] : index+1]
			fmt.Println("tempList = ", tempList)
			l := len(tempList)
			for j, vv := range tempList {
				fmt.Println("tempList[l-1-j] = ", tempList[l-1-j], " vv = ", vv)
				if tempList[l-1-j] != vv {
					flag = false
					break
				}
			}
			if !flag {
				sMap[v] = index
			} else {
				if l > maxLength {
					maxLength = l
					result = tempList
				}
			}
		} else {
			sMap[v] = index
		}
	}
	return strings.Join(result, "")
*/

/*
思路二:
逐一遍历s,若下一位与当前位的值相同,将下一位包括进此次回文,直至下一位的值不同,随后对这一段相同字符组成的串的首位进行逐一比较来确定前后边界
*/

func longestPalindrome(s string) string { // faster 100% less 100%
	if len(s) < 2 { // 肯定是回文，直接返回
		return s
	}
	// 最长回文的首字符索引，和最长回文的长度
	begin, maxLen := 0, 1

	// 在 for 循环中
	// b 代表回文的**首**字符索引号，
	// e 代表回文的**尾**字符索引号，
	// i 代表回文`正中间段`首字符的索引号
	// 在每一次for循环中
	// 先从i开始，利用`正中间段`所有字符相同的特性，让b，e分别指向`正中间段`的首尾
	// 再从`正中间段`向两边扩张，让b，e分别指向此`正中间段`为中心的最长回文的首尾
	for i := 0; i < len(s); { // 以s[i]为`正中间段`首字符开始寻找最长回文。
		if len(s)-i <= maxLen/2 {
			// 因为i是回文`正中间段`首字符的索引号
			// 假设此时能找到的最长回文的长度为l, 则，l <= (len(s)-i)*2 - 1
			// 如果，len(s)-i <= maxLen/2 成立
			// 则，l <= maxLen - 1
			// 则，l < maxLen
			// 所以，无需再找下去了。
			break
		}

		b, e := i, i
		for e < len(s)-1 && s[e+1] == s[e] { // 此循环用来检测一串相同字符串,循环结束后s[b:e+1]是一串相同的字符串
			e++
		}

		// 下一个回文的`正中间段`的首字符只会是s[e+1]
		i = e + 1
		for e < len(s)-1 && b > 0 && s[e+1] == s[b-1] { // s[e+1]为相同字符串的下一位,s[b-1]为相同字符串的上一位,循环结束后s[b:e+1]是这次能找到的最长回文
			e++
			b--
		}

		newLen := e + 1 - b
		// 创新记录的话，就更新记录
		if newLen > maxLen {
			begin = b
			maxLen = newLen
		}
	}
	return s[begin : begin+maxLen]
}

func main() {
	var s = "babaaaaaadada"
	result := longestPalindrome(s)
	fmt.Println(result)
}
