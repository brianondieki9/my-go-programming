package main

import "fmt"

var currentState bool

func main() {

	//print the current assigned default value
	fmt.Println(currentState)

	//assign the value to the variable
	currentState = true

	//print the new value
	fmt.Println(currentState)

	//print the type of the value
	fmt.Printf("%T\n", currentState)

	a := 43
	b := 43
	fmt.Println(a == b)

}
