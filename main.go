package main

import (
	"fmt"

	"github.com/dzakaeryan20/gorithm/bst"
)

func main() {
	// bst.Hello("Welcome to Bst")

	// binary search tree
	tree := &bst.Node{Key: 100}
	tree.Insert(50)
	tree.Insert(200)
	tree.Insert(150)
	tree.Insert(180)

	// fmt.Println(tree)
	// fmt.Println(tree.Search(50))

	pangkat := &bst.Node{Key: 2}
	t := pangkat.Factorial(3, 4)
	fmt.Println(t)
	//pangkat.Factorial(3,4)
}
