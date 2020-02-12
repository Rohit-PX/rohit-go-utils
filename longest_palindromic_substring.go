package main

import (
	"fmt"
)

func main() {
	str := "zbzd"
	fmt.Println(str)
	lps := longestPalindromicSub(str)
	fmt.Println(lps)
}

func longestPalindromicSub(testStr string) string {
	n := len(testStr)
	maxSub := ""

	if n == 0 {
		return ""
	}

	if n == 1 {
		return testStr
	}

	if n == 2 {
		if testStr[0] == testStr[1] {
			return testStr
		} else {
			return string(testStr[0])
		}
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if IsPalindrome(string(testStr[i:j+1])) && len(testStr[i:j+1]) > len(maxSub) {
				maxSub = testStr[i : j+1]
			}
		}
	}
	return maxSub

}

func IsPalindrome(str string) bool {
	mid := len(str) / 2
	last := len(str) - 1
	for i := 0; i < mid; i++ {
		if str[i] != str[last-i] {
			return false
		}
	}
	return true
}
