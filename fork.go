package main

type Fork struct {
	inUse     bool
	timesUsed int
	inputch   chan string
	outputch  chan string
}
