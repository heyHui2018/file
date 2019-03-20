package main

import (
	"fmt"
)

/*
思路一：
表驱动法及从右往左加,当左边的小于相邻右边的数时,减去这个数.
*/

func romanToInt(s string) int { // faster 100% less 100%
	d := [4][]string{
		[]string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"},
		[]string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"},
		[]string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"},
		[]string{"", "M", "MM", "MMM"},
	}
	return d[3][num/1000] +
		d[2][num/100%10] +
		d[1][num/10%10] +
		d[0][num%10]
}

func main() {
	num := "MCMXCII"
	result := romanToInt(num)
	fmt.Println(result)
}
