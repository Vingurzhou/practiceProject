package main

import (
	"fmt"
	"time"
)

type ZeroEvenOdd struct {
	n   int
	ch1 chan struct{}
	ch2 chan struct{}
	ch3 chan struct{}
}

func NewZeroEvenOdd(n int) *ZeroEvenOdd {
	ch1 := make(chan struct{})
	ch2 := make(chan struct{})
	ch3 := make(chan struct{})
	return &ZeroEvenOdd{n,
		ch1,
		ch2,
		ch3,
	}
}
func (zeo ZeroEvenOdd) zero() {
	for i := 1; i <= zeo.n; i++ {
		zeo.ch1 <- struct{}{}
		printNumber(0)
		if i%2 == 0 {
			<-zeo.ch2
		} else {
			<-zeo.ch3
		}
	}
}
func (zeo ZeroEvenOdd) even() {
	for i := 2; i <= zeo.n; i++ {
		if i%2 == 0 {
			zeo.ch2 <- struct{}{}
			printNumber(i)
			<-zeo.ch1
		}
	}
}
func (zeo ZeroEvenOdd) odd() {
	for i := 1; i <= zeo.n; i++ {
		if i%2 == 1 {
			zeo.ch3 <- struct{}{}
			printNumber(i)
			<-zeo.ch1
		}
	}
}
func printNumber(i int) {
	fmt.Println(i)
}
func main() {
	zeo := NewZeroEvenOdd(5)
	go zeo.zero()
	go zeo.even()
	go zeo.odd()
	<-zeo.ch1
	time.Sleep(time.Minute)
}
