package main

import (
	"fmt"
	"math"
)

var array = []int{10, 1, 18, 30, 23, 12, 7, 5, 18, 233, 144}

func main() {
	fmt.Println("排序前:", array)
	result := radixSort(array)
	fmt.Println("排序后:", result)
}

func radixSort(theArray []int) []int {
	// 获取最大值vl
	vl := 0;
	for _, v := range theArray {
		if v > vl {
			vl = v
		}
	}
	// 获取最大值的位数
	var count int = 0;
	for vl%10 > 0 {
		vl = vl / 10;
		count++;
	}

	// 给桶中对应的位置放数据
	for i := 0; i < count; i++ {
		fmt.Println(theArray)
		theData := int(math.Pow10(i)); // 10的i次方
		// 建立空桶
		var bucket [10][10]int;
		for k := 0; k < len(theArray); k++ {
			theResidue := (theArray[k] / theData) % 10; // 取余
			var childArray [10]int;                     // = bucket[theResidue];//获取子数组
			for m := 0; m < 10; m++ {
				if (bucket[theResidue][m] == 0) {
					childArray[m] = theArray[k];
					bucket[theResidue][m] = childArray[m];
					break;
				} else {
					continue;
				}
			}
		}
		// 一遍循环完之后需要把数组二维数据进行重新排序，比如数组开始是10 1 18 30 23 12 7 5 18 233 144 ，循环个位数
		// 循环之后的结果为10 30 1 12 23 233 144 5 7 18 18 ，然后循环十位数，结果为1 5 7 10 12 18 18 23 30 233 144
		// 最后循环百位数，结果为1 5 7 10 12 18 18 23 30 144 233
		var x = 0;
		slice := make([]int, len(theArray))
		for p := 0; p < len(bucket); p++ {
			for q := 0; q < len(bucket[p]); q++ {
				if (bucket[p][q] != 0) {
					slice[x] = bucket[p][q];
					x++;
				} else {
					break;
				}
			}
		}

		for key, value := range slice {
			theArray[key] = value
		}
	}
	return theArray
}
