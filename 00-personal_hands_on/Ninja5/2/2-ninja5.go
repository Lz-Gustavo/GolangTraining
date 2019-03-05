package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	iceCream  []string
}

func main() {

	mapPersons := make(map[string]person)

	p1 := person{
		firstName: "Michael",
		lastName:  "Jones",
		iceCream:  []string{"Strawberry", "Pistachio"},
	}
	p2 := person{
		firstName: "Matthew",
		lastName:  "Douglas",
		iceCream:  []string{"Chocolate", "Vanilla"},
	}

	mapPersons["Jones"] = p1
	mapPersons["Douglas"] = p2

	for _, p := range mapPersons {

		fmt.Println(p.firstName, p.lastName)
		for _, v := range p.iceCream {
			fmt.Println(v)
		}
		fmt.Println()
	}
}
