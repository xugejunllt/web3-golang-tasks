package task1

/**
给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。

回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

**/

import (
	"fmt"
	"strconv"
	"testing"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	s := strconv.Itoa(x)
	return s == reverseString(s)
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func TestIsPalindrome(t *testing.T) {
	test := 123454321
	result := isPalindrome(test)
	fmt.Print(test, "是否回文数:", result)
}
