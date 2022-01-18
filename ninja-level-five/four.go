package main

import "fmt"

func main() {

	students := struct {
		firstname string
		lastname  string
		semester  []string
		grades    map[string][]string
	}{
		firstname: "Brian",
		lastname:  "Johns",
		grades:    map[string][]string{"physics": {"A", "B", "C", "D"}, "chemistry": {"D", "C", "A"}, "English": {"F"}},
		semester:  []string{"Spring", "Fall"},
	}
	fmt.Println(students)

	fmt.Println("-----------------------")
	fmt.Println(students.firstname)
	fmt.Println(students.lastname)

	for i, v := range students.semester {
		fmt.Println(i, v)
	}

	for k, v := range students.grades {
		fmt.Println(k)
		for j, x := range v {
			fmt.Println(j, x)
		}
	}
}
