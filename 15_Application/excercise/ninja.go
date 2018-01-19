//Package excercise - 15 Application - Ninja01
package excercise

import (
	"encoding/json"
	"fmt"
	"os"
)

type user struct {
	First string
	Age   int
}

type pex02 struct {
	First   string   `json:"First"`
	Last    string   `json:"Last"`
	Age     uint8    `json:"Age"`
	Sayings []string `json:"Sayings"`
}

func Ex01() {
	u1 := user{
		First: "James",
		Age:   32,
	}

	u2 := user{
		First: "Moneypenny",
		Age:   27,
	}

	u3 := user{
		First: "M",
		Age:   54,
	}

	users := []user{u1, u2, u3}

	fmt.Println("Slice of users: ", users)

	bs, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bs))
	bspretty, _ := json.MarshalIndent(users, "", "    ")
	fmt.Println(string(bspretty))

}

func Ex02() {
	s := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`
	fmt.Println(s)
	var pex []pex02
	json.Unmarshal([]byte(s), &pex)
	for _, v := range pex {
		fmt.Println(v.First, v.Last, v.Age)
		for _, v1 := range v.Sayings {
			fmt.Println(v1)

		}
	}
}

func Ex03() {
	fmt.Printf("\n")
	fmt.Printf("\n")
	fmt.Printf("\n")
	u1 := pex02{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := pex02{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := pex02{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []pex02{u1, u2, u3}

	fmt.Println(users)
	fmt.Printf("\n")
	fmt.Printf("\n")
	fmt.Printf("\n")
	if err := json.NewEncoder(os.Stdout).Encode(&users); err != nil {
		fmt.Println("Something wrong happens, here's the error:", err)
	}
}
