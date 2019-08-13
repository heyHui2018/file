package main

import (
	"fmt"
	"sort"
)

var array = []int{10, 1, 18, 30, 23, 12, 7, 5, 18, 17}

func main() {
	fmt.Println("排序前:", array)
	bucketSort(array)
	fmt.Println("排序后:", array)
}

func bucketSort(arr []int) {
	// 最大最小值
	max := arr[0]
	min := arr[0]
	length := len(arr)

	for i := 1; i < length; i++ {
		if arr[i] > max {
			max = arr[i]
		} else if arr[i] < min {
			min = arr[i]
		}
	}

	// 最大值和最小值的差
	diff := max - min

	// 桶列表
	bucketList := [][]int{}
	for i := 0; i < length; i++ {
		bucketList = append(bucketList, []int{})
	}

	// 每个桶的存数区间
	section := float64(diff) / float64(length-1)

	// 数据入桶
	for i := 0; i < length; i++ {
		// 当前数除以区间得出存放桶的位置 减1后得出桶的下标
		num := int(float64(arr[i])/section) - 1
		if (num < 0) {
			num = 0
		}
		bucketList[num] = append(bucketList[num], arr[i])
	}

	// 桶内排序
	for i := 0; i < len(bucketList); i++ {
		// jdk的排序速度当然信得过
		sort.Ints(bucketList[i]);
	}

	// 写入原数组
	index := 0;
	for _, v := range bucketList {
		for _, w := range v {
			arr[index] = w
			index++
		}
	}
}
