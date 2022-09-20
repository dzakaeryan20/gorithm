package bst

import "fmt"

func Hello(greeting string) {
	fmt.Println(greeting)
}

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

func (n *Node) Insert(k int) {
	if n.Key < k {
		// move right
		if n.Right == nil {
			n.Right = &Node{Key: k}
		} else {
			n.Right.Insert(k)
		}
	} else if n.Key > k {
		if n.Left == nil {
			n.Left = &Node{Key: k}
		} else {
			n.Left.Insert(k)
		}
	}
}

func (n *Node) Search(k int) bool {

	if n == nil {
		fmt.Println("masuk sini")
		return false
	}

	if n.Key < k {
		return n.Right.Search(k)
	} else if n.Key > k {
		return n.Left.Search(k)
	}

	return true
}

// 2*2*2*2*2*2
func (n *Node) Factorial(k int, num int) int {
	n.Left = &Node{Key: k}
	n.Right = &Node{Key: k}

	total := 0

	for i := 0; i < num; i += n.Key {
		if num-i == 1 {
			total = n.Left.Key * total
		} else {
			if total == 0 {
				total = n.Left.Key * n.Right.Key
			} else {
				total *= n.Left.Key * n.Right.Key
			}

		}

	}

	return total
}
