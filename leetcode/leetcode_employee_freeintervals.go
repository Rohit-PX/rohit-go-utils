package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

type Interval struct {
	Start int
	End   int
}

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

type Intervals []*Interval

func (a Intervals) Len() int           { return len(a) }
func (a Intervals) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Intervals) Less(i, j int) bool { return a[i].Start < a[j].Start }

func employeeFreeTime(schedule [][]*Interval) []*Interval {

	var allTimes Intervals
	var commonFreeTimes Intervals

	for _, s := range schedule {
		for _, i := range s {
			allTimes = append(allTimes, i)
		}
	}
	sort.Sort(allTimes)

	end := allTimes[0].End
	for i := 1; i < len(allTimes); i++ {
		if end < allTimes[i].Start {
			freeInterval := &Interval{Start: end, End: allTimes[i].Start}
			commonFreeTimes = append(commonFreeTimes, freeInterval)
		}
		end = Max(end, allTimes[i].End)
	}

	return commonFreeTimes
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
