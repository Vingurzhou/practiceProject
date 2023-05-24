package main

import (
	"fmt"
	"time"
)

type DiningPhilosophers struct {
}

func NewDiningPhilosophers() *DiningPhilosophers {
	return &DiningPhilosophers{}
}

func (dp *DiningPhilosophers) wantsToEat(philosopher int, pickLeftFork func(), pickRightFork func(), eat func(), putLeftFork func(), putRightFork func()) {

}
func main() {
	n := 100
	dp := NewDiningPhilosophers()
	for i := 0; i < n; i++ {
		go dp.wantsToEat(0, func() { fmt.Println("拿起左边的叉子") }, func() { fmt.Println("拿起右边的叉子") }, func() { fmt.Println("吃面") }, func() { fmt.Println("放下左边的叉子") }, func() { fmt.Println("放下右边的叉子") })
		go dp.wantsToEat(1, func() { fmt.Println("拿起左边的叉子") }, func() { fmt.Println("拿起右边的叉子") }, func() { fmt.Println("吃面") }, func() { fmt.Println("放下左边的叉子") }, func() { fmt.Println("放下右边的叉子") })
		go dp.wantsToEat(2, func() { fmt.Println("拿起左边的叉子") }, func() { fmt.Println("拿起右边的叉子") }, func() { fmt.Println("吃面") }, func() { fmt.Println("放下左边的叉子") }, func() { fmt.Println("放下右边的叉子") })
		go dp.wantsToEat(3, func() { fmt.Println("拿起左边的叉子") }, func() { fmt.Println("拿起右边的叉子") }, func() { fmt.Println("吃面") }, func() { fmt.Println("放下左边的叉子") }, func() { fmt.Println("放下右边的叉子") })
		go dp.wantsToEat(4, func() { fmt.Println("拿起左边的叉子") }, func() { fmt.Println("拿起右边的叉子") }, func() { fmt.Println("吃面") }, func() { fmt.Println("放下左边的叉子") }, func() { fmt.Println("放下右边的叉子") })
	}
	time.Sleep(time.Minute)
}
