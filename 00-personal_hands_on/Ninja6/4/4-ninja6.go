package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (s person) speak() {
	fmt.Println("Hi! my name is", s.first, s.last, "and im", s.age, "years old")
}

func main() {

	p := person{
		first: "Michael",
		last:  "Faraday",
		age:   32,
	}

	p.speak()
}
