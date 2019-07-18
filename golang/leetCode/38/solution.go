package main

import (
	"fmt"
)

/*
要求:
1:1
2:1个1->11
3:2个1->21
4:1个2、1个1->1211
5:1个1、1个2、2个1->111221
6:3个1、2个2、1个1->312211


尝试：
后一个数在前一个数基础上处理,通过for循环实现,

学习：


关键点：

*/

func countAndSay(n int) string { // faster 100.00% less 88.18%
	result := ""
	for i := 1; i <= n; i++ {
		result = count(i, result)
	}
	return result
}

func count(n int, basic string) string {
	switch n {
	case 0:
		return ""
	case 1:
		return "1"
	default:
		tempCount := 0
		tempNum := 0
		list := make([]byte, 0)
		for k, v := range basic {
			num := int(v - '0')
			if tempNum == 0 {
				tempNum = num
				tempCount++
			} else {
				if num == tempNum {
					tempCount++
				} else {
					list = append(list, byte(tempCount)+'0')
					list = append(list, byte(tempNum)+'0')
					tempCount = 1
					tempNum = num
				}
			}
			if k == len(basic)-1 {
				list = append(list, byte(tempCount)+'0')
				list = append(list, byte(tempNum)+'0')
			}
		}
		return string(list)
	}
}

func main() {
	result := countAndSay(5)
	fmt.Println(result)
}
