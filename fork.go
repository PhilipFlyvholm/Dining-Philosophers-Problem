package main

type Fork struct {
	id        int
	inUse     bool
	timesUsed int
	input     chan Request
	output    chan ForkState
}
type ForkState struct{
	inUse bool
	timesUsed int
}

func (f Fork) InnerLoop() {

	for {
		select {
		case stateRequest := <- f.input:
			stateRequest.requester.input <- ForkState{inUse: f.inUse,timesUsed: f.timesUsed}
		}
	}
}
