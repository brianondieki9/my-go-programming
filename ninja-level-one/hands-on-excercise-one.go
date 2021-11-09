package main

import "fmt"

func main() {
	x := 42
	y := "James Bond"
	z := true

	//print multipe statements
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	//print using a single statement

	fmt.Println(x, y, z)

	fmt.Printf("%v, %v, %v\n", x, y, z)

}
