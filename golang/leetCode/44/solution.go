package main

import (
	"fmt"
)

/*
要求:


尝试：
(未完成)

学习：
记录上一个*的位置及*对应匹配s中字符的位置,当出现不匹配又不是?和*时,若出现过*,则回滚,使*继续匹配这次的字符,若没出现过*则返回false,当s已完成匹配而p未完成时,只有p中剩余的均为*时才返回true

关键点：
当出现不匹配又不是?和*时,若出现过*,则回滚,使*继续匹配这次的字符

*/

func isMatch(s string, p string) bool { // faster 100% less 85%
	// matchIndex:待匹配字符下标
	// lastStarIndex:上一个*下标
	sIndex, pIndex, matchIndex, lastStarIndex := 0, 0, 0, -1
	// 用p匹配s,故遍历s
	for sIndex < len(s) {
		if pIndex < len(p) && (s[sIndex] == p[pIndex] || p[pIndex] == '?') {
			// 因遍历s,故需确保pIndex < len(p),当字符匹配或p中出现?时,两index均往后移动一位
			sIndex++
			pIndex++
		} else if pIndex < len(p) && p[pIndex] == '*' {
			// 因遍历s,故需确保pIndex < len(p),当p中出现*时,记录*的位置,记录此时s中开始被*匹配的位置,p的index往后移动一位
			lastStarIndex = pIndex
			matchIndex = sIndex
			pIndex++
		} else if lastStarIndex != -1 {
			// 既不匹配,p中对应的也不是*或?,但之前出现过*,则开始回滚,p的index回到*的后一位,待匹配字符下标继续往后,s的index同待匹配字符下标
			pIndex = lastStarIndex + 1
			matchIndex++
			sIndex = matchIndex
		} else {
			return false
		}
	}
	// 当s已匹配完,但p还未完时,只有p中剩余的均为*才为true
	for i := pIndex; i < len(p); i++ {
		if p[i] != '*' {
			return false
		}
	}
	return true
}

func main() {
	s := "aa"
	p := "a"
	result := isMatch(s, p)
	fmt.Println(result)
}
