package main

import "fmt"

func main() {
	friends := []string{"ramesh", "Bhopal", "Dumbass"}

	// for
	for i := 0; i < len(friends); i++ {
		fmt.Println(friends[i])
	}

	// for loop with range
	for i, v := range friends {
		fmt.Printf("index %v, Value is %v\n", i, v)
	}

	// for loop but omiting index, it can be used to omit value
	for _, v := range friends {
		fmt.Println(v)
	}
}
