package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	a := []int{1, 2, 3, 4, 5}
	n := len(a)

	argsWithProg := os.Args
	pivot, _ := (strconv.Atoi(argsWithProg[1]))
	pivot = pivot % n

	b := make([]int, len(a))

	for j := 0; j < len(a); j++ {
		idx := (n - pivot + j) % n
		b[j] = a[idx]
	}
	fmt.Printf("A: %v", a)
	fmt.Printf("\nB: %v\n", b)

}
