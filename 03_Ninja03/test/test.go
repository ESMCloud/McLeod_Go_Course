package test

import (
	"fmt"
	"time"
)

//CasesSw FUNC
func CasesSw() {

	n := "Bond"

	switch n {
	case "Money penny", "Bond", "Do Not":
		fmt.Println("money p or bond or not")
		//fallthrough
	case "m":
		fmt.Println("This is m")
		//fallthrough
	case "Q":
		fmt.Println("Q")
		//fallthrough //OCCHIO!!
	// case (7 == 9):
	// 	fmt.Println("Not print 1")
	// case (4 == 4):
	// 	fmt.Println("Also true, does it print?")
	default:
		fmt.Println("This is  default")
	}
}

//Print1000 FUNC
func Print1000() {
	for i := 0; i <= 1000; i++ {
		fmt.Println(i)
	}
}

//PrintRune FUNC
func PrintRune() {
	for i := 65; i <= 90; i++ {
		fmt.Printf("%d\n", i)
		for x := 0; x < 3; x++ {
			fmt.Printf("\t%#U\n", i)
		}
	}
}

//YearsAlive FUNC
func YearsAlive() {
	bd := 1986
	yn := time.Now().Year()

	for bd <= yn {
		//fmt.Printf("%T", bd)
		fmt.Println(bd)
		bd++
	}
}

//PrintModulo FUNC
func PrintModulo() {
	for i := 10; i <= 100; i++ {
		fmt.Printf("For number %v the Modulus/remainder is %v\n", i, i%4)
	}
}

//SimpleIF FUNC
func SimpleIF() {
	var x int
	fmt.Scanf("%d", &x)
	if x == 1 {
		fmt.Println("Got IT")
	} else if x == 2 {
		fmt.Println("Fuck 2")
	} else {
		fmt.Println("Out")
	}
}

//SwitchEmpty FUNC
func SwitchEmpty() {
	// var x int
	// fmt.Scanf("%d", &x)
	switch {
	case false:
		fmt.Println("This not print")
	case true:
		fmt.Println("This will print")
	}
}

//SwitchString FUNC
func SwitchString() {
	favSport := "Surfing"
	switch favSport {
	case "Dio":
		fmt.Println("This not print")
	case "Surfing":
		fmt.Println("Here we are")
	}
}

//PrintBool FUNC

func PrintBool() {
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)
}
