package main

import "fmt"

func main() {

	total := sum(2, 3, 4, 5, 6, 7, 8)
	fmt.Println("The total of the sum is: ", total)
}

// func (r receiver) identifier(parameter(s)) (return(s)) {code}
func sum(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0

	for _, v := range x {
		sum += v
	}
	return sum
}
