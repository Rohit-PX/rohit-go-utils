package main

import (
	"fmt"
)

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

func main() {
	root := Node{Data: 1}
	left := Node{Data: 2}
	right := Node{Data: 3}

	root.Left = &left
	root.Right = &right

	inOrder(&root)

}

func createNode() {
}

func inOrder(root *Node) {
	if root == nil {
		return
	}

	inOrder(root.Left)
	fmt.Printf("%d\n", root.Data)
	inOrder(root.Right)

}
