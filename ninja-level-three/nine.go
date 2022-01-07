package main

import "fmt"

func main() {
	//Empt switch
	switch {
	case false:
		fmt.Println("should not print")
	case true:
		fmt.Println("shoud print")
	}
}
