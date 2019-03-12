package main

import "fmt"

func main() {

	fmt.Println("hello from main")

	func() {
		fmt.Println("Im an anonymous func")
	}()

}
