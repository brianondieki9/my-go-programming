package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person
	licenseToKill bool
}

func main() {
	p1 := person{
		first: "brian",
		last:  "ondieki",
		age:   23,
	}

	p2 := person{
		first: "Justus",
		last:  "Ondieki",
		age:   44,
	}
	fmt.Println(p1.first, p1.last, p1.age)
	fmt.Println(p2.first, p2.last, p2.age)

	//creating secretagen type value
	sa1 := secretAgent{
		person:        person{first: "John", last: "Wanangwe", age: 23},
		licenseToKill: true,
	}

	sa2 := secretAgent{
		person:        person{first: "Alice", last: "Muchugi", age: 55},
		licenseToKill: false,
	}
	//printing the value of secret agent struct values that includes another struct
	fmt.Println(sa1.first, sa1.last, sa1.age, sa1.licenseToKill)
	fmt.Println(sa2.person.first, sa2.person.last, sa2.person.age, sa2.licenseToKill)
	fmt.Printf("%T\n", sa1)

}
