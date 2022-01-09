package main

import "fmt"

func main() {

	m := make(map[string]int)

	m["k1"] = 77
	m["k2"] = 88

	fmt.Println(m)

	for k, v := range m {
		fmt.Println(k, v)
	}

	x := make(map[int]string)
	x[0] = "ssss"
	x[44] = "badb"

	fmt.Println(x)

	for k, v := range x {
		fmt.Println(k, v)
	}
	fmt.Println(len(x))

	//delete key,value from map
	delete(x, 44)
	fmt.Println(x)

	v, prs := x[0]

	if prs {
		fmt.Println("value present")
	} else {
		fmt.Println("value is not present")
	}

	fmt.Println(prs)
	fmt.Println(v)
}
