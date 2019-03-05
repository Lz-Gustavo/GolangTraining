package main

import "fmt"

func main() {

	p := struct {
		name string
		role string
		age  int
	}{
		name: "Louis",
		role: "Engineer",
		age:  27,
	}

	fmt.Println(p.name, p.role, p.age)
}
