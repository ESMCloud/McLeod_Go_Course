package main

import "fmt"

func main() {

	// x := 83 / 40
	// y := 83 % 40
	// fmt.Println(x, y)

	// x := 0
	// for {
	// 	if x > 100 {
	// 		break
	// 	}
	// 	if x%2 != 0 {
	// 		x++
	// 		continue
	// 	}
	// 	fmt.Println(x)
	// 	x++
	// }

	x := 0
	for {
		if x > 100 {
			break
		}
		if x%2 == 0 {
			fmt.Println(x)
			x++
		}
	}
	fmt.Println("Done")
}
