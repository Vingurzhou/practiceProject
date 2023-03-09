/**
 * Created by zhouwenzhe on 2023/2/11
 */

package leetcode

import (
	"fmt"
)

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var tail *ListNode
	var head *ListNode
	var carry int
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}

		val := (n1 + n2 + carry) % 10
		fmt.Println(val)
		if head == nil {
			head = &ListNode{Val: val, Next: nil}
			tail = head
		} else {
			tail.Next = &ListNode{Val: val, Next: nil}
			tail = tail.Next
		}

		carry = (n1 + n2 + carry) / 10
	}
	if carry > 0 {
		tail.Next = &ListNode{Val: carry}
	}
	fmt.Println(head)
	return head
}
