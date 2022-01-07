package main

import "fmt"

type brian int

var y int

func main() {

//create a variable of type Brian
	var x brian

	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)

	//initialize variable x to 42
	x = 42
	fmt.Println(x)

	//Now let us convert the variable conversion
	y = int(x)
	fmt.Println(y)
}
