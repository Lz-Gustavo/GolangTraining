package main

import "fmt"

func main() {

	matrix := make([][]string, 2, 10)
	fmt.Println(matrix)

	matrix[0] = append(matrix[0], "James", "Bond", "Shaken, not stirred")
	matrix[1] = append(matrix[1], "Miss", "Moneypenny", "Helloooooooo, James")

	for i, v := range matrix {
		fmt.Println(i, v)
	}

}
