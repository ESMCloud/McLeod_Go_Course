package excercise

import (
	"fmt"
)

//the structure of a FUNCTION
//func (r receiver) identifier(parameters) (return(s)) { ... }

//Pointers is
func Pointers() {
	a := 42
	fmt.Println(a)
	fmt.Println(&a)

	b := &a
	fmt.Printf("%T\n", b)
	fmt.Println(b)
	fmt.Println(*b)
	*b = 44
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", *b)
	fmt.Println(b)
	fmt.Println(*b)
	fmt.Println(a)
	z := 5
	fmt.Println("This is BEFORE ther function UsePointers")
	fmt.Println(z)
	fmt.Println(&z)
	UsePointers(&z)
	fmt.Println("This is after ther function UsePointers")
	fmt.Println(z)
	fmt.Println(&z)
}

//UsePointers is
func UsePointers(p *int) {
	*p = 0
	fmt.Println("This is the address of p ", p)
	fmt.Println(*p)
}
