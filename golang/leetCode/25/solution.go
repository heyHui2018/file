package main

import (
	"fmt"
)

/*
思路一：
需反转k个节点。将各节点指针放入list，当满足k个节点的数量时（count%k == 0），将这一组进行反转（第一个和第k个，第二个和第k-1个。。。，用for循环），需注意当head的长度是k的倍数时，
因最后一个节点的.next==nil，故需在最后再次手动加入list并执行一次反转
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

// 翻转k个节点,并非第1个和第k个互换
func reverseKGroup(head *ListNode, k int) *ListNode { // faster 25.64% less 6.87%
	if head == nil || head.Next == nil || k == 1 {
		return head
	}
	copy := head
	count := 1
	var list []*ListNode
	for copy.Next != nil {
		list = append(list, copy)
		if count%k == 0 {
			for i := 0; i < k/2; i++ {
				list[len(list)-1-i].Val, list[len(list)-k+i].Val = list[len(list)-k+i].Val, list[len(list)-1-i].Val
			}
		}
		count++
		copy = copy.Next
	}
	list = append(list, copy)
	if len(list)%k == 0 {
		for i := 0; i < k/2; i++ {
			list[len(list)-1-i].Val, list[len(list)-k+i].Val = list[len(list)-k+i].Val, list[len(list)-1-i].Val
		}
	}
	return head
}

func main() {
	l1 := new(ListNode)
	l1.Val = 1
	l1.Next = new(ListNode)
	l1.Next.Val = 2
	// l1.Next.Next = new(ListNode)
	// l1.Next.Next.Val = 3
	// l1.Next.Next.Next = new(ListNode)
	// l1.Next.Next.Next.Val = 4
	// l1.Next.Next.Next.Next = new(ListNode)
	// l1.Next.Next.Next.Next.Val = 5
	result := reverseKGroup(l1, 2)
	for (result.Next != nil) {
		fmt.Println(result.Val)
		result = result.Next
	}
	fmt.Println(result.Val)
}
