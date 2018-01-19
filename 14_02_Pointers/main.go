package main

import (
	"fmt"
)

func main() {
	fmt.Println("")
	fmt.Println("Video 1 on pointers")

	// & give us the address of the value
	// * give us the value stored at that address

	a := 42
	fmt.Printf("%v\n", a)
	fmt.Printf("%T\n", a)
	fmt.Printf("%v\n", &a)
	fmt.Printf("%T\n", &a)
	fmt.Println("")

	b := "Ciao"
	fmt.Printf("%v\n", b)
	fmt.Printf("%T\n", b)
	fmt.Printf("%v\n", &b)
	fmt.Printf("%T\n", &b)
	fmt.Println("")

	c := []string{"Mio", "Tuo", "Suo"}
	fmt.Printf("%v\n", c)
	fmt.Printf("%T\n", c)
	fmt.Printf("%v\n", &c)
	for _, v := range c {
		fmt.Printf("%v\n", v)
		fmt.Printf("%v\n", &v)

	}
	fmt.Printf("%v %v %v\n", &c[0], &c[1], &c[2])
	fmt.Printf("%T\n", &c)
	fmt.Println("")

	fmt.Printf("This is the type of a: %T\n", a)
	ca := &a
	fmt.Printf("This is the type of ca: %T\n", ca)
	fmt.Printf("This is the type of *ca: %T\n", *ca)
	fmt.Printf("This is the type of &ca: %T\n", &ca) // **int PUNTATORE A PUNTATORE DI INT
	fmt.Println("This is the address of a: ", &a)
	fmt.Println("This is the address of ca: ", ca)
	fmt.Println("This is a the valueof a: (1)", *&a)
	fmt.Println("This is the value of ca (1): ", *ca)
	fmt.Println("This is the value of &ca (1): ", &ca)

	*ca = 55 //Assign 55 to the VALUE of address ca (ca is of type pointer to a memory location, *ca show the value: deference )
	fmt.Println("This is a the valueof a: (1)", *&a)
	fmt.Println("This is the value of ca (1): ", *ca)

	fmt.Println("")
	fmt.Println("Video 2 on pointers")
	x := 0
	fmt.Println("This is x: ", x)
	fmt.Println("This is the address of x: ", &x)
	foo(x)
	fmt.Println("This is x after foo(): ", x)
	pointers(&x)
	fmt.Println("This is x after pointers(): ", x)
	//testpointers(*x)
	testpointers(x)
	testpointers2(&x)
	fmt.Println("")
	fmt.Println("Video 3 on pointers - method set")
}

// No pointer
func foo(y int) {
	fmt.Println("This is y in foo() (imported by x): ", y)
	y = 43
	fmt.Println("This is y in foo() after assignment: ", y)
}

// No pointer
func pointers(p *int) {
	fmt.Printf("\n")
	fmt.Printf("This is the type of p: %T\n", p)
	fmt.Println("This is p in pointers() (imported by x): ", p)
	fmt.Println("This is *p in pointers() (imported by x): ", *p)
	*p = 55
	fmt.Println("This is p in pointers() after assignment: ", p)
	fmt.Println("This is *p in pointers() after assignment: ", *p)
}

func testpointers(p ...interface{}) {
	fmt.Printf("\n")
	fmt.Printf("This is the type of p of testpointers: %T\n", p)
	fmt.Printf("This is the type of &p of testpointers: %T\n", &p)
	fmt.Println("This is p in testpointers() (imported by x): ", p)
}

func testpointers2(p ...interface{}) {
	fmt.Printf("\n")
	fmt.Printf("This is the type of p of testpointers2: %T\n", p)
	fmt.Println("This is p in testpointers2() (imported by x): ", p)
	fmt.Printf("This is the type of &p of testpointers2: %T\n", &p)
	fmt.Println("This is &p in testpointers2() (imported by x): ", &p)
	fmt.Printf("This is the type of &p[0] of testpointers2: %T\n", p[0])
	fmt.Println("This is &p[0] in testpointers2() (imported by x): ", &p)
}
