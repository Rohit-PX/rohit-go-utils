package main

import (
	"fmt"
	"math"
)

var results [][]int

func main() {
	num := 16
	var currList []int
	helperFact(currList, num, 2)
	fmt.Printf("\n\n\nFINAL result: %v\n", results)
}

func helperFact(currList []int, n int, i int) {
	var j int
	for j = i; j <= int(math.Sqrt(float64(n))); j++ {
		var newList []int
		if (n % j) == 0 {
			newList = append(newList, currList...)
			newList = append(newList, j)
			newList = append(newList, int(n/j))
			results = append(results, newList)
			currList = append(currList, i)
			fmt.Printf("HELPER CALL: currlist=%v, n=%d, i=%d,n/j=%d, j=%d\n", currList, n, i, n/j, j)
			helperFact(currList, int(n/j), j)
		}
		//fmt.Printf("AFT CURR LIST: %v, n=%d, i=%d,j=%d\n", currList, n, i, j)
		currList = []int{}
	}
	fmt.Printf("For i=%d, result: %v\n", i, results)
}
