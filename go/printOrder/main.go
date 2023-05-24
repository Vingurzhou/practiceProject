package main

import (
	"time"
)

type Foo struct {
	firstJobDone  chan struct{}
	secondJobDone chan struct{}
}

func NewFoo() *Foo {
	firstJobDone := make(chan struct{})
	secondJobDone := make(chan struct{})
	return &Foo{firstJobDone: firstJobDone,
		secondJobDone: secondJobDone,
	}
}

func (f *Foo) first(printFirst func()) {
	printFirst()
	f.firstJobDone <- struct{}{}
}

func (f *Foo) second(printSecond func()) {
	<-f.firstJobDone
	printSecond()
	f.secondJobDone <- struct{}{}
}

func (f *Foo) third(printThird func()) {
	<-f.secondJobDone
	printThird()
}

func main() {
	nums := []int{1, 3, 2}
	f := NewFoo()
	for _, num := range nums {
		switch num {
		case 1:
			go f.first(func() { print("first") })
		case 2:
			go f.second(func() { print("second") })
		case 3:
			go f.third(func() { print("third") })
		default:
			panic(num)
		}
	}

	time.Sleep(time.Minute)
}
