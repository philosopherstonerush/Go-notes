// Bringing all the concepts together
// How to fetch different parts of user data at the same?

package main

import (
	"fmt"
	"sync"
	"time"
)

type Pokemon struct {
	name   string
	health int
	moves  []string
}

func main() {
	start := time.Now()
	pok := fetchDataSynchronously()
	fmt.Printf("Pokemon is %v\n", pok)
	timeDiff := time.Since(start)
	fmt.Println("The time it took for synchronous fetching is ", timeDiff)

	// Async --> optimized to the lowest time taking function
	start = time.Now()
	pok = fetchDataAsynchronously()
	fmt.Printf("Pokemon is %v\n", pok)
	timeDiff = time.Since(start)
	fmt.Println("The time it took for Asynchronous fetching is ", timeDiff)
}

func fetchDataSynchronously() *Pokemon {
	pok := &Pokemon{}
	pok.name = getPokemonName()
	pok.health = getPokemonHealth()
	pok.moves = getPokemonMoves()
	return pok
}

type response struct {
	identity string
	data     any
	err      error
}

func fetchDataAsynchronously() *Pokemon {
	wg := &sync.WaitGroup{}          // need to wait for three go routines to finish
	pok := &Pokemon{}                // empty return struct
	respch := make(chan response, 3) // since we are writing to the buffer three times, 3 is the buffer or the number of available boxes
	go getPokemonHealthA(respch, wg) // passing the respch and wg pointer
	go getPokemonMovesA(respch, wg)
	go getPokemonNameA(respch, wg)
	wg.Add(3)     // because three go routines
	wg.Wait()     // wait till all the go routines finish
	close(respch) // close the response channel or else the below range loop will execute indefinitely
	for res := range respch {
		switch res.identity {
		case "name":
			pok.name = res.data.(string) // assert type from any to the specific type
		case "moves":
			pok.moves = res.data.([]string)
		case "health":
			pok.health = res.data.(int)
		}

		// or
		// switch res.data.(type) {
		// case string: executes if its a string type of data
		// blah blah
		// }
	}
	return pok
}

func getPokemonName() string {
	time.Sleep(time.Millisecond * 500)
	return "charizard"
}

func getPokemonHealth() int {
	time.Sleep(time.Millisecond * 200)
	return 10
}

func getPokemonMoves() []string {
	time.Sleep(time.Millisecond * 300)
	return []string{"Flamethrower", "willo-wisp", "fly", "stomp"}
}

func getPokemonNameA(respch chan response, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 500)
	respch <- response{
		identity: "name",
		data:     "charizard",
		err:      nil,
	}
	wg.Done() // to say its finished
}

func getPokemonHealthA(respch chan response, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 200)
	respch <- response{
		identity: "health",
		data:     10,
		err:      nil,
	}
	wg.Done()
}

func getPokemonMovesA(respch chan response, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 300)
	respch <- response{
		identity: "moves",
		data:     []string{"Flamethrower", "willo-wisp", "fly", "stomp"},
		err:      nil,
	}
	wg.Done()
}
