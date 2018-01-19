package main

import (
	"fmt"
	"math"
)

type Square struct {
	side float64
}
type Circle struct {
	radius float64
}

func (s Square) Area() float64 {
	return s.side * s.side
}

func (r Circle) Area() float64 {
	return r.radius * r.radius * math.Pi
}

type Shape interface {
	Area() float64
}

func main() {
	//defer bar()
	// x := foo()
	// b, s := bar()
	// fmt.Println(x, b, s)
	// y := []int{1, 2, 3, 4, 5, 6}
	// fmt.Println(fooVariadic(y...), barVariadic(y)) //The difference here is how the slice is passed to a func, with or witout the variadic inferrencing
	circ := Circle{
		radius: 13.554,
	}

	squa := Square{
		side: 13,
	}

	info(circ)
	info(squa)

	fmt.Println(anonFunc(12))

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

func info(s Shape) {
	fmt.Println(s.Area())
}

func anonFunc(y int) func() {
	return func() {
		y++
	}(y)
}
