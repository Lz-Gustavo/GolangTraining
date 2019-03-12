package main

import (
	"fmt"
	"math/rand"
)

func foo() func() int {

	return func() int {
		return rand.Int()
	}
}

func main() {

	f := foo()
	g := f()
	w := foo()()

	fmt.Printf("%T\n", f)
	fmt.Printf("%T\t%d\n", g, g)
	fmt.Printf("%T\t%d\n", w, w)
}
