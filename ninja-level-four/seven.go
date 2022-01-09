package main

import "fmt"

func main() {

	//create slices of type string to store information for a single person
	p1 := []string{"James", "Bond", "Shaken, not stirred"}
	p2 := []string{"Miss", "Moneypenny", "Helloooooo, James."}

	//print the slices seperately
	fmt.Println(p1, p2)

	//create a multi dimensional slice to store all the slices

	persons := [][]string{p1, p2}

	fmt.Println(persons)

	for i, v := range persons {
		fmt.Println(i)
		for j, q := range v {
			fmt.Println(j, q)
		}
	}
	fmt.Println("------------------------------------------------------")
	fmt.Println("Looping in multidimensional array without using range")

	for i := 0; i < len(persons); i++ {
		for j := 0; j < len(persons[i]); j++ {
			fmt.Println(persons[i][j])
		}
		fmt.Println()
	}
}
