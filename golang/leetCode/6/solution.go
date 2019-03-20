package main

import (
	"fmt"
	"bytes"
)

/*
思路一：(未完成)
找规律
*/

func convert(s string, numRows int) string {// faster 98.69% less 90.91%
	/*
		numRows = 5
		0行，0, 8, 16, 24
		1行，1, 7,9, 15,17, 23,25
		2行，2, 6, 10, 14, 18, 22
		3行，3,5, 11,13, 19,21
		4行，4, 12, 20

		令p=numRows×2-2，则：
		0行， 0×p，1×p，...
		r行， r，1×p-r，1×p+r，2×p-r，2×p+r，...
		最后一行， numRow-1, numRow-1+1×p，numRow-1+2×p，...
	*/
	if numRows == 1 || len(s) <= numRows {
		return s
	}

	res := bytes.Buffer{}
	// p pace 步距
	p := numRows*2 - 2

	// 处理第一行
	for i := 0; i < len(s); i += p {
		res.WriteByte(s[i])
	}

	// 处理中间行
	for r := 1; r <= numRows-2; r++ {
		// 添加r行的第一个字符
		res.WriteByte(s[r])

		for k := p; k-r < len(s); k += p {
			res.WriteByte(s[k-r])
			if k+r < len(s) {
				res.WriteByte(s[k+r])
			}
		}
	}

	// 处理最后一行
	for i := numRows - 1; i < len(s); i += p {
		res.WriteByte(s[i])
	}

	return res.String()
}

func main() {
	s := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numRows := 5
	result := convert(s, numRows)
	fmt.Println(result)
}
