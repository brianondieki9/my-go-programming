package main

import "fmt"

type person struct {
	firstname        string
	lastname         string
	icecreamflavours []string
}

func main() {
	p1 := person{
		firstname: "brian",
		lastname:  "johns",
		icecreamflavours: []string{
			"vanilla", "latee", "chocolate",
		},
	}

	p2 := person{
		firstname:        "Mary",
		lastname:         "Wany",
		icecreamflavours: []string{"cocoa", "caramel", "cocoa"},
	}

	fmt.Println(p1)
	//fmt.Printf("%T", p1)
	fmt.Println(p2)

	fmt.Println("----------------------------------")
	for i, v := range p1.icecreamflavours {
		fmt.Println(i, v)
	}
	fmt.Println("----------------------------------")

	for i, v := range p2.icecreamflavours {
		fmt.Println(i, v)
	}

	fmt.Println("----------------------------------")
}
