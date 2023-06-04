package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

type Stack []interface{}

func (s *Stack) Push(i interface{}) {
	*s = append(*s, i)
}

func (s *Stack) Pop() interface{} {
	if !(*s).IsEmpty() {
		top := (*s)[len(*s)-1]
		*s = (*s)[:len(*s)-1]
		return top
	}
	return nil
}

func (s *Stack) IsEmpty() bool {
	if len(*s) > 0 {
		return false
	}
	return true
}

func (s *Stack) Peek() interface{} {
	if !(*s).IsEmpty() {
		return (*s)[len(*s)-1]
	}
	return nil
}
