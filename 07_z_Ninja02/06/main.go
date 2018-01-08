package main

import (
	"fmt"
)

const (
	// a = 2017 + iota
	// b = 2017 + iota
	// c = 2017 + iota
	// d = 2017 + iota
	a = 2017 + iota
	b = a + iota
	c = a + iota
	d = a + iota
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
