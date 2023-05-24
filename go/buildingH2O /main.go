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
	for {
		if len(h.chs2) != 2 {
			continue
		} else {
			<-h.chs2[0]
			h.chs2 = h.chs2[1:]
			break
		}
	}
}

func (h *H2O) oxygen(releaseOxygen func()) {
	ch := make(chan struct{})
	h.chs2 = append(h.chs2, ch)
	ch <- struct{}{}

	releaseOxygen()
	for {
		if len(h.chs1) != 3 {
			continue
		} else {
			<-h.chs1[0]
			h.chs1 = h.chs1[1:]
			break
		}
	}
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
		fmt.Println((h2o.chs1))
		if len(h2o.chs1) != 4 {
			continue
		} else {
			<-h2o.chs1[0]
			h2o.chs1 = h2o.chs1[1:]
			break
		}
	}

	time.Sleep(time.Minute)
}
