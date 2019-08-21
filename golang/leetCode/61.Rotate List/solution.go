package main

import (
	"fmt"
)

/*
要求：


解题思路：
1、链表转成数组,简化k(k = k % len(list))后翻转,再转成链表.
2、取出链表前k个数,在此期间若next==nil,则可知链表长度,进入简化k(k = k % len(list))使其小于链表长度,随后再次取出链表前k个数,将剩余链表加入结果,随后将链表首位相连,再截断一位

关键点：


*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight1(head *ListNode, k int) *ListNode { // faster 31.91% less 25.00%.
	if k == 0 || head == nil {
		return head
	}
	first := head
	temp := head
	list := []int{}
	for {
		list = append(list, head.Val)
		if head.Next == nil {
			break
		}
		head = head.Next
	}
	k = k % len(list)
	res := []int{}
	for i := 0; i < k; i++ {
		res = append(res, list[len(list)-k+i])
	}
	for i := 0; i < len(list)-k; i++ {
		res = append(res, list[i])
	}
	for i := 0; i < len(res); i++ {
		first.Val = res[i]
		first = first.Next
	}
	return temp
}

func rotateRight(head *ListNode, k int) *ListNode { // faster 100% less 75.00%.
	if k == 0 || head == nil {
		return head
	}

	tail := head
	for i := 0; i < k; i++ {
		if tail.Next == nil {
			// 处理 k 大于 list 的长度的情况
			// i+1 就是 list 的长度
			return rotateRight(head, k%(i+1))
		}
		tail = tail.Next
	}

	newTail := head
	for tail.Next != nil {
		newTail, tail = newTail.Next, tail.Next
	}

	newHead := newTail.Next
	// tail.Next=head 意为首尾相连形成环
	// newTail.Next=nil 意为往后去一位
	newTail.Next, tail.Next = nil, head

	return newHead
}

func main() {
	l1 := new(ListNode)
	l1.Val = 1
	l1.Next = new(ListNode)
	l1.Next.Val = 4
	l1.Next.Next = new(ListNode)
	l1.Next.Next.Val = 7
	l1.Next.Next.Next = new(ListNode)
	l1.Next.Next.Next.Val = 8
	k := 7
	result := rotateRight(l1, k)
	for result.Next != nil {
		fmt.Println(result)
		result = result.Next
	}
	fmt.Println(result)
}
