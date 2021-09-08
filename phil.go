package main

type Phil struct {
	rightFork *Fork
	leftFork  *Fork
	eating    bool
	inputch   chan string
	outputch  chan string
}
