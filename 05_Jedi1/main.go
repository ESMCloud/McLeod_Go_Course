package main

import "fmt"

type mytype int

var x mytype
var y int

func main() {
	fmt.Printf("%v, %T\n", x, x)
	x = 42
	fmt.Println(x)
	y = int(x)
	fmt.Printf("%T", y)
}
