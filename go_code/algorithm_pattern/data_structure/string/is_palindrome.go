package string

import "strings"

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	left, right := 0, len(s) - 1
	for left < right {
		for !isalnum(s[left]) {
			left++
		}
		for !isalnum(s[right]) {
			right--
		}
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func isalnum(ch byte) bool {
	return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')
}
