package main

import "fmt"

// generic function
func printWhatever[T any](t T) { // any is basically any datatype
	fmt.Println(t)
}

// generic struct
type genStruct[T any, B any] struct { // just use "comparable" when you are going to use it for map
	index T
	value B
}

func (g genStruct[T, B]) getValue() B {
	return g.value
}

func main() {
	// generic function
	printWhatever[string]("Charizard")

	// generic struct
	p := genStruct[string, string]{
		index: "Ramesh",
		value: "24",
	}
	fmt.Println(p.value)
}
