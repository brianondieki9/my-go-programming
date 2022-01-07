package main

import "fmt"

func main() {

	birthyear := 1987
	for {
		if birthyear > 2022 {
			break
		}
		fmt.Println(birthyear)
		birthyear++
	}
}
