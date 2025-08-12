package main

import (
	"container/heap"
	"fmt"
)

type MedianFinder struct {
	minHeap *MinHeap
	maxHeap *MaxHeap
}

// Implement Min Heap

type MinHeap []int

func (h MinHeap) Len() int { return len(h) }

// Logic to make it min heap is here
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }

func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j],h[i] }

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = (*h)[0:n-1]
	return x
}

func (h *MinHeap) Push(num interface{}) {
	*h = append(*h, num.(int))
}

func (h *MinHeap) Peek() int {
	if h.Len() < 1 {
		return -1
	}
	minHeapHead := (*h)[0]
	return minHeapHead
}

// Implement Max heap

type MaxHeap []int

func (h MaxHeap) Len() int { return len(h) }

// Logic to make it max heap is here
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }

func (h MaxHeap) Swap(i, j int) { h[i], h[j] = h[j],h[i] }

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := (*h)[n-1]
	*h = (*h)[0:n-1]
	return x
}

func (h *MaxHeap) Push(num interface{}) {
	*h = append(*h, num.(int))
}

func (h *MaxHeap) Peek() int {
	n := h.Len()
	if n < 1 {
		return -1
	}
	maxHeapHead := (*h)[0]
	return maxHeapHead
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	mf := &MedianFinder{}
	minH := MinHeap{}
	heap.Init(&minH)
	maxH := MaxHeap{}
	heap.Init(&maxH)
	mf.minHeap = &minH
	mf.maxHeap = &maxH
	return *mf
}


func (this *MedianFinder) AddNum(num int)  {
	if this.maxHeap.Len() == 0 &&  this.minHeap.Len() == 0  {
		// Since both heaps are empty add to max heap for now (we could also add to min heap if we wanted)
		heap.Push(this.minHeap,num)
		return
	} else if this.maxHeap.Len() == 0 {
		heap.Push(this.maxHeap,num)
	} else {
		heap.Push(this.minHeap,num)
	}
	// Check which heap should the current number go into
	//minHead := this.minHeap.Peek()
	maxHead := this.maxHeap.Peek()
    if num < maxHead {
		heap.Push(this.minHeap,num)
	} else {
		heap.Push(this.minHeap,num)
	}
	// Check if heaps are balanced in size (i.e don't differ in size by more than one)
	if Abs(this.minHeap.Len() - this.maxHeap.Len()) > 1 {
		if this.minHeap.Len() > this.maxHeap.Len() {
			popped := heap.Pop(this.minHeap)
			heap.Push(this.maxHeap,popped)

		} else {
			popped := heap.Pop(this.maxHeap)
			heap.Push(this.minHeap, popped)
		}
	}
}


func (this *MedianFinder) FindMedian() float64 {
	minHead := this.minHeap.Peek()
	maxHead := this.maxHeap.Peek()
	if this.minHeap.Len() == this.maxHeap.Len() {
		return float64(minHead+maxHead)/2
	} else if this.minHeap.Len() > this.maxHeap.Len() {
		return float64(minHead)
	}
	return float64(maxHead)
}

func (this *MedianFinder) PrintMedian()  {
	fmt.Printf("\nMin Heap: %v   Max Heap: %v", this.minHeap, this.maxHeap)
}


/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */

func main() {
	mf := Constructor()
	mf.AddNum(2)
	mf.PrintMedian()
	mf.AddNum(7)
	mf.PrintMedian()
	mf.AddNum(1)
	mf.PrintMedian()
	mf.AddNum(5)
	mf.PrintMedian()
	fmt.Printf("\nMedian: %v", mf.FindMedian())

	/*
	testMax := &MaxHeap{}
	heap.Push(testMax,2)
	heap.Push(testMax,1)
	heap.Push(testMax,5)
	fmt.Printf("\nMax peek: %d", testMax.Peek())

	testMin := &MinHeap{}
	heap.Push(testMin,2)
	heap.Push(testMin,1)
	heap.Push(testMin,5)
	fmt.Printf("\nMin peek: %d", testMin.Peek())

	 */
}


func Abs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}
