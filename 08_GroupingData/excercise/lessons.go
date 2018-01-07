package excercise

import (
	"fmt"
)

//First FUNC
func First() {
	fmt.Println("ciao")
}

//TestArray FUNC
func TestArray() {
	var x [5]int
	fmt.Println(x)
	x[3] = 42
	fmt.Println(x)
	fmt.Println(len(x))
}

//TestSlices FUNC
func TestSlices() {
	//COMPOSITE LITERAL
	// x := type {values}
	x := []int{4, 5, 7, 8, 42} //A SLICE allows to group together VALUES of the SAME TYPE
	fmt.Println(x)
	fmt.Println("Lunghezza dello slice: ", len(x))
	fmt.Println("Primo Elemento dello Slice ", x[0])
	for i, v := range x {
		fmt.Println(i, v)
	}

}

//SlicingSlice FUNC
func SlicingSlice() {
	//COMPOSITE LITERAL
	// x := type {values}
	x := []int{4, 5, 7, 8, 42} //A SLICE allows to group together VALUES of the SAME TYPE
	fmt.Println(x)
	fmt.Println(x[1:])
	fmt.Println(x[1:3])
	fmt.Println(x[:3])

	for i, v := range x {
		fmt.Println(i, v)
	}

	for i := 0; i < len(x); i++ {
		fmt.Println(i, x[i])
	}

	x = append(x, 100, 102)
	fmt.Println(x)

	y := []int{300, 400, 500}

	x = append(x, y...)
	fmt.Println(x)

	// Remove / Deleting element(s) from a SLICE with append. You overwrite part of that slice
	x = append(x[:2], x[4:]...)
	fmt.Println(x)

}

//SlicingMake FUNC
func SlicingMake() {
	//Creare uno slice con make permette di settare in anticipo la dimensione dell'array sottostante, evitando la creazione / distruzione degli array dinamici.

	x := make([]int, 10, 12)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	x[0] = 42
	x[9] = 999

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 3423)

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 3424)

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 3425)

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
}

//MdimensionalSlice FUNC
func MdimensionalSlice() {
	//Creare uno slice con make permette di settare in anticipo la dimensione dell'array sottostante, evitando la creazione / distruzione degli array dinamici.

	//jb := make([]string, 10) // non e' necessario dichiarare la capacita
	jb := []string{"James", "Bond", "Choccolate", "Martini"}
	fmt.Println(jb)
	mp := []string{"Miss", "Moneypenny", "Strawberry", "Hezelnut"}
	fmt.Println(mp)

	xp := [][]string{jb, mp} // MULTI DIMENSIONAL SLICE
	fmt.Println(xp)
	//TODO, NAVIGATE MD SLICES
}

//UnderlyingArray FUNC
func UnderlyingArray() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(x)

	y := append(x[:2], x[5:]...) // the underlying array stores the values of the new slice

	//z := append(x, 43, 43, 43, 43, 43, 43, 44) // new underlying array allocated
	fmt.Println(x) // THE VALUE OF THE UNDERLYING ARRAY (X in this case, is CHANGED TOO!)
	fmt.Println(y)
	//fmt.Println(z)
}

//I tipi MAP DEVONO ESSERE CREATI!!! Exemple: make(map["key type"]"data type")

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
