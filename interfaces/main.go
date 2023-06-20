package main

import "fmt"

// Interfaces define the behaviour of the underlying implementations. So no matter what you know that they behave in a way that you want it to behave

// interfaces are always named -(er)
type PokemonStorer interface {
	getAllPokemon() ([]string, error) // abstract function that takes no arguments but returns two values of slice of string values and error
	putPokemon(string) error          // abstract function that takes in a pokemon name and then returns a error status code
}

// Implementing the interface (underlying implementation)
type MongoDBpokemonStorer struct {
	// values
}

// attaching the interface abstract function definitions
func (m MongoDBpokemonStorer) getAllPokemon() ([]string, error) {
	return []string{"char", "blast"}, nil // nil is no error
}

func (m MongoDBpokemonStorer) putPokemon(pok string) error {
	fmt.Println("Inserted")
	return nil
}

// Something that user the interface
type pokemonstorage struct {
	pokemonstore PokemonStorer
}

func main() {
	// Initializing the upper struct
	pks := pokemonstorage{
		pokemonstore: MongoDBpokemonStorer{},
	}

	// so now you can say

	fmt.Println(pks.pokemonstore.getAllPokemon())
	err := pks.pokemonstore.putPokemon("Charizard")
	if err == nil {
		fmt.Println("Hi")
	} else {
		fmt.Println("damn problem")
	}

	// The main advantage of interfaces is that, now you can swap out the mongodb implementation of the underlying struct then use like postgress without bothering the upper architecture code. As we know whatever code that lies below, it always has the two abstract functions defined

}
