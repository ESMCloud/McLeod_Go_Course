package main

import "fmt"

func main() {
	var a string
	a := fmt.Scan("ask for a name:")
	fmt.Println("Hello %s!", a)
}
