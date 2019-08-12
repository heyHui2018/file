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
	for i := 1; i < len(array); i++ {
		for j := i; j > 0 && array[j] < array[j-1]; j-- {
			array[j], array[j-1] = array[j-1], array[j]
		}
	}
}
