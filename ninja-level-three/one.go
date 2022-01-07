package main

import "fmt"

func main() {

	//first method is to do it this way
	//for i := 1; i <= 10000; i++ {
	//	fmt.Println(i)
	//}

	counter := 1

	for counter <= 10000 {
		fmt.Println(counter)
		counter++
	}
}
