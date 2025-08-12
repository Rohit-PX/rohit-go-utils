package main

import (
	"fmt"
	"sort"
)

type Interval struct {
	Start int
	End   int
}
type Intervals []*Interval

func (a Intervals) Len() int           { return len(a) }
func (a Intervals) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Intervals) Less(i, j int) bool { return a[i].Start < a[j].Start }

func main() {
	i1 := &Interval{Start: 1, End: 2}
	i2 := &Interval{Start: 7, End: 8}
	i3 := &Interval{Start: 5, End: 8}

	var testIntervals Intervals
	testIntervals = append(testIntervals, i1)
	testIntervals = append(testIntervals, i2)
	testIntervals = append(testIntervals, i3)

	for _, i := range getFreeTime(testIntervals) {
		fmt.Printf("\n%d,%d", i.Start, i.End)
	}

	fmt.Println("\nBEFORE")
	for _, i := range testIntervals {
		fmt.Printf("\n%d,%d", i.Start, i.End)
	}

	sort.Sort(testIntervals)
	fmt.Println("\nAFTER")
	for _, i := range testIntervals {
		fmt.Printf("\n%d,%d", i.Start, i.End)
	}

}

func getFreeTime(intervals []*Interval) []*Interval {
	var freeTimes []*Interval
	for i := 0; i < len(intervals)-1; i++ {
		if intervals[i].End < intervals[i+1].Start {
			freeInterval := &Interval{
				Start: intervals[i].End,
				End:   intervals[i+1].Start,
			}
			freeTimes = append(freeTimes, freeInterval)
		}
	}
	return freeTimes

}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
