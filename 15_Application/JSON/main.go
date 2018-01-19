package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"
)

type person struct {
	First string
	Last  string
	Age   uint8
}

type person1 struct {
	First string
	Age   uint8
}

type jsonPerson struct { // UNMARSHAL VIDEO
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

type jsonPerson1 struct { // UNMARSHAL VIDEO
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

// ByAge to use sort.Stable or sort.Sort interface
type ByAge []person1

func (ba ByAge) Len() int           { return len(ba) }
func (ba ByAge) Swap(i, j int)      { ba[i], ba[j] = ba[j], ba[i] }
func (ba ByAge) Less(i, j int) bool { return ba[i].Age < ba[j].Age }

type ByName []person1

func (bn ByName) Len() int           { return len(bn) }
func (bn ByName) Swap(i, j int)      { bn[i], bn[j] = bn[j], bn[i] }
func (bn ByName) Less(i, j int) bool { return bn[i].First < bn[j].First }

func main() {
	fmt.Println("")
	fmt.Println("Video 1 on Application section - JSON Documentation")

	fmt.Println("")
	fmt.Println("Video 2 on Application section - JSON marshal")
	p1 := person{
		First: "James",
		Last:  "Bond",
		Age:   42,
	}
	p2 := person{
		First: "Monay",
		Last:  "Penny",
		Age:   24,
	}

	people := []person{p1, p2} //Is not a best practice to put all in one line, but sometimes is more readable
	fmt.Println(people)

	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bs)) //Converto lo slice di bytes []bytes in string

	fmt.Println("")
	fmt.Println("Video 3 on Application  - JSON UNmarshal")

	s := `[{"First":"James","Last":"Bond","Age":42},{"First":"Monay","Last":"Penny","Age":24}]`
	fmt.Printf("Type of s: %T\n", s)
	byteslice := []byte(s)
	fmt.Printf("Type of byteslice: %T\n", byteslice)
	var newpeople []person
	//err = json.Unmarshal(byteslice, &newpeople)
	err = json.Unmarshal([]byte(s), &newpeople)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("\nAll the data:", newpeople)
	fmt.Printf("Type of newpeople: %T\n", newpeople)
	for i, v := range newpeople {
		fmt.Println("\nPERSON NUMBER: ", i+1)
		fmt.Println(v.First, v.Last, v.Age)
	}

	//json.Unmarshal(, &byteslice)

	//MY tests
	// var jp jsonPerson
	// json.Unmarshal(bs, &jp)
	// fmt.Println(jp)
	// fmt.Printf("Type of jp: %T\n", jp)
	fmt.Println("")
	fmt.Println(`Video 4 of "Application"  - JSON encode and decode + Write interface`)
	fmt.Println("Hello, playground")             //Println implementa Fprintln
	fmt.Fprintln(os.Stdout, "Hello, playground") //Fprintln ask work a io.Writer type, os.Stdout is a pointer to a file so it's implicitly of type Writer *p []bytes
	io.WriteString(os.Stdout, "Hello, playground")

	//My Tests
	// f, err := os.Create("/Users/lcancelliere/gocode/github.com/golang-code/McLeod_Go_Course/15_Application/JSON/test.txt")
	// defer f.Close()
	// if err != nil {
	// 	fmt.Println("Error in operning file: ", err)
	// }
	// fmt.Println(f.ReadAt(f, 1))
	// fmt.Printf("\n%T", f)

	fmt.Println("")
	fmt.Println("")
	fmt.Println(`Video 5 of "Application" - Sort`)
	xi := []int{4, 7, 3, 42, 99, 18, 16, 56, 12}
	xs := []string{"James", "Q", "M", "Moneypenny", "Dr. No"}
	fmt.Println(xi)
	fmt.Println(xs)
	sort.Strings(xs)
	sort.Ints(xi)
	fmt.Println(xs)
	fmt.Println(xi)

	fmt.Println("")
	fmt.Println("")
	fmt.Println(`Video 6 of "Application" - Sort Custom`)

	pp1 := person1{"James", 34}
	pp2 := person1{"Moneypenny", 27}
	pp3 := person1{"Q", 64}
	pp4 := person1{"M", 56}

	people2 := []person1{pp1, pp2, pp3, pp4}

	fmt.Println(people2)

	sort.Stable(ByAge(people2)) //STABLE E' LA VERSIONE STABILE DI sort.Sort()
	fmt.Println("Sorted by age: ", people2)
	sort.Stable(ByName(people2))
	fmt.Println("Sorted by name: ", people2)

}
