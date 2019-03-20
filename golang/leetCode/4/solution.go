package main

import (
	"fmt"
)

/*
思路一：
判断总长是奇还是偶来确定中位数(数组长度为奇数时,中间那个;为偶数时,中间两个的平均值),随后遍历nums1、nums2来获取中位数
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {faster 96.34% less 54%
	al := len(nums1)
	bl := len(nums2)
	cl := al + bl
	if cl == 0 {
		return 0
	} else if cl == 1 {
		if al == 0{
			return float64(nums2[0])
		}else{
			return float64(nums1[0])
		}
	}
	targetIndex := 0
	if cl%2 == 0 {
		targetIndex = cl / 2 // 偶数
	} else {
		targetIndex = (cl - 1) / 2 // 奇数
	}
	var c []int
	ai := 0
	bi := 0
	ci := 0
	for ai < al && bi < bl {
		if nums1[ai] < nums2[bi] {
			// c[ci] = nums1[ai]
			if ci == targetIndex-1 || ci == targetIndex {
				c = append(c, nums1[ai])
			}
			ci++
			ai++
		} else {
			if ci == targetIndex-1 || ci == targetIndex {
				c = append(c, nums2[bi])
			}
			// c[ci] = nums2[bi]
			ci++
			bi++
		}
	}
	for ai < al {
		if ci == targetIndex-1 || ci == targetIndex {
			// c[ci] = nums1[ai]
			c = append(c, nums1[ai])
		}
		ci++
		ai++
	}
	for bi < bl {
		if ci == targetIndex-1 || ci == targetIndex {
			// c[ci] = nums2[bi]
			c = append(c, nums2[bi])
		}
		ci++
		bi++
	}
	if cl%2 == 0 {
		return float64(c[0]+c[1]) / 2
	} else {
		return float64(c[1])
	}
}
*/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {//faster 96.34% less 40%
	lenMis, i := len(nums1), 0
	lenNjs, j := len(nums2), 0
	res := make([]int, lenMis+lenNjs)

	for k := 0; k < lenMis+lenNjs; k++ {
		if i == lenMis ||
			(i < lenMis && j < lenNjs && nums1[i] > nums2[j]) {
			res[k] = nums2[j]
			j++
			continue
		}

		if j == lenNjs ||
			(i < lenMis && j < lenNjs && nums1[i] <= nums2[j]) {
			res[k] = nums1[i]
			i++
		}
	}
	l := len(res)
	if l%2 == 0 {
		return float64(res[l/2]+res[l/2-1]) / 2.0
	}

	return float64(res[l/2])
}

func main() {
	var a = []int{}
	var b = []int{1}
	result := findMedianSortedArrays(a, b)
	fmt.Println(result)
}
