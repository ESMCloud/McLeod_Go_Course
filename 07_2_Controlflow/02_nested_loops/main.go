package main

import (
	"fmt"
)

func main() {
	for x := 0; x <= 10; x++ {
		fmt.Printf("The outer loop counter: %d\n", x)
		for y := 0; y < 3; y++ {
			fmt.Printf("\t\tThe inner loop counter: %d\n", y)
		}
	}

}
