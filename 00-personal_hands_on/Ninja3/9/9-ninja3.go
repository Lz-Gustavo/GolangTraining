package main

import "fmt"

func main() {

	favSport := "Chess"

	switch favSport {
	case "Soccer":
		fmt.Println("Ur dumb")
	case "Baseball":
		fmt.Println("America fanboy")
	case "Swimming":
		fmt.Println("Im starting to like you...")
	case "Chess":
		fmt.Println("THATS MY BOY")
	default:
		fmt.Println("just dont say LoL...")
	}
}
