package main

import "fmt"

// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效
func main() {
	fmt.Println(isValid("{}{}[][]()()"))
	fmt.Println(isValid("{{{([])}}}"))
	fmt.Println(isValid("{{{[])}}}"))
}

func isValid(s string) bool {
	if len(s)%2 != 0 { // s 长度必须是偶数
		return false
	}
	m := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	stack := []rune{}
	for _, c := range s {
		if m[c] == 0 {
			stack = append(stack, c) //入栈
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != m[c] {
				return false
			}
			stack = stack[:len(stack)-1] //出栈
		}
	}
	return len(stack) == 0
}
