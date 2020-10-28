package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("invalid input arguments")
		os.Exit(1)
	}
	s := os.Args[1]

	r := camelCase(s)
	fmt.Println(r)
}

func camelCase(s string) int {
	c := 1

	for _, char := range s {
		schar := string(char)
		//	fmt.Println(string(char))
		if schar == strings.ToUpper(schar) {
			c++
		}
	}

	return c
}
