package main

import (
	"fmt"
	"time"
)

type FooBar struct {
	n   int
	ch1 chan struct{}
	ch2 chan struct{}
}

func NewFooBar(n int) *FooBar {
	ch1 := make(chan struct{})
	ch2 := make(chan struct{})
	return &FooBar{n,
		ch1,
		ch2,
	}
}
func (fb FooBar) foo() {
	for i := 0; i < fb.n; i++ {
		fb.ch1 <- struct{}{}
		printFoo()
		<-fb.ch2
	}
}
func (fb FooBar) bar() {
	for i := 0; i < fb.n; i++ {
		fb.ch2 <- struct{}{}
		printBar()
		<-fb.ch1
	}
}
func printFoo() {
	fmt.Println("foo")
}
func printBar() {
	fmt.Println("bar")
}
func main() {
	fb := NewFooBar(100)
	go fb.foo()
	go fb.bar()
	<-fb.ch1
	time.Sleep(time.Minute)
}
