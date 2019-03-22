package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	First string
	Age   int
}

func main() {

	u1 := User{
		First: "James",
		Age:   32,
	}

	u2 := User{
		First: "Moneypenny",
		Age:   27,
	}

	u3 := User{
		First: "M",
		Age:   54,
	}

	users := []User{u1, u2, u3}

	fmt.Println(users)

	// your code goes here

	usersJS, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(usersJS))

	var fd *os.File
	fd, err = os.Create("./json-output.json")
	if err != nil {
		fmt.Println(err)
	}

	_, err = fmt.Fprint(fd, string(usersJS))
	if err != nil {
		fmt.Println(err)
	}
}
