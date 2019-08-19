package main

import (
	"fmt"
)

/*
要求：


解题思路：
第一个数会出现(n-1)!次,即res[0]=k/(n-1)!,k=k%(n-1)!;第二个数会出现(n-2)!,以此类推...根据此规律逐一生成目标组合中的数即可,需注意每个数仅出现一次

关键点：
第一个数会出现(n-1)!次,即res[0]=k/(n-1)!,k=k%(n-1)!;第二个数会出现(n-2)!,以此类推

*/

func getPermutation(n int, k int) string { // faster 100% less 100%.
	res := make([]byte, n)
	if n == 0 {
		return ""
	}
	// 计算阶乘
	p := 1
	for i := 2; i < n; i++ {
		p *= i
	}
	// 因从0开始,需--
	k--
	// 剩余的待排数,当k被选出后,需从待排数中移除
	rec := make([]byte, n)
	for i := 0; i < n; i++ {
		rec[i] = byte(i) + '1'
	}
	// 根据规律逐一生成
	for i := 0; i < n-1; i++ {
		idx := k / p
		res[i] = rec[idx]
		rec = append(rec[:idx], rec[idx+1:]...)
		k %= p
		p /= (n - i - 1)
	}
	// 加上最后一个数
	res[n-1] = rec[0]
	return string(res)
}

func main() {
	n := 3
	k := 6
	result := getPermutation(n, k)
	fmt.Println(result)
}
