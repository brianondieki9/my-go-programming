package main

import "fmt"

type vehicle struct {
	doors string
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {

	t := truck{
		vehicle:   vehicle{doors: "four", color: "green"},
		fourWheel: true,
	}

	s := sedan{
		vehicle: vehicle{doors: "five", color: "blue"},
		luxury:  true,
	}

	//print out the values of the types created
	fmt.Println(t)
	fmt.Println(s)

	//printing a single field for each type value created

	fmt.Println("---------------------------")
	fmt.Println("Doors: ", t.doors, "Color:", t.color, "is fourwheel:", t.fourWheel)

	fmt.Println("---------------------------")

	fmt.Println("Doors: ", s.doors, "Color:", s.color, "is luxury:", s.luxury)

}
