package main

import (
	"fmt"
)

/*
思路一：
依次与1000、500、100...进行比较,然后通过取整取余操作完成转换.
func intToRoman(num int) string { // faster 74.9% less 11.11%
	var s []string
	for {
		switch {
		case num >= 1000:
			for i := 0; i < num/1000; i++ {
				s = append(s, "M")
			}
			num = num % 1000
		case num >= 900:
			s = append(s, "CM")
			num = num - 900
		case num >= 500:
			s = append(s, "D")
			num = num - 500
		case num >= 400:
			s = append(s, "CD")
			num = num - 400
		case num >= 100:
			for i := 0; i < num/100; i++ {
				s = append(s, "C")
			}
			num = num % 100
		case num >= 90:
			s = append(s, "XC")
			num = num - 90
		case num >= 50:
			s = append(s, "L")
			num = num - 50
		case num >= 40:
			s = append(s, "XL")
			num = num - 40
		case num >= 10:
			for i := 0; i < num/10; i++ {
				s = append(s, "X")
			}
			num = num % 10
		case num == 9:
			s = append(s, "IX")
			num = num - 9
		case num >= 5:
			s = append(s, "V")
			num = num - 5
		case num == 4:
			s = append(s, "IV")
			num = num - 4
		case num >= 1:
			for i := 0; i < num; i++ {
				s = append(s, "I")
			}
			num = 0
		}
		if num == 0 {
			break
		}
	}
	return strings.Join(s, "")
}
*/
/*
思路二：
表驱动法
*/

func intToRoman(num int) string { // faster 100% less 100%
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
	num := 99
	result := intToRoman(num)
	fmt.Println(result)
}
