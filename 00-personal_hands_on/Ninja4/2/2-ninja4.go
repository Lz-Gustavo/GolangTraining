package main

import "fmt"

func main() {

	// make syntax
	//xi := make([]int, 0, 16)
	//xi = append(xi, 1, 2, 3, 4, 5, 6, 7, 8)

	// composite literal
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	for i, v := range xi {
		fmt.Println(i, v)
	}

	fmt.Printf("Slice type: %T\n", xi)
}
