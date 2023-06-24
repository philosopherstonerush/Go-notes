package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

// context is mainly used when you have to put restrictions on the duration/ or other factors.
// context can also be used to send data into a go-routine, then its guaranteed to be race free

func main() {
	start := time.Now()            // time at the moment is stored
	pokemon, err := fetchPokemon() // context with restrictions
	if err != nil {
		fmt.Printf("The error is %v\n", err)
	} else {
		fmt.Printf("The pokemon received is %v\n", pokemon)
	}
	timeDiff := time.Since(start) // the time that has been since start is calculated
	fmt.Printf("The time it took is %v\n", timeDiff)

	// Passing contexts
	ctx := context.WithValue(context.Background(), "Pok", "charizard") // parent context, key, value
	pok := getPokemon(ctx)
	fmt.Printf("The pokemon is %v\n", pok)
}

// context with restrictions
func fetchPokemon() (string, error) {
	// context with timeout, nobody wants an user to wait for 500 milliseconds for a result, this withTimeout() func takes 2 arguements and sends back two arguements, input argument 1 --> parent context, which for here we have used an empty context which is provided by Background() func and then the second arguement is the time. Cancel, happens if context is exited manually
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*100)
	defer cancel() // wait until it returns
	type pokeStruct struct {
		pokemon string
		err     error
	}
	pokemonch := make(chan pokeStruct, 1)
	go func() {
		pokemon, err := sendHttpPokemonRequest()
		if err != nil {
			log.Fatal(err)
			pokemonch <- pokeStruct{
				pokemon: "",
				err:     err,
			}
		} else {
			pokemonch <- pokeStruct{
				pokemon: pokemon,
				err:     err,
			}
		}
	}()

	select {
	// Done()
	// 1. Context time limit exceeded
	// 2. Cancel is invoked manually
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-pokemonch: // we got the result if within the time frame
		return res.pokemon, nil
	}
}

func sendHttpPokemonRequest() (string, error) {
	time.Sleep(time.Millisecond * 500)
	return "charizard", nil
}

// passing context
func getPokemon(ctx context.Context) any {
	var pok any = ctx.Value("Pok")
	return pok
}
