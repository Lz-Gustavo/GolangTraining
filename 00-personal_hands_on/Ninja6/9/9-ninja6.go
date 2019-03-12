package main

import "fmt"

// callback function 'foo'
func foo(f func(a ...interface{}) (int, error)) bool {

	_, err := f("Strange way to call a Println hahha")

	if err != nil {
		return false
	}
	return true
}

func main() {

	status := foo(fmt.Println)

	fmt.Println(status)
}
