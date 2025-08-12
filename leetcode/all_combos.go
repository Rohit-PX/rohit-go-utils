package main

import (
	"fmt"
	"os"
	"strconv"
)

func getNumComb(n int) int {
	count := 0
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			if i*j <= n {
				count = count + 1
				fmt.Printf("\nCurrent combo: (%d,%d)", i, j)
			}
		}
	}
	return count*2 + 3
}

func main() {
	nodes, _ := strconv.Atoi(os.Args[1])
	fmt.Printf("\nPossible combos for %d nodes: %d\n", nodes, getNumComb(nodes))
}
