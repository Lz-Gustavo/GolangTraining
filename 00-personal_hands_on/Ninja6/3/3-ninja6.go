package main

import "fmt"

func foo() {
	defer bar()
	fmt.Println("Hello from foo")
}

func bar() {
	fmt.Println("Hello from bar")
}

func main() {
	foo()
}
