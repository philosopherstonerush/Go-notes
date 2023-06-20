package main

import (
	"fmt"
)

// More datatypes
var (
	floatVar32 float32 = 1.0
	floatVar64 float64 = 1.00 // Always use float64
	name       string  = "Foo"
	character  rune    = 'f' // Double quotes define a string
	intVar32   int32   = 1   // Can be positive or negative
	intVar64   int64   = 2
	intVar     int     = 1
	uintVar    uint    = 1 // Can only be positive
	uintVar32  uint32  = 1
	uintVar64  uint64  = 23
	uintVar8   uint8   = 0x1 // Same as byte
	byteVar    byte    = 0x2
)

// Structs
type Pokemon struct {
	name   string
	health int
	moves  int
}

// attaching function to the struct
func (pok Pokemon) getNameA() string {
	return pok.name
}

// A function that takes the struct and returns the name of the pokemon, that is non attached
func getName(pok Pokemon) string { // the function takes the argument name first then the datatype
	return pok.name
}

// Custom types
type Weapon string // Now you have a datatype called Weapon that is a String

func main() {
	//Initializing struct
	pok := Pokemon{
		name:   "Charizard",
		health: 100,
		moves:  4,
	}

	fmt.Printf("The pokemon is: %v\n", pok)              // verbose print, v puts whatever you send into the string
	fmt.Printf("The pokemon is: %+v\n", pok)             // Another way to print the structs but this time you get the name of the struct variables
	fmt.Printf("The health is: %d\n", pok.health)        // Accessing the variables of your struct
	fmt.Printf("Getting the name: %v\n", getName(pok))   // Non-attached function
	fmt.Printf("Getting the name: %v\n", pok.getNameA()) // Attached

	// Maps
	gems := map[string]int64{}      // basically like a dictionary with string key and int64 value
	gemsM := make(map[string]int64) // this is also right

	// Initializing, you can specify the values within the brackets or pass like this below
	gems["Diamond"] = 10000
	gems["ruby"] = 5000
	gemsM["ass"] = 100
	// Checks if the gems key value is present
	age, ok := gems["Diamond"]
	if !ok {
		fmt.Println("Diamond is not in the map")
	} else {
		fmt.Printf("Age is : %d\n", age)
	}
	// deleting map values
	delete(gems, "Diamond")
	// iterating over the values
	for k, v := range gems {
		// you can have just k there for keys only as well
		fmt.Printf("The key is %v and the value is %v\n", k, v)
	}

	// slices - they can shrink or expand, much like lists!
	friends := []string{"Ramesh", "balu"}
	Otherfriends := make([]string, 3) // You have to specify how many values you are going to store initially

	for k, v := range friends {
		fmt.Println(k, v)
	}
	for k, v := range Otherfriends {
		fmt.Println(k, v)
	}
	// appending
	friends = append(friends, "Bhopal")

	// Arrays - Fixed size, they dont shrink or increase
	girlfriends := [2]string{"Mia", "Reema"}

}
