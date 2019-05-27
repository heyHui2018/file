package main

import (
	"fmt"
)

/*
思路一：(未完成,需加计算i1-1-j1的逻辑)
遍历s并放入map,当发现重复时(此时index=i1,上次出现为i0,上次重复时后一次的index为j1),计算i1-i0,计算i1-1-j1,与现有的长度作比较,较长的放入result,最终输出result
if len(s) == 0 {
		return 0
	}
	sMap := make(map[string]int, len(s))
	sList := strings.Split(s, "")
	result := 1
	temp := 0
	for k, v := range sList {
		if kOld, ok := sMap[v]; ok {
			if kOld+1-temp > result {
				result = kOld + 1 - temp
			}
			if kOld+1 > temp {
				temp = kOld + 1
			}
			fmt.Println("111111111temp = ", temp, " result = ", result)
			if k+1-temp > result {
				result = k + 1 - temp
			}
			fmt.Println("2222222222result = ", result)
		}
		sMap[sList[k]] = k
	}
	if len(s)-temp > result {
		result = len(s) - temp
	}
	return result
*/

func lengthOfLongestSubstring(s string) int {
	// location[s[i]] == j 表示：s中第i个字符串,上次出现在s的j位置,所以在s[j+1:i]中没有s[i]
	// location[s[i]] == -1 表示：s[i] 在s中第一次出现
	location := [256]int{} // 只有256长是因为假定输入的字符串只有ASCII字符
	for i := range location {
		location[i] = -1 // 先设置所有的字符都没有见过
	}

	maxLen, left := 0, 0

	for i := 0; i < len(s); i++ {
		// 说明s[i]已经在s[left:i+1]中重复了并且s[i]上次出现的位置在location[s[i]]
		if location[s[i]] >= left {
			left = location[s[i]] + 1 // 在s[left:i+1]中去除s[i]字符及其之前的部分
		} else if i+1-left > maxLen {
			maxLen = i + 1 - left
		}
		location[s[i]] = i
	}
	return maxLen
}

func main() {
	s := "abcb"
	result := lengthOfLongestSubstringH(s)
	fmt.Println(result)
}
