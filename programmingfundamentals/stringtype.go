package main

import "fmt"

func main() {

	s := "Hello, playground"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	//printing the decimal value of the characters
	fmt.Println("printing the decimal value of the characters")
	bs := ([]byte(s))
	fmt.Println(bs)

	//printing the type of the array(byte is similar to uint8)
	fmt.Printf("%T\n", bs)

	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}

	fmt.Println(" ")
	//printing the utf8 value of the characters

	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U\n", s[i])
	}

	//printing the hexademical
}
