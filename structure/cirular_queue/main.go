package main

import "fmt"

type CirularQueue struct {
	array [4]int
	head  int
	tail  int
}

func (this *CirularQueue) Push(val int) {
	// fmt.Println((this.tail+1)%(len(this.array)+1), this.head)
	if (this.tail+1)%(len(this.array)+1) == this.head {
		panic("full")
	}
	this.array[this.tail] = val
	this.tail++
}
func (this *CirularQueue) Pop() {
	if this.head == this.tail {
		panic("empty")
	}
	fmt.Println(this.array[this.head])
	this.head++
}
func main() {
	cq := &CirularQueue{array: [4]int{}, head: 0, tail: 0}
	fmt.Println(cq)
	cq.Push(1)
	cq.Push(2)
	cq.Push(3)
	cq.Push(4)
	fmt.Println(cq)
	cq.Pop()
	cq.Pop()
	cq.Pop()
	cq.Pop()
	fmt.Println(cq)
}
