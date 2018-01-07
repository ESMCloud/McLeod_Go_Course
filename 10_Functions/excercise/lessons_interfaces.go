package excercise

import (
	"fmt"
)

type person3 struct {
	first string
	last  string
}

type secretAgent2 struct {
	person3
	ltk bool
}

//// creo un nuovo tipo di tipo "interface" accostato a qualsiasi istanza che abbia il metodo speak()
type human interface {
	speak()
}

type hotdog int

// func (r receiver) identifier(parameters) (return(s)) { ... code}
func (s secretAgent2) speak() { // speak() e' la funzione che verra associata a tutte le istanze di TIPO secretAgent, la variabile "s" avra' accesso ai dati dell'istanza (nome, cognome, etc)
	fmt.Println("My Name is ", s.last, ",", s.first, s.last, "!", "The secret agent speak")
	if s.ltk {
		fmt.Println("...And I'm going to kill you!")
	} else {
		fmt.Println("...And I'll not hurt to anyone!")
	}
}

func (p person3) speak() { // speak() e' la funzione che verra associata a tutte le istanze di TIPO persona3, la variabile "s" avra' accesso ai dati dell'istanza (nome, cognome, etc)
	fmt.Println("Here I'm ", p.last, ",", p.first, p.last, "!", "The person speak")
}

//Interfaces are the function(s) of a TYPE
func Interfaces() {
	sa1 := secretAgent2{
		person3: person3{
			"James",
			"Bond",
		},
		ltk: true,
	}
	sa2 := secretAgent2{
		person3: person3{
			"Miss",
			"Moneypenny",
		},
		ltk: false,
	}

	p1 := person3{
		first: "Dr.",
		last:  "Yes",
	}

	fmt.Println(sa1)
	fmt.Println(sa2)
	sa1.speak()
	sa2.speak()

	fmt.Println(p1)
	p1.speak()

	IntMethod(sa1)
	IntMethod(sa2)
	IntMethod(p1)

	// conversion
	var x hotdog = 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	var y int
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)

}

//IntMethod are the function(s) of a TYPE
func IntMethod(h human) {
	switch h.(type) {
	case person3:
		fmt.Println("I was passed into barrrrrr", h.(person3).last)
	case secretAgent2:
		fmt.Println("I was passed into barrrrrr", h.(secretAgent2).last)
	}
	fmt.Println("I was passed into IntMethod", h)
}

//AnonFunc1 is
func AnonFunc1(get int) {

	func(x int) {
		fmt.Println("Ciao, hai ", get, " anni!")
	}(get) //sintassi per passare valori alla funzione anonima

}

//AssignFuncVar is
func AssignFuncVar() {

	x := func() {
		fmt.Println("My first func expression!")
	} //sintassi per passare valori alla funzione anonima
	x()
	fmt.Printf("%T\t, %p,\t", x, x)
}
