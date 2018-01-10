package main

import (
	"fmt"

	"github.com/golang-code/McLeod_Go_Course/10_Functions/excercise"
)

func main() {
	//excercise.SimpleFunc()
	//excercise.Bar("Luca")
	//fmt.Println(excercise.Zoo("CuCCO"))
	//x, y := excercise.Mouse("Pippo ", "Pluto")
	//fmt.Println(x, y)
	//fmt.Println(excercise.Variad(2, 3, 4, 5, 6, 7, 8, 9))
	// xi := []int{100, 3, 4, 5, 6, 7, 8, 9}
	// x := excercise.UnfurlingSlice(xi...)
	// fmt.Println(x)
	//lesson n 12_4 //DEFERRING A FUNC
	// defer foo()
	// bar()
	//excercise.Methods()
	//excercise.Interfaces()
	//excercise.AnonFunc1(100)
	// excercise.AssignFuncVar()
	//excercise.TestCallBack()
	//fmt.Println(excercise.Incrementor()())
	// x := excercise.FactorialFunc1(10)
	// fmt.Println(x)
	excercise.TestNinja04()
}

//lesson n 12_4 //DEFERRING A FUNC

func foo() {
	fmt.Println("hello from foo")
}

func bar() {
	fmt.Println("hello from bar")
}
