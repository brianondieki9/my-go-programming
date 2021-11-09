package main

import "fmt"

var x int
var y string
var z bool

func main() {

	fmt.Println("Value for int: ", x)
	fmt.Println("Value for String: ", y)
	fmt.Println("Value for boolean: ", z)

	fmt.Printf("%v %v %v \n", x, y, z)
}
