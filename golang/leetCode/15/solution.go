package main

import (
	"fmt"
	"sort"
)

/*
思路一：
三重循环遍历-超时
func threeSum(nums []int) [][]int {
	result := [][]int{}
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
		c:
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					temp := []int{nums[i], nums[j], nums[k]}
					sort.Ints(temp)
					for p := range result {
						sort.Ints(result[p])
						if reflect.DeepEqual(result[p], temp) {
							continue c
						}
					}
					result = append(result, temp)
				}
			}
		}
	}
	return result
}
*/

/*
先排序，遍历，此次循环数字与上次相同时continue
*/

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}

	for i := range nums {
		// 避免添加重复的结果
		// i>0 是为了防止nums[i-1]溢出
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		l, r := i+1, len(nums)-1

		for l < r {
			s := nums[i] + nums[l] + nums[r]
			switch {
			case s < 0:
				// 较小的 l 需要变大
				l++
			case s > 0:
				// 较大的 r 需要变小
				r--
			default:
				res = append(res, []int{nums[i], nums[l], nums[r]})
				// 为避免重复添加，l 和 r 都需要移动到不同的元素上。
				l, r = next(nums, l, r)
			}
		}
	}
	return res
}

func next(nums []int, l, r int) (int, int) {
	for l < r {
		switch {
		case nums[l] == nums[l+1]:
			l++
		case nums[r] == nums[r-1]:
			r--
		default:
			l++
			r--
			return l, r
		}
	}
	return l, r
}

func main() {
	nums := []int{7, -10, 7, 3, 14, 3, -2, -15, 7, -1, -7, 6, -5, -1, 3, -13, 6, -15, -10, 14, 8, 5, -10, -1, 1, 1, 11, 6, 8, 5, -4, 0, 3, 10, -12, -6, -2, -6, -6, -10, 8, -5, 12, 10, 1, -8, 4, -8, -8, 2, -9, -15, 14, -11, -1, -8, 5, -13, 14, -2, 0, -13, 14, -12, 12, -13, -3, -13, -12, -2, -15, 4, 8, 4, -1, -6, 11, 11, -7, -12, -2, -8, 10, -3, -4, -6, 4, -14, -12, -5, 0, 3, -3, -9, -2, -6, -15, 2, -11, -11, 8, -11, 8, -7, 8, 14, -5, 4, 10, 3, -1, -15, 10, -6, -11, 13, -5, 1, -15}
	result := threeSum(nums)
	fmt.Println(result)
}
