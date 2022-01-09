package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 55, 5, 6, 6, 3, 1}

	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T", x)
}
