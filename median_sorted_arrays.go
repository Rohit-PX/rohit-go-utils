package main

import "fmt"

func main() {
	arr1 := []int{7, 11, 18, 19, 21, 25}
	arr2 := []int{1, 3, 8, 9, 15}
	fmt.Printf("\n\nMedian of %v and \n %v\n is: %f\n", arr1, arr2, findMedianSortedArrays(arr1, arr2))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) { // Left array should be smaller so call with reversed arguments
		return findMedianSortedArrays(nums2, nums1)
	}
	x := len(nums1)
	y := len(nums2)

	low := 0
	high := x
	for low <= high {
		partX := (low + high) / 2
		partY := (x+y+1)/2 - partX

		maxLeftX := getMaxLeft(nums1, partX)
		minRightX := getMinRight(nums1, partX)

		maxLeftY := getMaxLeft(nums2, partY)
		minRightY := getMinRight(nums2, partY)

		//fmt.Printf("low: %d \nhigh: %d \npartX: %d \npartY: %d\n", low, high, partX, partY)
		//fmt.Printf("maxLeftX: %d \nminRightX: %d \nmaxLeftY: %d \nminRightY: %d", maxLeftX, minRightX, maxLeftY, minRightY)

		if (maxLeftX <= minRightY) && (maxLeftY <= minRightX) {
			// check odd or even total length
			if (x+y)%2 == 0 {
				return float64(Max(maxLeftX, maxLeftY)+Min(minRightX, minRightY)) / 2
			} else {
				return float64(Max(maxLeftX, maxLeftY))
			}
		} else if maxLeftX > minRightY {
			high = partX - 1
		} else {
			low = partX + 1
		}

	}
	return 1.38
}

func getMaxLeft(arr []int, p int) int {
	var m int
	if p == 0 {
		m = -1000
	} else {
		m = arr[p-1]
	}
	return m
}

func getMinRight(arr []int, p int) int {
	var m int
	if p == len(arr) {
		m = 12345
	} else {
		m = arr[p]
	}
	return m
}
func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
