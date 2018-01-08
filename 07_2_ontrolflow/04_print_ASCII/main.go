package main

import "fmt"

func main() {

	for x := 0; x <= 244; x++ {

		fmt.Printf("Decimal: %d\tHexadecimal: %#x\tUnicode: %#U\tBinary: %b\n", x, x, x, x)
	}
}
