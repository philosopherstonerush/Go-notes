package main

import "fmt"

type pokemon struct {
	name   string
	health int
}

func (pok pokemon) getName() string {
	return pok.name
}

// pokemon struct got inherited by this bigger struct
type megapokemon struct {
	pokemon
	megamove string
}

func main() {
	fmt.Print("tis nothin but a scratch\n")
	m := megapokemon{
		pokemon: pokemon{
			name:   "charizard",
			health: 200,
		},
		megamove: "blizzard",
	}
	fmt.Printf("struct inherited and initializaed %+v\n", m)
	fmt.Printf("This is the function that got inherited, name of the pokemon is %v", m.getName())
}
