package main

import (
	"fmt"
)

type Queue struct {
	array [4]int
	front int
	rear  int
}

func (this *Queue) Add(val int) {
	if this.rear == len(this.array)-1 {
		panic("queue full")
	}
	this.rear++
	this.array[this.rear] = val
	return
}
func (this *Queue) Get() {
	if this.rear == this.front {
		panic("queue empty")
	}
	this.front++
	fmt.Println(this.array[this.front])
	return
}

func main() {
	queue := Queue{array: [4]int{},
		front: -1, rear: -1}
	fmt.Println(queue)
	queue.Add(1)
	queue.Add(2)
	queue.Add(3)
	queue.Add(4)
	fmt.Println(queue)
	queue.Get()
	fmt.Println(queue)
	queue.Get()
	fmt.Println(queue)

}
