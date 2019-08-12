package main

import (
	"fmt"
)

var array = []int{10, 6, 1, 2, 7, 9, 3, 4, 5, 8}

func main() {
	fmt.Println("排序前:", array)
	selectionSort(array)
	fmt.Println("排序后:", array)
}

func selectionSort(array []int) {
	for i := 0; i < len(array); i++ {
		min := i; // 最小元素的下标
		for j := i + 1; j < len(array); j++ {
			if array[j] < array[min] {
				min = j; // 找最小值
			}
		}
		// 交换位置
		temp := array[i];
		array[i] = array[min];
		array[min] = temp;
	}
}
