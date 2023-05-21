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

func (f *Foo) first() {
	//printFirst()
	//<-f.firstJobDone
	printFirst()
	f.firstJobDone <- struct{}{}
}

func (f *Foo) second() {
	//f.firstJobDone <- struct{}{}
	//printSecond()
	//<-f.secondJobDone
	<-f.firstJobDone
	printSecond()
	f.secondJobDone <- struct{}{}
}

func (f *Foo) third() {
	//f.secondJobDone <- struct{}{}
	//printThird()
	<-f.secondJobDone
	printThird()
}

func printFirst() {
	print("first")
}

func printSecond() {
	print("second")
}

func printThird() {
	print("third")
}

func main() {
	f := NewFoo()
	go f.third()
	go f.second()
	go f.first()
	time.Sleep(time.Minute)
}
