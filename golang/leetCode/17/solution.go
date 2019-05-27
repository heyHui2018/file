package main

import (
	"fmt"
)

/*
思路一：
表驱动法
dMap := map[string][]string{
		"2": []string{"a", "b", "c"},
		"3": []string{"d", "e", "f"},
		"4": []string{"g", "h", "i"},
		"5": []string{"j", "k", "l"},
		"6": []string{"m", "n", "o"},
		"7": []string{"p", "q", "r", "s"},
		"8": []string{"t", "u", "v"},
		"9": []string{"w", "x", "y", "z"},
	}
未完成-无法预知循环层数
*/

func letterCombinations(digits string) []string { // faster 100% less 5.06%
	m := map[byte][]string{
		'2': []string{"a", "b", "c"},
		'3': []string{"d", "e", "f"},
		'4': []string{"g", "h", "i"},
		'5': []string{"j", "k", "l"},
		'6': []string{"m", "n", "o"},
		'7': []string{"p", "q", "r", "s"},
		'8': []string{"t", "u", "v"},
		'9': []string{"w", "x", "y", "z"},
	}
	if digits == "" {
		return nil
	}
	ret := []string{""}
	// 让digits中所有的数字都有机会进行迭代。
	for i := 0; i < len(digits); i++ {
		temp := []string{}
		// 让ret中的每个元素都能添加新的字母。
		for j := 0; j < len(ret); j++ {
			// 把digits[i]所对应的字母，多次添加到ret[j]的末尾
			for k := 0; k < len(m[digits[i]]); k++ {
				temp = append(temp, ret[j]+m[digits[i]][k])
			}
			fmt.Println(temp)
		}
		ret = temp
	}
	return ret
}

func main() {
	digits := "233"
	result := letterCombinations(digits)
	fmt.Println(result)
}
