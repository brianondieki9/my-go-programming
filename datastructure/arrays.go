package main

import "fmt"

func main() {
	var x [5]int
	x[4] = 56
	x[0] = 22
	x[2] = 555
	fmt.Println(x)
	fmt.Println(len(x))
}
