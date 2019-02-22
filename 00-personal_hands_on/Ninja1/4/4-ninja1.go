package main

import "fmt"

type newVar int

var x newVar

func main() {

	fmt.Println(x)

	fmt.Printf("%T\n", x)

	x = 42

	fmt.Println(x)

}
