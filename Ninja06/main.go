package main

import (
	"fmt"
)

func main() {
	defer bar()
	// x := foo()
	// b, s := bar()
	// fmt.Println(x, b, s)
	y := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(fooVariadic(y...), barVariadic(y)) //The difference here is how the slice is passed to a func, with or witout the variadic inferrencing
}

func foo() int {
	return 1
}
func bar() (int, string) {
	defer func() {
		fmt.Println("foo DEFER ran")
	}()
	return 2, "ciao"
}

func fooVariadic(x ...int) int {
	var sum int
	for _, v := range x {
		sum += v
	}
	return sum
}

func barVariadic(x []int) int {
	var sum int
	for _, v := range x {
		sum += v
	}
	return sum
}
