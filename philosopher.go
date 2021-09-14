package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Philosopher struct {
	id        int
	leftFork  ForkReference
	rightFork ForkReference
	state     PhilosopherState
	input     chan *Fork
	//output    chan bool
}

func NewPhilosopher(_id int, _leftFork *Fork, _rightFork *Fork) Philosopher {
	return Philosopher{
		id:        _id,
		leftFork:  ForkReference{fork: _leftFork},
		rightFork: ForkReference{fork: _rightFork},
		input:     make(chan *Fork, 2),
	}
}

type PhilosopherState struct {
	eating     bool
	timesEaten int
}

type ForkReference struct {
	fork   *Fork
	inHand bool
}

func (p *Philosopher) InnerLoop(startWithReserve bool) {
	rand.Seed(time.Now().UnixNano())
	if startWithReserve {
		p.reserveForks()
	}
	for {
		timeout := make(chan bool, 1)
		go func() {
			time.Sleep(time.Second * 7)
			timeout <- true
		}()
		select {
		case fork := <-p.input:
			var id = fork.id
			if p.leftFork.fork.id == id {
				p.leftFork.inHand = true
			} else if p.rightFork.fork.id == id {
				p.rightFork.inHand = true
			}
			if p.isBothForksInHand() {
				p.eat()
			}
		case <-timeout:
			p.think()
			//fmt.Println("Timeout Happened", p.id)
		}
	}
}

func (p *Philosopher) deselectForks() {
	p.leftFork.inHand = false
	p.rightFork.inHand = false
}

func (p *Philosopher) reserveForks() {
	if p.state.eating {
		return
	}
	if !p.leftFork.inHand {
		p.leftFork.fork.input <- Request{
			requester:   p,
			requestType: reserve,
		}
	}
	if !p.rightFork.inHand {
		p.rightFork.fork.input <- Request{
			requester:   p,
			requestType: reserve,
		}
	}
}

func (p *Philosopher) eat() {
	if p.state.eating {
		return
	}
	fmt.Println("Philosopher", p.id, "starting to eat")
	p.state.timesEaten++
	p.state.eating = true
	time.Sleep(time.Second * time.Duration(rand.Intn(4)+3)) //Sleep for between 1 to 3 seconds
	p.think()

}
func (p *Philosopher) think() {
	if !p.state.eating {
		return
	}
	fmt.Println("Philosopher", p.id, "starting to think")
	p.state.eating = false
	p.deselectForks()
	p.leftFork.fork.input <- Request{
		requester:   p,
		requestType: dismiss,
	}
	p.rightFork.fork.input <- Request{
		requester:   p,
		requestType: dismiss,
	}
	time.Sleep(time.Second)
	p.reserveForks()
}

func (p *Philosopher) isBothForksInHand() bool {
	return (p.leftFork.inHand && p.rightFork.inHand)
}
