package main

import "fmt"

func main() {

	var x int
	fmt.Scanf("%d", &x)

	if x < 3 {
		fmt.Println("nice", x)
	} else if x < 10 {
		fmt.Println("big", x)
	} else {
		fmt.Println("huge", x)
	}
}
