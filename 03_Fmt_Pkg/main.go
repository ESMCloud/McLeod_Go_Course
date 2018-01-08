package main

import (
	"fmt"
)

var a int
var b string = "James Bond"
var c bool
var d bool = true

func main() {
	e := 32
	f := "Shake not stirred"
	g := `Miss Moneypenny says, "Helloooooo James!"`
	h := `Q. says,
	"I have some new toys for your, James."`

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(b, "says,", f)
	fmt.Println(g)
	fmt.Printf("%v", a)
	fmt.Println(h)

	s := fmt.Sprint(a, "Something else", b)
	fmt.Println(s)
}
