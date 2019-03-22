package main

import (
	"fmt"
)

/*
思路一：
取strs中某个字符并从头开始和其余字符串进行比较,都成功则加一位继续比较,否则返回.
*/

func longestCommonPrefix(strs []string) string { // faster 100% less 70%

}

func main() {
	strs := []string{"1", "2"}
	result := longestCommonPrefix(strs)
	fmt.Println(result)
}
