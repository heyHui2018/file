package main

import (
	"fmt"
	"sort"
)

/*
思路一：
遍历放入slice后排序，再遍历slice拼成ListNode。
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode { // faster 28.21% less 5.11%
	var slice []int
	for l1 != nil || l2 != nil {
		if l1 != nil {
			slice = append(slice, l1.Val)
			l1 = l1.Next
		}
		if l2 != nil {
			slice = append(slice, l2.Val)
			l2 = l2.Next
		}
	}
	if len(slice) == 0 {
		return nil
	}
	sort.Ints(slice)
	temp := new(ListNode)
	result := temp
	for k, v := range slice {
		temp.Val = v
		if k == len(slice)-1 {
			break
		}
		temp.Next = new(ListNode)
		temp = temp.Next
	}
	return result
}
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode { // faster 100.00% less 100.00%
	// 有一条链为nil，直接返回另一条链
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	// 此时，两条链都不为nil，可以直接使用l.Val，而不用担心panic
	// 在l1和l2之间，选择较小的节点作为head，并设置好node
	var head, node *ListNode
	if l1.Val < l2.Val {
		head = l1
		node = l1
		l1 = l1.Next
	} else {
		head = l2
		node = l2
		l2 = l2.Next
	}

	// 循环比较l1和l2的值，始终选择较小的值连上node
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			node.Next = l1
			l1 = l1.Next
		} else {
			node.Next = l2
			l2 = l2.Next
		}
		node = node.Next
	}
	if l1 != nil {
		// 连上l1剩余的链
		node.Next = l1
	}
	if l2 != nil {
		// 连上l2剩余的链
		node.Next = l2
	}
	return head
}

func main() {
	var l1 *ListNode
	l1 = nil
	var l2 *ListNode
	l2 = nil
	// l1 := new(ListNode)
	// l1.Val = 2
	// l1.Next = new(ListNode)
	// l1.Next.Val = 6
	// l1.Next.Next = new(ListNode)
	// l1.Next.Next.Val = 7

	// l2 := new(ListNode)
	// l2.Val = 4
	// l2.Next = new(ListNode)
	// l2.Next.Val = 5
	// l2.Next.Next = new(ListNode)
	// l2.Next.Next.Val = 9
	result := mergeTwoLists(l1, l2)
	fmt.Println(result)
	for (result.Next != nil) {
		fmt.Println(result.Val)
		result = result.Next
	}
	fmt.Println(result.Val)
}
