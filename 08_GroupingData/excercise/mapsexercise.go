package excercise

import (
	"fmt"
)

//Mapsex01 Func
func Mapsex01() {
	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 24,
	}
	fmt.Println(m)
	fmt.Println(m["James"])
	fmt.Println(m["NONESISTO"]) // se richiedo una chiave non esistente, il valore di ritorno e' zero
	// Per controllare se esiste una chiave, si usa l'idioma "Comma / , ok idiom"
	if v, ok := m["NONESISTO"]; ok {
		fmt.Println(v)
	} else {
		fmt.Println("Chiave non esistente")
	}

	if v, ok := m["Miss Moneypenny"]; ok {
		fmt.Println("QUESTO E IL DATO: ", v)
	} else {
		fmt.Println("Chiave non esistente")
	}
	m["todd"] = 33
	for k, v := range m {
		fmt.Println(k, v)
	}

	xi := []int{4, 5, 7, 8, 9, 42}

	for i, v := range xi {
		fmt.Println(i, v)
	}

	delete(m, "James")
	fmt.Println(m)
	delete(m, "XYZ") // Nessun errore se cancello una entry inesistente, usare il ", ok"
	if _, ok := m["XYZ"]; ok {
		delete(m, "XYZ")
	} else {
		fmt.Println("Key not present")
	}

}
