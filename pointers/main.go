package main

import "fmt"

// Pointers

type pokemon struct {
	name   string
	health int
}

func (pok *pokemon) TakeDamage(damage int) {
	pok.health -= damage
}

func main() {
	p := pokemon{
		name:   "Blastoise",
		health: 100,
	}
	fmt.Printf("Before damage %+v", p)
	p.TakeDamage(10)
	fmt.Printf("After damage %+v", p)
}
