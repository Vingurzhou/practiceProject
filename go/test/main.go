package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func ll(nums []int, head *ListNode) *ListNode {
	fnode := head
	for _, num := range nums {
		temp := ListNode{Val: num}
		head.Next = &temp
		head = &temp
	}
	return fnode.Next
}
func main() {
	a := []int{1, 23, 4, 5}
	fmt.Println(ll(a, &ListNode{}))
}
