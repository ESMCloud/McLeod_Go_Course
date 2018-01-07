package excercise

import (
	"fmt"
)

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

// func (r receiver) identifier(parameters) (return(s)) { ... code}
func (s secretAgent) speak() { // speak() e' la funzione che verra associata a tutte le istanze di TIPO secretAgent, la variabile "s" avra' accesso ai dati dell'istanza (nome, cognome, etc)
	fmt.Println("My Name is ", s.last, ",", s.first, s.last, "!")
	if s.ltk {
		fmt.Println("...And I'm going to kill you!")
	} else {
		fmt.Println("...And I'll not hurt to anyone!")
	}
}

//Methods are the function(s) of a TYPE
func Methods() {
	sa1 := secretAgent{
		person: person{
			"James",
			"Bond",
		},
		ltk: true,
	}
	sa2 := secretAgent{
		person: person{
			"Miss",
			"Moneypenny",
		},
		ltk: false,
	}

	fmt.Println(sa1)
	fmt.Println(sa2)
	sa1.speak()
	sa2.speak()

}
