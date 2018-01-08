package main

import (
	"fmt"
)

func main() {
	var x string
	fmt.Println("Give your name:")
	fmt.Scan(&x)
	fmt.Println("Hello &s", x)
}
