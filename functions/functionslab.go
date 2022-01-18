package main

import "fmt"

func main() {
	foo()
	bar("John")
	s1 := woo("peter")
	fmt.Println(s1)

	x, y := mouse("Jane", "Wanjiru")
	fmt.Println(x)
	fmt.Println(y)
}

//func (r receiver) identifier(parameters) (return(s)){....code....}

func foo() {
	fmt.Println("Hello from foo")
}

//everything in go is passed by VALUE
func bar(s string) {
	fmt.Println("Hello", s)
}

func woo(s string) string {
	return fmt.Sprint("Hello from Woo: ", s)
}

func mouse(firstname string, lastname string) (string, bool) {
	a := fmt.Sprint(firstname, " ", lastname, `, says "Hello"`)
	b := true

	return a, b
}
