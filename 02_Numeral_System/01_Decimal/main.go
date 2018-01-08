package main

import (
	"fmt"
)

func main() {
	// 	for g := 0; g <= 170; g++ {
	// 		fmt.Printf("%d \t %b \t %x \t %q \n", g, g, g, g)
	// 	}
	var x int
	fmt.Println("Give me a binary number: ")
	fmt.Scanf("%b", &x)
	fmt.Printf("%d\n", x)

}
