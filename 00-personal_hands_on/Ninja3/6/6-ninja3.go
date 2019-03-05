package main

import "fmt"

func main() {

	cond := true

	if cond {
		fmt.Println("Thats if statement")
		cond = false
	}

	if !cond {
		fmt.Println(" In action ...")
	}

}
