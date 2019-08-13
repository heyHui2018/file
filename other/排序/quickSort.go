package main

import (
	"fmt"
)

// 双边扫描

var array = []int{10, 6, 1, 2, 7, 9, 3, 4, 5, 8}

func main() {
	fmt.Println("排序前:", array)
	quickSort(0, len(array)-1)
	fmt.Println("排序后:", array)
}

func quickSort(left, right int) {
	var i, j, t, temp int
	if left > right {
		return
	}
	temp = array[left]
	i = left
	j = right
	for {
		if i == j {
			break
		}
		for {
			if array[j] < temp || i >= j {
				break
			}
			j--
		}
		for {
			if array[i] > temp || i >= j {
				break
			}
			i++
		}
		if i < j {
			t = array[i]
			array[i] = array[j]
			array[j] = t
		}
	}
	array[left] = array[i]
	array[i] = temp
	fmt.Println("排序中:", array)
	quickSort(left, i-1)
	quickSort(i+1, right)
}
