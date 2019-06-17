package main

import (
	"fmt"
)

/*
思路一：
遍历放入slice后排序，再遍历slice拼成ListNode。
func mergeKLists(lists []*ListNode) *ListNode { // faster 98.81% less 38.48%
	var slice []int
	for k := range lists {
		temp := lists[k]
		if temp == nil {
			continue
		}
		for temp.Next != nil {
			slice = append(slice, temp.Val)
			temp = temp.Next
		}
		slice = append(slice, temp.Val)
	}
	if len(slice) == 0 {
		return nil
	}
	sort.Ints(slice)
	l := new(ListNode)
	result := l // 因l一直被迭代,若没有此句,则最终返回的l是最新被初始化迭代的
	for m := range slice {
		l.Val = slice[m]
		if m == len(slice)-1 {
			break
		}
		l.Next = new(ListNode)
		l = l.Next
	}
	return result
}
*/

/*
思路二：
归并，两两合并、递归。
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func merge(lists []*ListNode) *ListNode { // faster 98.81% less 100%
	length := len(lists)
	half := length / 2
	if length == 1 {
		return lists[0]
	}
	if length == 2 {
		var (
			l0, l1   = lists[0], lists[1]
			res, cur *ListNode
		)
		if l0 == nil {
			return l1
		}
		if l1 == nil {
			return l0
		}
		if l0.Val < l1.Val {
			res, cur, l0 = l0, l0, l0.Next
		} else {
			res, cur, l1 = l1, l1, l1.Next
		}
		for l0 != nil && l1 != nil {
			if l0.Val < l1.Val {
				cur.Next, l0 = l0, l0.Next
			} else {
				cur.Next, l1 = l1, l1.Next
			}
			cur = cur.Next
		}
		if l0 != nil {
			cur.Next = l0
		}
		if l1 != nil {
			cur.Next = l1
		}
		return res
	}
	return mergeKLists([]*ListNode{mergeKLists(lists[half:]), mergeKLists(lists[:half])})
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	return merge(lists)
}

func main() {
	var list []*ListNode
	l1 := new(ListNode)
	l1.Val = 1
	l1.Next = new(ListNode)
	l1.Next.Val = 4
	l1.Next.Next = new(ListNode)
	l1.Next.Next.Val = 7
	list = append(list, l1)
	l2 := new(ListNode)
	l2.Val = 2
	l2.Next = new(ListNode)
	l2.Next.Val = 5
	l2.Next.Next = new(ListNode)
	l2.Next.Next.Val = 8
	list = append(list, l2)
	l3 := new(ListNode)
	l3.Val = 3
	l3.Next = new(ListNode)
	l3.Next.Val = 6
	l3.Next.Next = new(ListNode)
	l3.Next.Next.Val = 9
	list = append(list, l3)
	result := mergeKLists(list)
	fmt.Println(result)
	for (result.Next != nil) {
		fmt.Println(result.Val)
		result = result.Next
	}
	fmt.Println(result.Val)
}
