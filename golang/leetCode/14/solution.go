package main

import (
	"fmt"
)

/*
思路一：
取strs中某个字符并从头开始和其余字符串进行比较,都成功则加一位继续比较,否则返回.
*/

func longestCommonPrefix(strs []string) string { // faster 100% less 56.24%
	switch (len(strs)) {
	case 0:
		return ""
	case 1:
		return strs[0]
	default:
		count := 1
	a:
		for i := 1; i <= len(strs[0]); i++ { // 遍历第一个单词
			prefix := strs[0][:i]
			for j := 1; j < len(strs); j++ { // 遍历数组
				if len(strs[j]) < i { // strs[j]长度小于strs[0]时，返回strs[0][:i-1]
					return strs[0][:i-1]
				}
				if prefix != strs[j][:i] {
					break a
				}
			}
			count++
		}
		if len(strs[0]) < count {
			return strs[0][:count-1]
		}
		return strs[0][:count-1]
	}
}

func main() {
	strs := []string{"a", "a"}
	result := longestCommonPrefix(strs)
	fmt.Println(result)
}
