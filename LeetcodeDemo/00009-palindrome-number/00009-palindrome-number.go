package _0009_palindrome_number

import "strconv"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	s := strconv.Itoa(x)
	sLen := len(s)
	fLen := sLen / 2
	for i := 0; i <= fLen; i++ {
		if s[i] != s[sLen-1-i] {
			return false
		}
	}
	return true
}
