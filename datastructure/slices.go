package main

import "fmt"

func main() {
	//composite literal example
	//x := type{values}

	x := []int{2, 3, 4, 5, 6, 7}
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	//printing using range
	for k, v := range x {
		fmt.Println(k, v)
	}

	//slicing slices
	fmt.Println(x[0:6])

	for i := 0; i < len(x); i++ {
		fmt.Println(x[i])
	}

	//append slices
	x = append(x, 155, 156)
	fmt.Println(x)

	y := []int{88, 89, 90}

	x = append(x, y...)

	fmt.Println(x)

	//delete from slice
	x = append(x[0:8], x[9:]...)
	fmt.Println(x)

	//create slice using make built in function
	z := make([]int, 10, 100)
	z[88] = 56
	fmt.Println(z)
	fmt.Println("length of the slice is", len(z))
	fmt.Println("capacity is", cap(z))

}
