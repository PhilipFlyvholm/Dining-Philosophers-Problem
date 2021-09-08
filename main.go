package main

import "fmt"

func main() {
	fmt.Println("hey")
	//var phil = makePhil(1, rightforkpointer, leftforkpointer)
	//fmt.Println("made a philosopher his id is:", phil.id)
	forks := make([]*Fork, 5)
	for i := 0; i < 5; i++ {
		var fork = makeFork()
		forkp := &fork
		forks[i] = forkp
	}
	philosophers := make([]Phil, 5)
	var phil Phil
	// link all philosophers together
	for i := 0; i < 5; i++ {
		phil = makePhil(i)
		phil.leftFork = forks[i]
		phil.rightFork = forks[(i+1)%5]
		philosophers[i] = phil
	}
}

func think() {

}

func eat() {

}

func makePhil(number int) Phil {
	phil := Phil{number, nil, nil, false, make(chan bool, 1), make(chan bool, 1)}
	return phil
}

func makeFork() Fork {
	fork := Fork{false, make(chan bool, 0), make(chan bool, 0)}
	return fork
}

type Fork struct {
	inUse   bool
	inputch chan bool
	output  chan bool
}

type Phil struct {
	id        int
	rightFork *Fork
	leftFork  *Fork
	eating    bool
	inputch   chan bool
	outputch  chan bool
}
