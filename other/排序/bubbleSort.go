package main

import (
	"fmt"
)

var array = []int{10, 6, 1, 2, 7, 9, 3, 4, 5, 8}

func main() {
	fmt.Println("排序前:", array)
	bubbleSort(array)
	fmt.Println("排序后:", array)
}

func bubbleSort(array []int) {
	for i := 0; i < len(array)-1; i++ {
		isSort := true
		for j := 0; j < len(array)-1-i; j++ {
			temp := 0
			if array[j] < array[j+1] {
				temp = array[j]
				array[j] = array[j+1]
				array[j+1] = temp
				isSort = false
			}
		}
		if (isSort) {
			break
		}
	}
}
