package main

import "fmt"

type brian int

var y int

func main() {
	var x brian
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)

	x = 42
	fmt.Println(x)

	//conversion
	y = int(x)
	fmt.Println(y)
}
