package main

import (
	"fmt"

	"github.com/golang-code/McLeod_Go_Course/07_Ninja03/test"
)

func main() {

	if x := 41; x == 40 {
		fmt.Println("our value was 40")
	} else if x == 42 {
		fmt.Println("our value was 42")
	} else {
		fmt.Println("fuck")
	}
	for i := 0; i < 100; i++ {
		if i == 0 {
			fmt.Printf("Il numero %d e' ZERO\n", i)
		} else if i%2 == 0 {
			fmt.Printf("Il numero %d e' PARI\n", i)
		} else {
			fmt.Printf("Il numero %d e' DISPARI\n", i)
		}
	}

	test.CasesSw()
	test.Print1000()
	test.PrintRune()
	test.YearsAlive()
	test.PrintModulo()
	test.SimpleIF()
	test.SwitchEmpty()
	test.SwitchString()
	test.PrintBool()

}
