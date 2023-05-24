package main

import (
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
func (fb FooBar) foo(printFoo func()) {
	for i := 0; i < fb.n; i++ {
		fb.ch1 <- struct{}{}
		printFoo()
		<-fb.ch2
	}
}
func (fb FooBar) bar(printBar func()) {
	for i := 0; i < fb.n; i++ {
		fb.ch2 <- struct{}{}
		printBar()
		<-fb.ch1
	}
}
func main() {
	n := 2
	fb := NewFooBar(n)
	go fb.foo(func() { print("foo") })
	go fb.bar(func() { print("bar") })
	<-fb.ch1
	time.Sleep(time.Minute)
}
