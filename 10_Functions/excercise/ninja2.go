package excercise

import (
	"fmt"
)

type man struct {
	first string
	last  string
	age   uint8
}

func (m man) speak() {
	fmt.Println("Ehy", m.first, " ", m.last)
	fmt.Println("You're", m.age, "years old!")
}

//TestNinja04 is
func TestNinja04() {
	m := man{
		first: "Luca",
		last:  "Canc",
		age:   99,
	}

	m.speak()
}
