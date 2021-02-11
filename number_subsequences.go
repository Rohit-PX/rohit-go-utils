package main

import (
	"fmt"
	"sort"
)

const (
	mod = 1000000007
)

func main() {
	arr := []int{9, 25, 9, 28, 24, 12, 17, 8, 28, 7, 21, 25, 10, 2, 16, 19, 12, 13, 15, 28, 14, 12, 24, 9, 6, 7, 2, 15, 19, 13, 30, 30, 23, 19, 11, 3, 17, 2, 14, 20, 22, 30, 12, 1, 11, 2, 2, 20, 20, 27, 15, 9, 10, 4, 12, 30, 13, 5, 2, 11, 29, 5, 3, 13, 22, 5, 16, 19, 7, 19, 11, 16, 11, 25, 29, 21, 29, 3, 2, 9, 20, 15, 9}
	fmt.Printf("Number of subsequences: %d\n", numSubseq(arr, 32))
}

func numSubseq(nums []int, target int) int {
	sort.Ints(nums)
	start := 0
	ans := 0
	end := len(nums) - 1
	for start <= end {
		if nums[start]+nums[end] <= target {
			ans = ans + int(powInt(2, int64(end-start)))
			ans = ans % mod
			start++
		} else {
			end--
		}
	}
	return ans
}

func powInt(x, y int64) int64 {

	if y == 1 {
		return x
	}
	if y == 0 {
		return 1
	}

	var ans int64
	ans = 1
	if y%2 == 0 {
		ans = powInt(x, y/2)
		ans = ans * ans

	} else {
		ans = powInt(x, y-1)
		ans = ans * x
	}

	return ans % mod
}
