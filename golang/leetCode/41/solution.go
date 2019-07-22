package main

import (
	"fmt"
)

/*
要求:
时间复杂度O(n)、空间复杂度固定

尝试：
(未完成)

学习：
因是最小正整数,故其取值定在[1,len(nums)+1]中,故可用nums[k]存放k+1,当某个k存放的不是对应的k+1时,结果便是此k+1
nums[i] != nums[nums[i]-1] 即 k+1 != nums[k],则将其交换位置,通过for循环交换,最终会将1/2/3...放到其位置上(若nums中存在的话),若nums中不存在,则其位置为其余数,在最后遍历nums即可发现

关键点：
发现取值空间在[1,len(nums)+1]中,可通过nums[k]存放k+1来计数,明白nums[i] != nums[nums[i]-1] 即为 k+1 != nums[k]

*/

func firstMissingPositive(nums []int) int { // faster 100% less 98.53%
	for i := 0; i < len(nums); i++ {
		for 1 <= nums[i] && nums[i] < len(nums)+1 && nums[i] != nums[nums[i]-1] {
			fmt.Println("0000000 i = ", i, " nums[i]-1 = ", nums[i]-1, " nums[i] = ", nums[i], " nums[nums[i]-1] = ", nums[nums[i]-1])
			// 当 for 的判断语句成立时，
			// nums[i]-1 就是 k ，nums[i] 的值是 k+1
			// nums[i] != nums[nums[i]-1] 即是 k+1 != nums[k] ，这说明
			// 1. k+1 存在与 nums 中，
			// 2. k+1 还没有在他该在的 nums[k] 中
			// 通过互换，让 k+1 到 nums[k] 中去
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]

			// fmt.Println(nums)

			// 使用 for 而不是 if 是因为
			// nums[i] 中的新值，有可能是另一个 k+1 ，需要再让其归为
			// 如果使用 if ，而这个新的 k+1 又只有一个的话，
			// 这个新的 k+1 不会被处理到，不会被放在 nums[k] 中
			fmt.Println("11111111 nums = ", nums)
		}
		fmt.Println("22222222 nums = ", nums)
	}
	// 循环结束后，所有 1<=k+1<=len(nums) 且 k+1 存在于nums中，都会被存放于 nums[k] 中

	// 整理后，第一个不存在的 k+1 就是答案
	for k := range nums {
		if nums[k] != k+1 {
			return k + 1
		}
	}
	return len(nums) + 1
}

func main() {
	nums := []int{7, 9, 10, 11, 8, 9, 13}
	result := firstMissingPositive(nums)
	fmt.Println(result)
}
