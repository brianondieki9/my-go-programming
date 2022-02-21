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

//all types that implement the speak() function below will also be of type human(inheritance)
type human interface {
	speak()
}

func (s secretAgent) speak() {
	fmt.Println("I am", s.firstname, s.lastname)
}

func (p person) speak() {
	fmt.Println("I am", p.firstname, p.lastname)
}

//inheritance allows polymophism to be used, this function can take different types that are of type human
func bar(h human) {
	fmt.Println("I have been passed in Bar", h)
	h.speak()
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
		licenseToKill: false,
	}

	p1 := person{
		firstname: "Dr.",
		lastname:  "Yes",
	}

	sa1.speak()
	sa2.speak()

	sa1.countCharacters()
	sa2.countCharacters()

	fmt.Println(sa1, sa2)

	bar(sa1)
	bar(sa2)
	bar(p1)
}
