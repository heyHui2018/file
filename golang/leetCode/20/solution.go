package main

import (
	"errors"
	"fmt"
)

/*
思路一：
分解后遍历s,当是前半个时,入栈,后半个时和栈头匹配,成功则继续,不成功则返回false.最后需检查栈的长度,>0时返回false.
func isValid(s string) bool { // faster 100% less 13.07%
	cMap := map[string]string{
		")": "(",
		"]": "[",
		"}": "{",
	}
	var stack Stack
	sList := strings.Split(s, "")
	for _, v := range sList {
		if _, ok := cMap[v]; ok {
			character, err := stack.Pop()
			if err != nil {
				return false
			}
			if character.(string) != cMap[v] {
				return false
			}
		} else {
			stack.Push(v)
		}
	}
	if stack.Len() > 0 {
		return false
	}
	return true
}
*/

func isValid(s string) bool { // faster 100% less 83.97%
	size := len(s)

	stack := make([]byte, size) // 简易栈
	top := 0

	for i := 0; i < size; i++ {
		c := s[i]
		switch c {
		case '(':
			stack[top] = c + 1 // '('+1 is ')'
			top++
		case '[', '{':
			stack[top] = c + 2
			top++
		case ')', ']', '}':
			if top > 0 && stack[top-1] == c {
				top--
			} else {
				return false
			}
		}
	}
	return top == 0
}

func main() {
	s := "["
	result := isValid(s)
	fmt.Println(result)
}

type Stack []interface{}

func (stack Stack) Len() int {
	return len(stack)
}
func (stack Stack) Cap() int {
	return cap(stack)
}
func (stack *Stack) Push(value interface{}) {
	*stack = append(*stack, value)
}
func (stack Stack) Top() (interface{}, error) {
	if len(stack) == 0 {
		return nil, errors.New("Out of index, len is 0")
	}
	return stack[len(stack)-1], nil
}

func (stack *Stack) Pop() (interface{}, error) {
	theStack := *stack
	if len(theStack) == 0 {
		return nil, errors.New("Out of index, len is 0")
	}
	value := theStack[len(theStack)-1]
	*stack = theStack[:len(theStack)-1]
	return value, nil
}

func (stack Stack) IsEmpty() bool {
	return len(stack) == 0
}
