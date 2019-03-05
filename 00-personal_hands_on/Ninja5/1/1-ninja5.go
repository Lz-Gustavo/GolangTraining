package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	iceCream  []string
}

func main() {

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

	fmt.Println(p1.firstName, p1.lastName)
	for _, v := range p1.iceCream {
		fmt.Println(v)
	}

	fmt.Println(p2.firstName, p2.lastName)
	for _, v := range p2.iceCream {
		fmt.Println(v)
	}
}
