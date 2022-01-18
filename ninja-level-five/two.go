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
		firstname: "Mary",
		lastname:  "Wany",
		icecreamflavours: []string{
			"cocoa", "caramel", "cocoa",
		},
	}

	personmap := make(map[string]person)

	personmap[p1.lastname] = p1
	personmap[p2.lastname] = p2

	//loop through the map and print all the items stored in the map
	fmt.Println("----------------------------")
	for _, v := range personmap {
		fmt.Println(v.firstname)
		fmt.Println(v.lastname)

		//range over the slice of ice flavours
		for _, x := range v.icecreamflavours {
			fmt.Println(x)
		}
		fmt.Println("----------------------------")
	}
}
