package main

import "fmt"

var x int
var y string
var z bool

func main() {

	x = 42
	y = "James Bond"
	z = true

	//fmt.Println("Value for int: ", x)
	//fmt.Println("Value for String: ", y)
	//fmt.Println("Value for boolean: ", z)

	s := fmt.Sprintf("%v %v %v \n", x, y, z)
	//fmt.Printf("%v %v %v \n", x, y, z)
	fmt.Println(s)

}
