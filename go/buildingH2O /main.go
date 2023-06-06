package main

import (
	"fmt"
	"time"
)

type H2O struct {
	chs1    []chan struct{}
	chs2    []chan struct{}
	num1    int
	num2    int
	barrier chan string
}

func NewH2O() *H2O {
	var chs1 []chan struct{}
	var chs2 []chan struct{}
	return &H2O{chs1: chs1,
		chs2: chs2,
	}
}

func (h *H2O) hydrogen(releaseHydrogen func()) {
	ch := make(chan struct{})
	h.chs1 = append(h.chs1, ch)
	ch <- struct{}{}

	releaseHydrogen()

}

func (h *H2O) oxygen(releaseOxygen func()) {
	ch := make(chan struct{})
	h.chs2 = append(h.chs2, ch)
	ch <- struct{}{}

	releaseOxygen()

}

func main() {
	water := "OOHHHH"
	h2o := NewH2O()
	for _, character := range water {
		switch character {
		case 'H':
			go h2o.hydrogen(func() { print("H") })
		case 'O':
			go h2o.oxygen(func() { print("O") })
		default:
			panic(character)
		}
	}

	for {
		if len(h2o.chs1) == 4 && len(h2o.chs2) == 2 {
			fmt.Println(len(h2o.chs1) == 4, len(h2o.chs2) == 2)
			break
		}
	}
	time.Sleep(time.Minute)
}
