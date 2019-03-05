package main

import "fmt"

func main() {

	birth := 1993
	for birth < 2019 {
		fmt.Println("On", birth, "I was alive!")
		birth++
	}
}
