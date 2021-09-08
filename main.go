package main

func main() {
	makeForks()
	makePhils()

}


func makeForks() {
	var fork1 := Fork{make(chan string, 0), make(chan string, 0)}
	var fork2 := Fork{make(chan string, 0), make(chan string, 0)}
	var fork3 := Fork{make(chan string, 0), make(chan string, 0)}
	var fork4 := Fork{make(chan string, 0), make(chan string, 0)}
	var fork5 := Fork{make(chan string, 0), make(chan string, 0)}
}
func makePhils() {
	var phil1 := Phil{&fork1, &fork2, make(chan string, 0), make(chan string, 0)}
	var phil2 := Phil{&fork2, &fork3, make(chan string, 0), make(chan string, 0)}
	var phil3 := Phil{&fork3, &fork4, make(chan string, 0), make(chan string, 0)}
	var phil4 := Phil{&fork4, &fork5, make(chan string, 0), make(chan string, 0)}
	var phil5 := Phil{&fork5, &fork1, make(chan string, 0), make(chan string, 0)}
}