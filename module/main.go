package main

import (
	"fmt"
	"mypack/pack2" // module name is mypack, inside it we have the folder pack2
)

// Every file in the same directory as the main.go file is considered to be of the same package
func main() {
	fmt.Println("This is the mainpackage")
	fmt.Println(getGreeting()) // from pack1.go doesnt need to be imported since its in the same directory but compiler doesnt know the code exists in a separate file. Run "go mod init <name>" to create a module file that tracks it.
	// You can then compile it using "go build -o <name of the output file>" -o is output, make sure the name ends with .exe
	// Making use of folders, see pack2
	fmt.Println(pack2.GetName())
	// Make sure to build again!
}
