package main

type Philosopher struct {
	id         int
	leftFork   *Fork
	rightFork  *Fork
	timesEaten int
	input      chan ForkState
	output     chan bool
}
