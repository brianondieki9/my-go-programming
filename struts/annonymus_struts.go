package main

import "fmt"

func main() {

	p1 := struct {
		first string
		last  string
		age   int
	}{
		first: "brian",
		last:  "Mosoti",
		age:   23,
	}

	fmt.Println(p1)
	fmt.Println(p1.first, p1.last, p1.age)
}
