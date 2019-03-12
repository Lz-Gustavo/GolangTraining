package main

import "fmt"

func main() {

	fmt.Println("Hello from main")

	x := func(n int) {
		fmt.Println("Hello from non-named func assigned to a variable... funcs are first-class citizens in go hahah and u just passed me the value", n)
	}

	x(10)
}
