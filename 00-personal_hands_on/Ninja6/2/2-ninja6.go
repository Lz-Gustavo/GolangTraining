package main

import "fmt"

func main() {

	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	sum := foo(xi...)
	sum2 := bar(xi)

	fmt.Println(sum, sum2)
}

func foo(xi ...int) int {

	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}

func bar(xi []int) int {

	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}
