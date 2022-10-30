package main

import "fmt"

type Node struct {
	val   int
	left  *Node
	right *Node
}

func preOrder(node *Node) {
	if node != nil {
		fmt.Println(node.val)
		preOrder(node.left)
		preOrder(node.right)
	}
}
func infixOrder(node *Node) {
	if node != nil {
		infixOrder(node.left)
		fmt.Println(node.val)
		infixOrder(node.right)
	}
}
func postOrder(node *Node) {
	if node != nil {
		postOrder(node.left)
		postOrder(node.right)
		fmt.Println(node.val)
	}
}
func main() {
	root := &Node{val: 1}
	root.left = &Node{val: 2}
	root.right = &Node{val: 3}
	root.left.left = &Node{val: 4}
	root.left.right = &Node{val: 5}
	root.right.right = &Node{val: 6}
	// preOrder(root)
	infixOrder(root)
	// postOrder(root)
}
