package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Fork struct {
	id    int
	state ForkState
	input chan Request
	queue chan *Philosopher
}

func NewFork(_id int) Fork {
	return Fork{id: _id, input: make(chan Request, 2), queue: make(chan *Philosopher, 2)}
}

type ForkState struct {
	inUse     bool
	timesUsed int
}

func (f *Fork) InnerLoop() {
	rand.Seed(time.Now().UnixNano())
	for {
		select {
		case request := <-f.input:
			switch request.requestType {
			case reserve:
				//Reserve
				f.queue <- request.requester
				if !f.state.inUse {
					f.giveForkToNextPhilosopher()
				}
			case dismiss:
				//Dismiss
				if f.state.inUse {
					f.state.inUse = false
					time.Sleep(time.Second * time.Duration(rand.Intn(4)+3))
					if len(f.queue) >= 1 && !f.state.inUse {
						f.giveForkToNextPhilosopher()
					}
				}
			case printState:
				fmt.Println("Fork status, id:", f.id, "times used:", f.state.timesUsed, "in use:", f.state.inUse)
			}
		}
	}
}

func (f *Fork) giveForkToNextPhilosopher() {
	next := <-f.queue
	if next == nil {
		return
	}
	f.state.inUse = true
	f.state.timesUsed++
	next.input <- f
}
