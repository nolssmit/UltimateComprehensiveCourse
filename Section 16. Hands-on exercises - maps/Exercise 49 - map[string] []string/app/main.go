package main

import (
	"fmt"
)

func main() {
	m := make(map[string] [] string) // map has value of a slice of string
	m["bond_james"] = []string{"shaken, not stirred", "martinis", "fast cars"}
	m[`moneypenny_jenny`] = []string{`intelligence`, `literature`, `computer science`}
	m[`no_dr`] = []string{`cats`, `ice cream`, `sunsets`}

	fmt.Println("=====================================================================")
	fmt.Printf("%#v\n", m)
	fmt.Println("=====================================================================")

	for k, v := range m {
		fmt.Println(k)
		fmt.Println(" ")
		for i, v2 := range v {  // the value is the slice of string
			fmt.Println(i, v2)
		}
		fmt.Println("-------------------------------------------")
	}
}