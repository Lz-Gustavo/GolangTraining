package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

func changeMe(p *person, fn string, ln string, age int) {

	(*p).firstName = fn
	(*p).lastName = ln
	(*p).age = age
}

func main() {

	p := person{
		firstName: "George",
		lastName:  "Lucas",
		age:       41,
	}
	fmt.Println(p)

	changeMe(&p, "Ronaldo", "Nazario", 39)
	fmt.Println(p)
}
