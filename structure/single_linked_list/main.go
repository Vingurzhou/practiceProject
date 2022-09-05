package main

import "fmt"

type ListNode struct {
	val  int
	next *ListNode
}

func (this *ListNode) insert(ln *ListNode) {
	temp := this
	for temp.next != nil {
		temp = temp.next
	}
	temp.next = ln
}
func (this *ListNode) show() {
	temp := this
	for temp.next != nil {
		fmt.Println(temp)
		temp = temp.next
	}
	if temp.next == nil {
		fmt.Println(temp)
	}

}
func main() {
	head := &ListNode{}
	head.show()
	// head.insert(&ListNode{val: 1})
	// head.show()
}
