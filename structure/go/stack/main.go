package main

import "fmt"

type Stack struct {
	top   int
	array [5]int
}

func (this *Stack) push(val int) {
	if len(this.array)-1 == this.top {
		panic("stack full")
	}
	this.top++
	this.array[this.top] = val
}
func (this *Stack) pop(val int) {
	if -1 == this.top {
		panic("stack empty")
	}
	this.array[this.top] = val
	this.top--

}
func main() {
	stack := &Stack{
		top: -1,
	}
	fmt.Println(stack)
	stack.push(1)
	stack.push(1)
	stack.push(1)
	stack.push(1)
	stack.push(1)
	fmt.Println(stack)
	// stack.push(1)
	stack.pop(1)
	stack.pop(1)
	stack.pop(1)
	stack.pop(1)
	stack.pop(1)
	fmt.Println(stack)
	// stack.pop(1)
}
