package main

import "fmt"

type brian int

func main() {
	var x brian
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
}
