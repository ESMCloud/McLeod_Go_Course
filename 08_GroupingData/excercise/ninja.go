package excercise

import (
	"fmt"
)

//Ninja04_01 Func
func Ninja04_01() {
	x := [5]int{10, 20, 30, 40, 50}
	// x[0] = 10
	// x[1] = 20
	// x[2] = 30
	// x[3] = 40
	// x[4] = 50
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T", x)
}

//Ninja04_02 Func
func Ninja04_02() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", x)
}

//Ninja04_03 Func
func Ninja04_03() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(x[:5])
	fmt.Println(x[5:])
	fmt.Println(x[2:7])
	fmt.Println(x[1:6])
	fmt.Println(x)

}

//Ninja04_04 Func
func Ninja04_04() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(x)
	x = append(x, 52)
	fmt.Println(x)
	x = append(x, 53, 54, 55)
	fmt.Println(x)
	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...)
	fmt.Println(x)

}

//Ninja04_05 Func
func Ninja04_05() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(x)
	y := append(x[:3], x[6:]...)
	fmt.Println(y)

}

//Ninja04_06 Func
func Ninja04_06() {
	x := make([]string, 50, 50)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	x = []string{` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`}
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	for i := 0; i < len(x); i++ {
		fmt.Println(i, x[i])
	}
}

//Ninja04_07 Func
func Ninja04_07() {
	jb := []string{"James", "Bond", "Shaken, not stirred"}
	mp := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	xp := [][]string{jb, mp}
	for k, v := range xp {
		fmt.Println(k)
		for j, val := range v {
			fmt.Println(j, val)

		}
	}
}

//Ninja04_08 Func
func Ninja04_08() {
	m := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}
	fmt.Println(m)
	for k, v := range m {
		fmt.Println(k)
		for j, val := range v {
			fmt.Println("\t", j, val)
		}
	}

	// fmt.Println(m)
	// fmt.Println(m["James"])
	// fmt.Println(m["NONESISTO"]) // se richiedo una chiave non esistente, il valore di ritorno e' zero
	// // Per controllare se esiste una chiave, si usa l'idioma "Comma / , ok idiom"
	// if v, ok := m["NONESISTO"]; ok {
	// 	fmt.Println(v)
	// } else {
	// 	fmt.Println("Chiave non esistente")
	// }
}

//Ninja04_09 Func
func Ninja04_09() {
	m := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}
	m[`fleming_ian`] = []string{`steaks`, `cigars`, `espionage`}
	fmt.Println(m)
	for k, v := range m {
		fmt.Println(k)
		for j, val := range v {
			fmt.Println("\t", j, val)
		}
	}

	// fmt.Println(m)
	// fmt.Println(m["James"])
	// fmt.Println(m["NONESISTO"]) // se richiedo una chiave non esistente, il valore di ritorno e' zero
	// // Per controllare se esiste una chiave, si usa l'idioma "Comma / , ok idiom"
	// if v, ok := m["NONESISTO"]; ok {
	// 	fmt.Println(v)
	// } else {
	// 	fmt.Println("Chiave non esistente")
	// }
}

//Ninja04_10 Func
func Ninja04_10() {
	m := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}
	m[`fleming_ian`] = []string{`steaks`, `cigars`, `espionage`}
	fmt.Println(m)
	if _, ok := m["bond_james"]; ok {
		delete(m, "bond_james")
		//fmt.Println(m)
		for k, v := range m {
			fmt.Println(k)
			for j, val := range v {
				fmt.Println("\t", j, val)
			}
		}
	} else {
		fmt.Println("Chiave non esistente")
	}

}
