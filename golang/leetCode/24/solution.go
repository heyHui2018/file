package main

import (
	"fmt"
)

/*
思路一：
计数，是偶数位时，与前一位互换。
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode { // faster 100.00% less 67.41%
	if head == nil || head.Next == nil {
		return head
	}
	copy := head
	count := 1
	for copy.Next != nil {
		count++
		if count%2 == 0 {
			copy.Val, copy.Next.Val = copy.Next.Val, copy.Val
		}
		copy = copy.Next
	}
	return head
}

func main() {
	l1 := new(ListNode)
	l1.Val = 1
	l1.Next = new(ListNode)
	l1.Next.Val = 4
	l1.Next.Next = new(ListNode)
	l1.Next.Next.Val = 7
	result := swapPairs(l1)
	fmt.Println(result)
	for (result.Next != nil) {
		fmt.Println(result.Val)
		result = result.Next
	}
	fmt.Println(result.Val)
}
