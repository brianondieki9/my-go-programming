package main

import (
	"fmt"
)

func main() {
	x := 43
	y := 34.33

	fmt.Println(x)
	fmt.Println(y)

	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)

	var a int = 343
	var b float64 = 34.33

	fmt.Println(a)
	fmt.Println(b)

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)

	testVariable := 345
	fmt.Println(testVariable)

	n, err := fmt.Printf("%T", testVariable)
	fmt.Println(n)
	fmt.Println(err)

	fmt.Println("base 2")
	fmt.Printf("%b\n", testVariable)

	fmt.Println("the character represented by the corresponding Unicode code point")
	fmt.Printf("%c\n", testVariable)

	fmt.Println("base 10")
	fmt.Printf("%d\n", testVariable)

	fmt.Println("base 8")
	fmt.Printf("%o\n", testVariable)

	fmt.Println("base 8 with 0o prefix")
	fmt.Printf("%O\n", testVariable)

	fmt.Println("a single-quoted character literal safely escaped with Go syntax")
	fmt.Printf("%q\n", testVariable)

	fmt.Println("base 16, with lower-case letters for a-f")
	fmt.Printf("%x\n", testVariable)

	fmt.Println("base 16, with upper-case letters for A-F")
	fmt.Printf("%X\n", testVariable)

	fmt.Println("Unicode format: U+1234; same as")
	fmt.Printf("%U\n", testVariable)

}
