package main

import (
	"fmt"
)

/*
要求:


尝试：
切分每个单词进而生成对应字母表的string,作为map的key,最后遍历map输出

学习：


关键点：


*/

func groupAnagrams(strs []string) [][]string { // faster 15.45% less 95.83%
	keyMap := make(map[string][]string)
	for k := range strs {
		if _, ok := keyMap[sort(strs[k])]; !ok {
			var slice []string
			slice = append(slice, strs[k])
			keyMap[sort(strs[k])] = slice
		} else {
			keyMap[sort(strs[k])] = append(keyMap[sort(strs[k])], strs[k])
		}
	}
	var result [][]string
	for j := range keyMap {
		result = append(result, keyMap[j])
	}
	return result
}

func sort(str string) string {
	slice := make([]int, 26)
	for k := range str {
		slice[str[k]-'a'] += 1
	}
	var key string
	for j := range slice {
		key += string(slice[j])
	}
	return key
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	result := groupAnagrams(strs)
	fmt.Println(result)
}
