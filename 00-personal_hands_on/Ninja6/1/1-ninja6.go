package main

import (
	"fmt"
	"math/rand"
)

func main() {

	x := foo()
	y, s := bar()

	fmt.Println(x, y, s)
}

func foo() int {
	return rand.Int()
}

func bar() (int, string) {
	return rand.Int(), "random string"
}
