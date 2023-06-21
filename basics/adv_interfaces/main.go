package main

import "fmt"

// Advanced interfaces

type pokemongetter interface {
	getAll() ([]string, error)
}

// This interface inherits the definition seen in pokemon getter
type pokemonstorer interface {
	pokemongetter
	put(string) error
}

// upper level struct that uses the interface
type pokemon struct {
	pokemonstore pokemonstorer
	name         string
}

// lowerlevel struct
type pokemonDB struct {
	// values
}

func (pok pokemonDB) getAll() ([]string, error) {
	return []string{"Charizard, blastoise, squirtle"}, nil
}

func (pok pokemonDB) put(s string) error {
	fmt.Printf("Pokemon %v was inserted into the DB\n", s)
	return nil
}

// Passing functions as arguements
type name func() (s string)

func fireType() string {
	return "charizard"
}

func getName(obj int, fn name) string {
	if obj == 1 {
		return fn()
	}
	return ""
}

func main() {
	p := pokemon{
		pokemonstore: pokemonDB{},
		name:         "charchar",
	}

	fmt.Println(p.pokemonstore.getAll())
	fmt.Println(p.pokemonstore.put("Charizard"))

	// Passing function as arguments
	fmt.Printf("Aight here we go again, %v", getName(1, fireType))
}
