package main

import "fmt"

func main() {
	switch {
	case 1 == 1:
		fmt.Println("not specified")
		fallthrough
	case 4 != 4:
		fmt.Println("at all")
		fallthrough
	default:
		fmt.Println("note: never use fallthrough!")
	}
}
