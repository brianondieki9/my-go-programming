package main

import "fmt"

func main() {

	favSport := "soccer"

	switch favSport {
	case "basketball":
		fmt.Println("You are a good basketball player")
	case "soccer":
		fmt.Println("You are a good soccer player")
	case "football":
		fmt.Println("You are a good football player")
	default:
		fmt.Println("You do not like any sport")
	}
}
