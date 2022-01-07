package main

import "fmt"

func main() {
	x := 42
	fmt.Printf("%d\n", x)
	fmt.Printf("%b\n", x)
	fmt.Printf("%#X\n", x)

	y := 42 << 1
	fmt.Printf("%d\n", y)
	fmt.Printf("%b\n", y)
	fmt.Printf("%#X\n", y)

}
