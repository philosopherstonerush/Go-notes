package main

import (
	"fmt"
	"time"
)

// To make a go routine, you just specify "go" keyword

func main() {
	// Making an async goroutine call to the getPokemon function
	go getPokemon()
	// or to making anonymous function calls
	go func() {
		result := getPokemon()
		fmt.Println(result)
	}()
	// the above code wont work, to get results --> channel
	// variables that are defined as channel variables must be named like result(ch)
	resultch := make(chan string) // could be anything, float64 etc
	// channel is of two types, 1) buffered channel 2) unbuffered channel
	// resultch is an unbuffered channel
	// a channel in golang will block if its full
	resultbuffch := make(chan string, 2) // buffered channel
	// If the channel gets to 2 or exceeds then, it will block or deadlock occurs
	// buffering is basically the number of boxes where you can put stuff into
	// If production is higher that is, you put too much into the box, more than you consume then its a deadlock and vice versa
	resultbuffch <- getPokemon() // <- is used to put stuff into the buffer
	resultbuffch <- getPokemon()
	close(resultbuffch) // we close the channel after we finish putting stuff into it
	close(resultch)
	for { // for loop that executes always
		result, ok := <-resultbuffch // to take whatever that is stored in the buffer
		if !ok {
			fmt.Println("Done printing whatever was inside the channel")
			break
		} else {
			fmt.Println(result)
		}
	}
}

func getPokemon() string {
	time.Sleep(time.Second * 2) // wait for 2 seconds
	return "blastoise"
}
