package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
思路一：
l1、l2遍历直接相加后放入l3,有进位时放入temp,最后遍历l3输出
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var s3 []int
	var temp, flag int
Exit:
	for {
		switch {
		case l1 == nil && l2 == nil:
			if flag == 1 {
				s3 = append(s3, 1)
			}
			break Exit
		case l1 == nil:
			temp = l2.Val + flag
			l2 = l2.Next
		case l2 == nil:
			temp = l1.Val + flag
			l1 = l1.Next
		default:
			temp = l1.Val + l2.Val + flag
			l1 = l1.Next
			l2 = l2.Next
		}
		flag = 0
		if temp >= 10 {
			flag = 1
			temp = temp - 10
		}
		s3 = append(s3, temp)
	}
	var l3 *ListNode
	for i := 1; i <= len(s3); i++ {
		intTemp := s3[len(s3)-i]
		temp := new(ListNode)
		temp.Val = intTemp
		temp.Next = l3
		l3 = temp
	}
	return l3
}
*/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	resPre := &ListNode{}
	cur := resPre
	carry := 0

	for l1 != nil || l2 != nil || carry > 0 {
		sum := carry

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		carry = sum / 10

		cur.Next = &ListNode{Val: sum % 10}
		cur = cur.Next
	}

	return resPre.Next
}

func main() {
	l1 := new(ListNode)
	l1.Val = 5
	l1.Next = new(ListNode)
	l1.Next.Val = 4
	l1.Next.Next = new(ListNode)
	l1.Next.Next.Val = 3

	l2 := new(ListNode)
	l2.Val = 5
	l2.Next = new(ListNode)
	l2.Next.Val = 9
	l2.Next.Next = new(ListNode)
	l2.Next.Next.Val = 4

	result := addTwoNumbers(l1, l2)
	for {
		if result.Next == nil {
			fmt.Println("111111", result.Val)
			break
		}
		fmt.Println(result.Val)
		result = result.Next
	}
}
