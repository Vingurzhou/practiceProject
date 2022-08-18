package main

import (
	"fmt"
	"strconv"
	"strings"
)

//   Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func ll(nums []string, head *ListNode) *ListNode {
	fnode := head
	for _, num := range nums {
		x, _ := strconv.Atoi(num)
		temp := ListNode{Val: x}
		head.Next = &temp
		head = &temp
	}
	return fnode.Next
}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	s1 := ""
	for l1 != nil {
		s1 += strconv.Itoa(l1.Val)
		l1 = l1.Next
	}
	s2 := ""
	for l2 != nil {
		s2 += strconv.Itoa(l2.Val)
		l2 = l2.Next
	}

	n1, _ := strconv.Atoi(s1)
	n2, _ := strconv.Atoi(s2)
	n3 := strconv.Itoa(n1 + n2)
	lst := strings.Split(n3, "")
	for i, j := 0, len(lst)-1; i < j; i, j = i+1, j-1 {
		lst[i], lst[j] = lst[j], lst[i]
	}
	return ll(lst, &ListNode{})
}
func main() {
	fmt.Println(addTwoNumbers(&ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}, &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}))
}
