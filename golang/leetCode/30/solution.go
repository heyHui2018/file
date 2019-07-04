package main

import (
	"fmt"
)

/*
思路一：
找出words中所有的排列组合,然后从s中找出出现的index(未完成)
*/

/*
思路二：
见注释
*/
func findSubstring(s string, words []string) []int { // faster 49.10% less 73.21%
	res := make([]int, 0)
	lens := len(s)
	lenws := len(words)
	lenw := len(words[0])
	// 若s长度小于words中子串连起来的长度则结果一定为空
	if lens == 0 || lenws == 0 || lens < lenws*lenw {
		return res
	}
	// map record保存words中各子串的值,value为其出现次数
	record := make(map[string]int, lenws)
	for _, v := range words {
		record[v]++
	}
	// map remain保存剩余待匹配的子串的值,value为其待匹配次数
	remain := make(map[string]int, lenws)
	for k, v := range record {
		remain[k] = v
	}
	// count指目前已匹配的子串个数
	left, right, count := 0, 0, 0
	// 重置,将remain map重置,不可使用remain=record来重置,会导致remain和record指向同一个顶层map,在修改remain时同时也会修改record
	reset := func() {
		for k, v := range record {
			remain[k] = v
		}
		count = 0
	}
	// 移动left,则left原来指向的单词的次数需+1,已匹配的子串数需-1,left指向下一个单词
	moveLeft := func() {
		remain[s[left:left+lenw]]++
		count--
		left += lenw
	}
	// 因为每次移动left或right都是一次移动lenw个字符,故此层循环设定i<lenw是用于指定遍历是从子串的第i个字符开始
	// 加上第二层的for lens-left >= lenws*lenw,即可实现遍历s中所有字符的目的
	for i := 0; i < lenw; i++ {
		left, right = i, i
		reset()
		for lens-left >= lenws*lenw {
			word := s[right : right+lenw]
			remainTimes, ok := remain[word]
			switch {
			case !ok:
				// 不存在这个单词
				// 从目前right所指向单词的下一个单词开始
				left, right = right+lenw, right+lenw
				reset()
			case remainTimes == 0:
				// 单词出现的次数已用完
				// 移动left,直到remain[word]的次数=1停止,此case可能连续出现多次
				moveLeft()
			default:
				// 单词还有剩余匹配次数
				// 匹配次数-1,已匹配个数+1,right指向下一个单词
				remain[word]--
				right += lenw
				count++
				// 判断此时是否已完成一次全匹配
				if count == lenws {
					res = append(res, left)
					// 避免重复统计s[left+len(words[0]):right]中的信息,仅需统计目前right指向单词的下一个单词即可
					moveLeft()
				}
			}
		}
	}
	return res
}

func main() {
	s := "cccab"
	words := []string{"ca", "cc"}
	result := findSubstring(s, words)
	fmt.Println(result)
}
