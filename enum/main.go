package main

import "fmt"

// ENUMS

type pokemon int

const (
	Charizard  pokemon = 0
	Blastoise  pokemon = 1
	wigglypuff pokemon = 2
)

// or with iota which bascially increments everything below

const (
	marie pokemon = iota
	meena
	sona
	charchar
)

// Attaching functions to enums
func (pok pokemon) String() string { // String function automatically sends the name according to placeholder specifier
	switch pok {
	case Charizard:
		return "charizard"
	case Blastoise:
		return "blast"
	case wigglypuff:
		return "wiggly"
	default:
		panic("Shit")
	}
}

// Function demonstrating enums
func getHealth(pok pokemon) int {
	switch pok {
	case Charizard:
		return 100
	case Blastoise:
		return 200
	default:
		panic("Im too lazy") // stops execution
	}
}

func main() {
	fmt.Printf("Pokemon is %s, health is %d\n", Charizard.String(), getHealth(Charizard))
	// or
	fmt.Printf("Pokemon is %s, health is %d", Blastoise, getHealth((Blastoise)))
}
