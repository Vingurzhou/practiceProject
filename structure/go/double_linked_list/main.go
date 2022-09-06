package main

import "fmt"

type ListNode struct {
	val  int
	pre  *ListNode
	next *ListNode
}

func (this *ListNode) Insert(ln *ListNode) {
	for this.next != nil {
		this = this.next

	}
	this.next = ln
	ln.pre = this
}
func (this *ListNode) Show() {
	for this.next != nil {
		fmt.Print(this, "\t")
		this = this.next
	}
	if this.next == nil {
		fmt.Println(this)
	}

}
func (this *ListNode) Delete(val int) {
	for this.next != nil {
		if this.val == val {
			this.next.pre = this.pre
			this.pre.next = this.next
		}
		this = this.next
	}
	if this.next == nil {
	}
}
func main() {
	head := &ListNode{}
	head.Show()
	head.Insert(&ListNode{val: 1})
	head.Insert(&ListNode{val: 2})
	head.Insert(&ListNode{val: 3})
	head.Show()
	head.Delete(2)
	head.Show()
}
