package main

import (
	"fmt"
)

/*
思路一：
表驱动法及从右往左加,当左边的小于相邻右边的数时,减去这个数.
*/

func romanToInt(s string) int { // faster 100% less 70%
	romanMap := make(map[string]int, 8)
	romanMap["I"] = 1
	romanMap["V"] = 5
	romanMap["X"] = 10
	romanMap["L"] = 50
	romanMap["C"] = 100
	romanMap["D"] = 500
	romanMap["M"] = 1000

	total := 0
	for i := len(s) - 1; i >= 0; i-- {
		if i == len(s)-1 {
			total = romanMap[string(s[i])]
			continue
		}
		if romanMap[string(s[i])] < romanMap[string(s[i+1])] {
			total = total - romanMap[string(s[i])]
		} else {
			total = total + romanMap[string(s[i])]
		}
	}
	return total
}

func main() {
	num := "MCMXCII"
	result := romanToInt(num)
	fmt.Println(result)
}
