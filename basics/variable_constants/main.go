package main

import "fmt"

// used to print out onto the console

// Global scope
var firstname = "john"      // type infered
var lastname string = "doe" // type explicit

// instead of repeating var multiple times, you can say

var (
	firstborn  string = "ramesh"
	secondborn        = "bhopal"
)

// constants are lowercase and declared in the global scope

const maincharacterage = 45
const (
	firstbornage  = 20
	secondbornage = 18
)

// default values, if just initialized without assigment
var babyone string
var babyoneage int

func main() {
	// local scope
	// you can use := to tell the compiler to infer it. This cannot be used in the global scope
	var secondcharater string = "vipul"
	var secondcharacterfirstborn = "mygod"
	secondcharacterage := 18

	fmt.Println("Local variables " + secondcharacterfirstborn + secondcharater)
	fmt.Println(secondcharacterage)

	fmt.Println("default" + babyone)
	fmt.Println("default int ")
	fmt.Println(babyoneage)
}
