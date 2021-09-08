package main

import (
	"fmt"
)

type PhilInputMessage struct {
	isInUse   bool
	rightFork bool
}

type ForkInputMessage struct {
	requester *Phil
	rightFork bool
}

type Phil struct {
	id         string
	rightFork  *Fork
	leftFork   *Fork
	eating     bool
	timesEaten int
	input      chan *PhilInputMessage
	output     chan bool
}

type Fork struct {
	id        string
	inUse     bool
	timesUsed int
	input     chan *ForkInputMessage
	output    chan *Phil
}

func main() {
	fmt.Println("Welcome to the dining philosophers")

	var fork1 = Fork{id: "1", input: make(chan *Phil), output: make(chan *Phil)}
	var fork2 = Fork{id: "2", input: make(chan *Phil), output: make(chan *Phil)}

	var phil1 = Phil{id: "Philosopher", rightFork: &fork1, leftFork: &fork2}

	go phil1.looper()
	go fork1.looper()
	go fork2.looper()

	var input string
	fmt.Scanln(&input)

	//var fork2 = Fork{id: "2"}
	//var fork3 = Fork{id: "3"}
	//var fork4 = Fork{id: "4"}
	//var fork5 = Fork{id: "5"}

	//var phil1 = Phil{rightFork: fork1, leftFork: fork2}
	//var phil2 = Phil{rightFork: fork2, leftFork: fork3}
	//var phil3 = Phil{rightFork: fork3, leftFork: fork4}
	//var phil4 = Phil{rightFork: fork4, leftFork: fork5}
	//var phil5 = Phil{rightFork: fork5, leftFork: fork1}
}

func (f Fork) looper() {
	var isInUse bool

	for {
		select {
		case msg := <-f.input:
			//phil.input <- &PhilInputMessage{isInUse: f.inUse, rightFork: phil.rightFork}¨
			msg.requester.input <- &PhilInputMessage{isInUse: f.inUse, rightFork: Phil.rightFork}

		case f.output <- isInUse:
			fmt.Println("f.output <---- isInUse")
		}
	}
}

func (f Fork) isInUse() bool {
	return <-f.output
}

func (f Fork) setInUse(inUse bool) {
	f.input <- inUse
}

func (p Phil) Looper() {
	var isRightInUse bool
	var isLeftInUse bool

	for {
		select {
		case isInUse := <-p.input:

		}
	}
}

func (p Phil) askIfForkInUse(right Fork, left Fork) {
	right.input <- &ForkInputMessage{requester: &p, rightFork: true}
	left.input <- &ForkInputMessage{requester: &p, rightFork: false}
}
