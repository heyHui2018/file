package main

import (
	"fmt"
)

/*
要求:


尝试：
递归,当前组合的结果基于其子集的结果,如当前组合为{1,2},则有两种组合1-2和2-1,则{1,2,3}有3-1-2,1-3-2,1-2-3,3-2-1,2-3-1,2-1-3 6中组合,在{1,2}的2种组合的基础上产生的,在迭代过程中通过map去重

学习：


关键点：
[1,2,1]加2出现的[2,1,2,1]和[2,1,1]加2出现的[2,1,2,1]会重复,通过map解决此问题,key为遍历slice组成的string

*/

func permuteUnique(nums []int) [][]int { // faster 38.10% less 50.00%
	result := make([][]int, 0)
	if len(nums) < 1 {
		return result
	}
	result = append(result, nums[:1])
	resultMap := make(map[string]int)
	for i := 1; i < len(nums); i++ {
		result = p(result, nums[i], resultMap)
	}
	return result
}

func p(base [][]int, num int, resultMap map[string]int) [][]int {
	result := make([][]int, 0)
	for i := range base {
		for j := range base[i] {
			// 此循环中均为num放此base[i]之前,故当num == base[i][j]时continue,可去重
			if num == base[i][j] {
				continue
			}
			// 此处不可使用temp:=base[i][:j],因在temp = append(temp, num)后,会使base[i]也改变
			var temp []int
			temp = append(temp, base[i][:j]...)
			temp = append(temp, num)
			temp = append(temp, base[i][j:]...)
			key, ok := check(temp, resultMap)
			if ok {
				continue
			}
			result = append(result, temp)
			resultMap[key] = 1
		}
		// 加上num在base[i]末尾的情况
		var temp []int
		temp = append(base[i], num)
		key, ok := check(temp, resultMap)
		if ok {
			continue
		}
		result = append(result, temp)
		resultMap[key] = 1
	}
	return result
}

func check(slice []int, resultMap map[string]int) (string, bool) {
	key := ""
	for k := range slice {
		key += string(slice[k])
	}
	_, ok := resultMap[key];
	return key, ok
}

func main() {
	nums := []int{1, 1, 2, 2}
	result := permuteUnique(nums)
	for _, v := range result {
		fmt.Println(v)
	}
}
