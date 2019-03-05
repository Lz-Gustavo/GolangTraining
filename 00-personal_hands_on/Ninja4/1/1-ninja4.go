package main

import "fmt"

func main() {

	ai := [5]int{1, 2, 3, 4, 5}
	fmt.Println(ai)

	for i, v := range ai {
		fmt.Println(i, v)
	}

	fmt.Printf("Array type: %T\n", ai)
}
