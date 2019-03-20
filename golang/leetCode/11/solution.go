package main

import (
	"fmt"
)

/*
思路一：
输入数组中各元素的index即为其横坐标,值为纵坐标,寻找两个index形成面积最大的矩形.
穷举法-时间复杂度O(n^2)
func maxArea(height []int) int { // faster 9.97% less 12.50%
	var area, temp, h int
	for i, v := range height {
		for ii, vv := range height {
			h = v
			if v > vv {
				h = vv
			}
			temp = (ii - i) * h
			if temp < 0 {
				temp = 0 - temp
			}
			if temp > area {
				area = temp
			}
		}
	}
	return area
}
*/

/*
思路二：
穷举法是从头单向遍历到尾,故同时从头(ai)尾(aj)向中间遍历可耗时更少.当ai<aj时,i++,当ai>aj时,j--.因为只有短板移动,面积才可能会增加.
func maxArea(height []int) int { // faster 47.35% less 87.50%
	// 从两端开始寻找,至少保证了宽度是最大值
	i, j := 0, len(height)-1
	max := 0
	for i < j {
		a, b := height[i], height[j]
		h := min(a, b)
		area := h * (j - i)
		if max < area {
			max = area
		}
		// 朝着area具有变大的可能性方向变化
		if a < b {
			i++
		} else {
			j--
		}
	}
	return max
}
func min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}
*/

func maxArea(height []int) int { // 思路二改进版：faster 100% less 87.50%
	var tempA, tempI, tempJ int
	i := 0
	j := len(height) - 1
	h := height[i]
	if height[j] < height[i] {
		h = height[j]
	}
	area := (j - i) * h
	for i < j {
		if height[i] < height[tempI] {
			tempI = i
			i++
			continue
		}
		if height[j] < height[tempJ] {
			tempJ = j
			j--
			continue
		}
		h = height[i]
		if height[j] < height[i] {
			h = height[j]
		}
		tempA = (j - i) * h
		if tempA > area {
			area = tempA
		}
		if height[i] >= height[j] {
			tempJ = j
			j--
		} else {
			tempI = i
			i++
		}
	}
	return area
}

func main() {
	height := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	result := maxArea(height)
	fmt.Println(result)
}
