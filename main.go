package main

import (
	"fmt"
)

func main() {
	var forks = make([]Fork, 5)
	for i := 0; i < 5; i++ {
		forks[i] = Fork{id: i}
		go forks[i].InnerLoop()
	}
	var philosophers = make([]Philosopher, 5)
	for i := 0; i < 5; i++ {
		var leftForkID = i - 1
		if leftForkID < 0 {
			leftForkID = 5
		}
		var rightForkID = i
		philosophers[i] = Philosopher{id: i, leftFork: &forks[leftForkID], rightFork: &forks[rightForkID]}
	}

	fmt.Println("Welcome to the dining philosophers")
}

type Request struct {
	requester *Philosopher
}
