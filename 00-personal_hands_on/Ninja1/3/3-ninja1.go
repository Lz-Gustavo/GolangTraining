package main

import "fmt"

// global
var x = 42
var y = "James Bound"
var z = true

func main() {

	s := fmt.Sprintf("%d\t%s\t%v", x, y, z)
	fmt.Println(s)
}
