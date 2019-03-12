package main

import "fmt"

func main() {

	semDown, semUp, semGet := lockStatus()

	fmt.Println("===Semaphore behaviour===")
	fmt.Println("Current value is:", semGet())

	semUp()
	fmt.Println("Now is:", semGet())

	semDown()
	fmt.Println("And now is:", semGet())
}

func lockStatus() (func(), func(), func() int) {
	value := 0

	return func() {
			value = 0
		}, func() {
			value = 1
		}, func() int {
			return value
		}
}
