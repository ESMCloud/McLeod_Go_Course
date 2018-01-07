package excercise

import (
	"fmt"
)

//Clousure1 is
func Clousure1() {
	t := evenSum(sum, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}...)
	fmt.Println(t)
	Incrementor()
}

//Incrementor s
func Incrementor() func() int {
	var x int
	return func() int {
		return x + 101
	}
}
