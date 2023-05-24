package main

import (
	"fmt"
	"time"
)

type FizzBuzz struct {
	n   int
	ch1 chan struct{}
	ch2 chan struct{}
	ch3 chan struct{}
	ch4 chan struct{}
}

func NewFizzBuzz(n int) *FizzBuzz {
	ch1 := make(chan struct{})
	ch2 := make(chan struct{})
	ch3 := make(chan struct{})
	ch4 := make(chan struct{})
	return &FizzBuzz{n: n,
		ch1: ch1,
		ch2: ch2,
		ch3: ch3,
		ch4: ch4,
	}
}

func (fb *FizzBuzz) fizz(printFizz func()) {
	for i := 1; i <= fb.n; i++ {
		<-fb.ch1
		printFizz()
	}
}
func (fb *FizzBuzz) buzz(printBuzz func()) {
	for i := 1; i <= fb.n; i++ {
		<-fb.ch2
		printBuzz()
	}
}
func (fb *FizzBuzz) fizzbuzz(printFizzBuzz func()) {
	for i := 1; i <= fb.n; i++ {
		<-fb.ch3
		printFizzBuzz()
	}
}
func (fb *FizzBuzz) number(printNumber func(int)) {
	for i := 1; i <= fb.n; i++ {
		<-fb.ch4
		printNumber(i)
	}
}
func main() {
	n := 15
	fb := NewFizzBuzz(n)
	go fb.fizz(func() { print("fizz ") })
	go fb.buzz(func() { print("buzz ") })
	go fb.fizzbuzz(func() { print("fizzbuzz ") })
	go fb.number(func(x int) { fmt.Printf("%d ", x) })
	for i := 1; i <= fb.n; i++ {
		switch {
		case i%3 == 0 && i%5 != 0:
			fb.ch1 <- struct{}{}
		case i%3 != 0 && i%5 == 0:
			fb.ch2 <- struct{}{}
		case i%3 == 0 && i%5 == 0:
			fb.ch3 <- struct{}{}
		case i%3 != 0 && i%5 != 0:
			fb.ch4 <- struct{}{}
		default:
			panic(i)
		}
	}
	time.Sleep(time.Minute)
}
