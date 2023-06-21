package main

import (
	"reflect"
	"testing"
)

// How to run tests?
// go mod init <projectname>
// go test ./... ---> baically runs every test file
// go test ./... -v ---> gives extra information while running
// go test ./ -run <test-function-name> ---> runs the specific test method

// Function that tests that other function in main.go, this must be prefixed by Test-
func TestReturnPokemon(t *testing.T) { // t *testing.T is boilerplate, go internal thingy
	expected := "chari"
	have := returnPokemon()

	if have != expected {
		t.Errorf("different pokemon returned") // Basically how the error is thrown
	}
}

// Testing structs

func TestPokemonStruct(t *testing.T) {
	expected := pokemon{
		name:   "Charizard",
		health: 100,
	}
	have := pokemon{
		name:   "chari",
		health: 50,
	}

	if !reflect.DeepEqual(expected, have) { // Checks if both struct objects are clones
		t.Errorf("Not equal")
	}
}
