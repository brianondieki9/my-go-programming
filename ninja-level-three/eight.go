package main

import "fmt"

func main() {
	score := 89

	if score < 50 {
		fmt.Println("Your grade is D")
	} else if score > 50 && score < 60 {
		fmt.Println("Your grrade is C")
	} else if score > 60 && score < 80 {
		fmt.Println("Your grade is B")
	} else {
		fmt.Println("Your grade is A")
	}
}
