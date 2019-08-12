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
	gap := 1;
	for gap < len(array) {
		gap = gap*3 + 1;
	}
	for gap > 0 {
		for i := gap; i < len(array); i++ {
			tmp := array[i];
			j := i - gap;
			// 跨区间排序
			for j >= 0 && array[j] > tmp {
				array[j+gap] = array[j];
				j -= gap;
			}
			array[j+gap] = tmp;
		}
		gap = gap / 3;
	}
}
