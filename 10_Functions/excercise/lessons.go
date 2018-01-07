package excercise

import (
	"fmt"
)

//the structure of a FUNCTION
//func (r receiver) identifier(parameters) (return(s)) { ... }

//SimpleFunc is an example of the most simple function: with no parameters and no return
func SimpleFunc() {
	fmt.Println("ciao da simpleFunc")
}

//Bar is a simple function that takes a string as argument
func Bar(s string) {
	fmt.Println("ciao dalla Bar function, ", s)
}

//Zoo is a simple function that takes a string as argument and return a string
func Zoo(s string) string {
	return fmt.Sprint("ciao dalla Zoo function, ", s)
}

//Mouse is a simple function that takes a string as argument with multiple returns
func Mouse(fn string, ln string) (string, bool) {
	return fmt.Sprint("ciao dalla Mouse function, ", fn, ln), true
}

//Variad is a  function that takes unlimited number of parameters as argument
func Variad(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	sum := 0
	for _, v := range x {
		sum += v
	}
	//fmt.Println(sum)
	return sum
}
