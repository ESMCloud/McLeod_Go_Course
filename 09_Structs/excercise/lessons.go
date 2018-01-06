package excercise

import (
	"fmt"
)

//lesson 1, struct definition
type person struct {
	first string
	last  string
	age   int8
}

//Lesson 2, embedded data structure

type secretAgent struct {
	person
	ltk bool
}

//FirstStruct01 FUNC
func FirstStruct01() {
	//HINT
	p1 := person{
		first: "James",
		last:  "Bond",
		age:   32,
	}
	p2 := person{
		first: "Miss",
		last:  "Monepenny",
		age:   24,
	}
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p1.first, p1.last)
	fmt.Println(p2.first, p2.last)
}

//EmbeddedDataStructure FUNC
func EmbeddedDataStructure() {
	//HINT
	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   32,
		},
		ltk: true,
	}
	sa2 := secretAgent{
		person: person{
			first: "Miss",
			last:  "Monepenny",
			age:   24,
		},
		ltk: false,
	}
	fmt.Println(sa1)
	fmt.Println(sa2)
	fmt.Println(sa1.first, sa1.last, sa1.ltk)
	fmt.Println(sa2.first, sa2.last, sa2.ltk)
}

//AnonStructure FUNC
func AnonStructure() {
	// quando si vuole usare una struttura limitata come scope ad una parte di codice (niente pollution e variabili in giro per il progetto), si puo' usare una struttura anonima
	p1 := struct {
		first string
		last  string
		age   uint8
	}{
		first: "James",
		last:  "Bond",
		age:   32,
	}

	fmt.Println(p1)
	fmt.Printf("%T", p1)
}
