package main

import (
	"fmt"
	"math/rand"
	"time"
)

type ForkStatus int

const (
	Available ForkStatus = iota // Available = 0
	InUse                       // InUse = 1
)

type Fork struct {
	id     int
	status ForkStatus
	state  ForkState
	input  chan Request
	queue  chan *Philosopher
}

func NewFork(_id int) Fork {
	return Fork{id: _id, input: make(chan Request, 5), queue: make(chan *Philosopher, 2)}
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
				fmt.Println("Reserving", f.id)
				f.queue <- request.requester
				if f.status == Available && !f.state.inUse {
					f.giveForkToNextPhilosopher()
				}
			case dismiss:
				//Dismiss
				fmt.Println("Dismissing", f.id)
				if f.status == InUse {
					fmt.Println("Dismissed", f.id)
					f.state.inUse = false
					f.status = Available
					time.Sleep(time.Second * time.Duration(rand.Intn(3)+1)) //Sleep for between 1 to 3 seconds
					if len(f.queue) > 1 && f.status == Available && !f.state.inUse {
						f.giveForkToNextPhilosopher()
					}
				}
			case printState:
				fmt.Println(fmt.Sprintf("Fork status for id %d: Is in use: %t, Has been used: %d times", f.id, f.state.inUse, f.state.timesUsed))

			}
		}
	}
}

func (f *Fork) giveForkToNextPhilosopher() {
	if f.status != Available {
		return
	}
	next := <-f.queue
	if next == nil {
		return
	}
	f.status = InUse
	f.state.inUse = true
	f.state.timesUsed++
	fmt.Println("Times used:", f.state.timesUsed)
	next.input <- f
}
