package main

import (
	"fmt"
	"time"
)

var forks []Fork
var philosophers []Philosopher

func calcNumOfForks(numOfPhils int) int {
	if numOfPhils < 1 {
		return 0
	}
	if numOfPhils == 1 {
		return 2
	}
	return numOfPhils
}

func main() {
	fmt.Println("Welcome to the dining philosophers")
	var amountOfPhilosophers int = 5
	amountOfForks := calcNumOfForks(amountOfPhilosophers)
	fmt.Println("At the table we have", amountOfPhilosophers, "philosophers sharing", amountOfForks, "forks")
	forks = make([]Fork, amountOfForks)
	for i := 0; i < amountOfForks; i++ {
		forks[i] = NewFork(i)
		go forks[i].InnerLoop()
	}
	philosophers = make([]Philosopher, 5)
	for i := 0; i < amountOfPhilosophers; i++ {
		var leftForkID = i - 1
		if leftForkID < 0 {
			leftForkID = amountOfForks - 1
		}
		var rightForkID = i
		philosophers[i] = NewPhilosopher(i, &forks[leftForkID], &forks[rightForkID])

		go philosophers[i].InnerLoop(true)
	}
	fmt.Println("Everything is setup")

	printInfo(5)
}

func printInfo(seconds int) {
	for _, fork := range forks {
		fork.input <- Request{requestType: printState}
	}
	time.Sleep(time.Second * time.Duration(seconds))
	printInfo(seconds)
}

type Request struct {
	requester   *Philosopher
	requestType RequestType
}

type RequestType int

const (
	reserve    RequestType = iota // reserve = 0
	dismiss                       // dismiss = 1
	printState                    // printState = 2
)
