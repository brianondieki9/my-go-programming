package main

import "fmt"

type person struct {
	firstname string
	lastname  string
}

type secretAgent struct {
	person
	licenseToKill bool
}

func (s secretAgent) speak() {
	fmt.Println("I am", s.firstname, s.lastname)
}

func (s secretAgent) countCharacters() {

	fmt.Println(len(s.firstname) + len(s.lastname))
}

// func (r receiver) identifier(parameters) (return(s)){code}
func main() {
	sa1 := secretAgent{
		person: person{
			firstname: "Brian",
			lastname:  "Ondieki",
		},
		licenseToKill: true}

	sa2 := secretAgent{
		person: person{
			firstname: "Peter",
			lastname:  "Mwaura",
		},
		licenseToKill: false}

	sa1.speak()
	sa2.speak()

	sa1.countCharacters()
	sa2.countCharacters()

	fmt.Println(sa1, sa2)
}
