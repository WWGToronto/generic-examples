package main

import "fmt"

// Node is a very simple implementation of a binary tree.
type Node[T any] struct {
	Value T
	Left  *Node[T]
	Right *Node[T]
}

func (n *Node[T]) CreateLeft(v T) {
	n.Left = &Node[T]{Value: v}
}

func (n *Node[T]) CreateRight(v T) {
	n.Right = &Node[T]{Value: v}
}

func (n *Node[T]) Print(prefix string) {
	fmt.Printf("%s<%v>\n", prefix, n.Value)
	if n.Right != nil {
		n.Right.Print(prefix + "    ")
	}
	if n.Left != nil {
		n.Left.Print(prefix + "    ")
	}
}

// main contains an example of how you can write and use custom data types.
func main() {
	// build a tree of pets
	binaryTree := Node[string]{Value: "pet"}
	binaryTree.CreateLeft("cat")
	binaryTree.Left.CreateLeft("calico")
	binaryTree.Left.CreateRight("black")
	binaryTree.Left.Right.CreateLeft("big")
	binaryTree.Left.Right.CreateRight("small")
	binaryTree.CreateRight("dog")
	binaryTree.Right.CreateLeft("poodle")
	binaryTree.Right.CreateRight("husky")

	// display!
	binaryTree.Print("")
}
