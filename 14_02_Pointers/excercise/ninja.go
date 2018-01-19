package excercise

import (
	"fmt"
)

type person2 struct {
	first      string
	last       string
	favFlavors []string
}

type vehicle struct {
	doors uint8
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

//Ninja05_01 play with structure
func Ninja05_01() {
	p1 := person2{
		first: "James",
		last:  "Bond",
		favFlavors: []string{
			"chocolate",
			"martini",
			"rum and coke",
		},
	}

	p2 := person2{
		first: "Miss",
		last:  "Moneypenny",
		favFlavors: []string{
			"strawberry",
			"vanilla",
			"capuccino",
		},
	}

	fmt.Println(p1.first)
	fmt.Println(p1.last)
	for i, v := range p1.favFlavors {
		fmt.Println(i, v)
	}

	fmt.Println(p2.first)
	fmt.Println(p2.last)
	for i, v := range p2.favFlavors {
		fmt.Println(i, v)
	}
}

//Ninja05_02 put a variable of your type person into a map
func Ninja05_02() {
	p1 := person2{
		first: "James",
		last:  "Bond",
		favFlavors: []string{
			"chocolate",
			"martini",
			"rum and coke",
		},
	}

	p2 := person2{
		first: "Miss",
		last:  "Moneypenny",
		favFlavors: []string{
			"strawberry",
			"vanilla",
			"capuccino",
		},
	}

	m := map[string]person2{
		p1.last: p1,
		p2.last: p2,
	}

	fmt.Println(m)

	for _, v := range m {
		fmt.Println(v.first)
		fmt.Println(v.last)
		for i, val := range v.favFlavors {
			fmt.Println(i, val)
		}
		fmt.Println("-------")
	}
}

//Ninja05_03 work with data structs
func Ninja05_03() {
	tr1 := truck{
		vehicle: vehicle{
			doors: 2,
			color: "Red",
		},
		fourWheel: true,
	}

	sd1 := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "Black",
		},
		luxury: false,
	}

	fmt.Println(tr1)
	fmt.Println(sd1)
	fmt.Println(tr1.fourWheel)
	fmt.Println(tr1.doors)
	fmt.Println(sd1.luxury)
	fmt.Println(sd1.doors)
}

//Ninja05_04 anon structs
func Ninja05_04() {
	x := struct {
		name      string
		last      string
		age       uint8
		interests []string
		enc       map[string]uint16
	}{
		name:      "XXX",
		last:      "YYY",
		age:       52,
		interests: []string{"Sex", "test"},
		enc: map[string]uint16{
			"RSA":   2048,
			"DSA":   1024,
			"ECDHA": 4096,
		},
	}

	fmt.Println(x.name)
	fmt.Println(x.last)
	fmt.Println(x.interests)

	for k, v := range x.enc {
		fmt.Println(k, v)
	}

	for i, val := range x.interests {
		fmt.Println(i, val)
	}
}
