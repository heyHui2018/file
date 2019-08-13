package main

import (
	"fmt"
)

var array = []int{10, 6, 1, 2, 7, 9, 3, 4, 5, 8}

func main() {
	fmt.Println("排序前:", array)
	heapSort(array)
	fmt.Println("排序后:", array)
}

func heapSort(array []int) {
	N := len(array) - 1
	//构造堆
	//如果给两个已构造好的堆添加一个共同父节点，
	//将新添加的节点作一次下沉将构造一个新堆，
	//由于叶子节点都可看作一个构造好的堆，所以
	//可以从最后一个非叶子节点开始下沉，直至
	//根节点，最后一个非叶子节点是最后一个叶子
	//节点的父节点，角标为N/2
	for k := N / 2; k >= 0; k-- {
		sink(array, k, N)
	}
	//下沉排序
	for N > 0 {
		//将大的放在数组后面，升序排序
		array[0], array[N] = array[N], array[0]
		N--
		sink(array, 0, N)
	}
}

//下沉（由上至下的堆有序化）
func sink(s []int, k, N int) {
	for {
		i := 2 * k
		if i > N { //保证该节点是非叶子节点
			break
		}
		if i < N && s[i+1] > s[i] { //选择较大的子节点
			i++
		}
		if s[k] >= s[i] { //没下沉到底就构造好堆了
			break
		}
		s[k], s[i] = s[i], s[k]
		k = i
	}
}