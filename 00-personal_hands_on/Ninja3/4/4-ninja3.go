package main

import "fmt"

func main() {

	birth := 1994
	for {

		fmt.Println(birth)
		birth++

		if birth > 2019 {
			break
		}
	}
}
