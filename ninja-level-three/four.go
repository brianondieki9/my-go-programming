package main

import "fmt"

func main() {

	for i := 10; i <= 100; i++ {
		remainder := i % 4

		if remainder == 0 {
			continue
		}

		fmt.Printf("The remainder of %v is %v \n", i, remainder)
	}
}
