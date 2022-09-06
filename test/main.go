package main

import "fmt"

func addTwoNumbers(l1, l2 []int) (head []int) {
	carry := 0
	i := 0
	for (i < len(l1)) || (i < len(l2)) {
		n1, n2 := 0, 0
		if i < len(l1) {
			n1 = l1[i]
		}
		if i < len(l2) {
			n2 = l2[i]
		}
		sum := n1 + n2 + carry
		sum, carry = sum%10, sum/10
		head = append(head, sum)
		fmt.Println(sum)
		i++
	}
	return head
}
func main() {
	addTwoNumbers([]int{1, 2, 3, 4, 5, 6}, []int{1, 2, 3})
}
