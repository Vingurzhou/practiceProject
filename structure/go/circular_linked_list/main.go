package main

import "fmt"

type ListNode struct {
	val  int
	next *ListNode
}

func (this *ListNode) Insert(ln *ListNode) {
	if this.next == nil {
		this.val = ln.val
		this.next = ln
	}
	for this.next != nil {
		this = this.next
	}
}
func (this *ListNode) Show() {}
func main() {
	head := &ListNode{}
	fmt.Println(head)
	head.Insert(&ListNode{val: 1})
	fmt.Println(head)
}
