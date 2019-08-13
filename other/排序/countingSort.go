package main

import (
	"fmt"
)

var array = []int{10, 6, 1, 2, 7, 9, 3, 4, 5, 8}

func main() {
	fmt.Println("排序前:", array)
	result := countingSort(array)
	fmt.Println("排序后:", result)
}

func countingSort(array []int) []int {
	lastArray := make([]int, len(array))
	for i := 0; i < len(array); i++ {
		count := 0
		for j := 0; j < len(array); j++ {
			if array[i] > array[j] {
				count++
			}
		}
		lastArray[count] = array[i]
	}
	return lastArray
}
