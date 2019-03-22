package main

import "fmt"

type person struct {
	name     string
	lastName string
	age      int
}

func (p *person) Speak() {
	fmt.Println("Hi, my name is", p.name, p.lastName, "and im", p.age, "years old")
}

type human interface {
	Speak()
}

func saySomething(h human) {
	h.Speak()
}

func main() {

	p := person{
		name:     "Edward",
		lastName: "Scissorhands",
		age:      28,
	}

	// this is possible
	saySomething(&p)

	p2 := &person{
		name:     "Kilo",
		lastName: "Ren",
		age:      33,
	}

	// and this
	saySomething(p2)
}
