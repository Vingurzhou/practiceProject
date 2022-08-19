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

func reverse(s string) string {
	runes := []rune(s)
	fmt.Println(runes)
	for front, back := 0, len(s)-1; front < back; front, back = front+1, back-1 {
		runes[front], runes[back] = runes[back], runes[front]
	}
	return string(runes)
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
	s1 = reverse(s1)
	s2 = reverse(s2)
	n1, _ := strconv.Atoi(s1)
	n2, _ := strconv.Atoi(s2)
	n3 := strconv.Itoa(n1 + n2)
	n3 = reverse(n3)
	lst := strings.Split(n3, "")
	return ll((lst), &ListNode{})
}
func main() {
	fmt.Println(addTwoNumbers(&ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}, &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}))
}
