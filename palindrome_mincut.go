package main

import "fmt"

const (
	problem = 132
)

func main() {
	fmt.Printf("LC%d. Palindrome Partitioning II", problem)
}

func minCutLC132(s string) int {
	fmt.Printf("LC132. Palindrome Partitioning II: %d", problem)
	n := len(s)
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}

	// Gap strategy
	for g := 0; g < n; g++ {
		for i, j := 0, g; j < n; i, j = i+1, j+1 {
			if g == 0 {
				dp[i][j] = true
			} else if g == 1 {
				dp[i][j] = (s[i] == s[j])
			} else {
				if (s[i] == s[j]) && dp[i+1][j-1] {
					dp[i][j] = true
				}
			}
		}
	}

	intDP := make([][]int, n)
	for i := 0; i < n; i++ {
		intDP[i] = make([]int, n)
	}

	for g := 0; g < n; g++ {
		for i, j := 0, g; j < n; i, j = i+1, j+1 {
			if g == 0 {
				intDP[i][j] = 0
			} else if g == 1 {
				if s[i] == s[j] {
					intDP[i][j] = 0
				} else {
					intDP[i][j] = 1
				}
			} else { //g > 1
				if dp[i][j] == true {
					// already a palindrome
					intDP[i][j] = 0
				} else {
					min := 9999999
					for k := i; k < j; k++ {
						left := intDP[i][k]
						right := intDP[k+1][j]
						if left+right+1 < min {
							min = left + right + 1
						}
					}
					intDP[i][j] = min
				}
			}
		}
	}

	fmt.Printf("MAP: %v", intDP)

	return intDP[0][n-1]

}
