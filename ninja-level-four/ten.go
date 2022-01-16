package main

import "fmt"

func main() {

	m := make(map[string][]string)

	first_person := []string{`bond_james`, `Shaken, not stirred`, `Martinis`, `Women`}
	second_person := []string{`moneypenny_miss`, `James Bond`, `Literature`, `Computer Science`}
	third_person := []string{`no_dr`, `Being evil`, `Ice cream`, `Sunset`}

	m["peter"] = first_person
	m["black"] = second_person
	m["brown"] = third_person

	//add a new record to the map
	m["jones"] = []string{`alexis sachez`, `meat stew`, `Science Fiction`, `Soccer`}

	//delete a record from the map and then print the map
	delete(m, "brown")
	delete(m, "peter")

	//print the map
	for k, v := range m {
		fmt.Println(k, v)
		for i, val := range v {
			fmt.Println(i, val)
		}
	}
}
