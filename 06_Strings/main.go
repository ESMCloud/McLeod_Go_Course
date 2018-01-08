package main

import (
	"fmt"
)

func main() {
	s := "Hello World!"
	fmt.Println(s)
	fmt.Printf("%T\n", s)
	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

	// How to see UTF-8 encoding of a slice of string?

	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U ", s[i])
	}
	fmt.Println("")

	for i, v := range s {
		fmt.Printf("\nAt index position: %d we have hex %#x corresponding to UTF-8 %#U\n", i, v, v)
		// fmt.Printf

	}
}
