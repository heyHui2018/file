package main

import (
	"fmt"
)

/*
思路一：
遍历链表并存入map,key为节点序号,value为节点指针,根据入参n,将总数count-n+2链到count-n上.
需注意单节点情况和删除头节点的情况.
func removeNthFromEnd(head *ListNode, n int) *ListNode { // faster 100% less 6.02%
	headMap := make(map[int]*ListNode)
	count := 1
	for (head.Next != nil) {
		headMap[count] = head
		count++
		head = head.Next
	}
	headMap[count] = head
	if count == 1 {
		return nil
	}
	if count == n {
		return headMap[2]
	}
	headMap[count-n].Next = headMap[count-n+2]
	return headMap[1]
}
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode { // faster 100% less 81.12%
	d, headIsNthFromEnd := getDaddy(head, n)
	if headIsNthFromEnd {
		// 删除head节点
		return head.Next
	}
	d.Next = d.Next.Next
	return head
}

// 获取倒数第N个节点的父节点
func getDaddy(head *ListNode, n int) (daddy *ListNode, headIsNthFromEnd bool) {
	daddy = head
	for head != nil {
		if n < 0 {
			daddy = daddy.Next // n < 0后,daddy从头开始自己变.(n=2,总6,则为正数第5即总+1-n,n--减到0时, daddy从头往后迭代,当head迭代完成时,daddy刚好指向目标的上一个节点)
		}
		n--
		head = head.Next
	}
	// n==0，说明链的长度等于n
	headIsNthFromEnd = n == 0
	return
}

func main() {
	l1 := new(ListNode)
	l1.Val = 1
	l2 := new(ListNode)
	l2.Val = 2
	l1.Next = l2
	l3 := new(ListNode)
	l3.Val = 3
	l2.Next = l3
	l4 := new(ListNode)
	l4.Val = 4
	l3.Next = l4
	l5 := new(ListNode)
	l5.Val = 5
	l4.Next = l5
	l6 := new(ListNode)
	l6.Val = 6
	l5.Next = l6
	result := removeNthFromEnd(l1, 2)
	if result != nil {
		for (result.Next != nil) {
			fmt.Println(result.Val)
			result = result.Next
		}
		fmt.Println(result.Val)
	}
}
