package main

import "fmt"

func main() {
	//loop the decimal numbers for all upper case
	counter := 1
	for i := 65; i <= 90; i++ {
		//loop and print three times
		fmt.Println(counter)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t%#U \n", i)
		}
		counter++
	}
}
