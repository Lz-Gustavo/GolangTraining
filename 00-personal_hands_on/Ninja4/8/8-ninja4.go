package main

import "fmt"

func main() {

	m := make(map[string][]string)

	m["bond_james"] = []string{`Shaken, not stirred`, `Martinis`, `Women`}
	m["moneypenny_miss"] = []string{`James Bond`, `Literature`, `Computer Science`}
	m["no_dr"] = []string{`Being evil`, `Ice cream`, `Sunsets`}

	for i := range m {
		fmt.Println(i, "likes:")
		for j, v := range m[i] {
			fmt.Printf("\t%d\t%s\n", j, v)
		}
	}
}
