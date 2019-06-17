package main

import (
	"fmt"
)

/*
思路一：

*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {

}

func main() {
	var list []*ListNode
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
