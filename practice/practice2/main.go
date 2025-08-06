package main

import (
	"fmt"
	"strconv"
)

//回文

func main() {
	palindrome := isPalindrome(123)
	palindrome1 := isPalindrome(121)
	fmt.Println(palindrome, palindrome1)
}

func isPalindrome(x int) bool {
	str := strconv.Itoa(x)
	length := len(str)
	for i := 0; i < length/2; i++ {
		if str[i] != str[length-i-1] {
			return false
		}
	}
	return true
}
