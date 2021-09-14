package main

import (
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
				//fmt.Println("Fork reserving", f.id)
				f.queue <- request.requester
				if !f.state.inUse {
					f.giveForkToNextPhilosopher()
				}
			case dismiss:
				//Dismiss
				//fmt.Println("Fork dismissing", f.id)
				if f.state.inUse {
					//fmt.Println("Fork dismissed", f.id)
					f.state.inUse = false
					time.Sleep(time.Second * time.Duration(rand.Intn(4)+3)) //Sleep for between 1 to 3 seconds
					if len(f.queue) >= 1 && !f.state.inUse {
						f.giveForkToNextPhilosopher()
					}
				}
			case printState:
				//fmt.Println(fmt.Sprintf("Fork status for id %d: Is in use: %t, Has been used: %d times", f.id, f.state.inUse, f.state.timesUsed))
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
	//fmt.Println("Fork times used:", f.state.timesUsed)
	next.input <- f
}
