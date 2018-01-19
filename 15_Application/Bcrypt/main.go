package main

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := `MyS3creTPw`
	fmt.Println(time.Now())
	pw, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	fmt.Println(time.Now())
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(pw))
	}

	loginPw := `MyS3creTPw1`
	if err := bcrypt.CompareHashAndPassword(pw, []byte(loginPw)); err != nil {
		fmt.Println("pass does not match")
	} else {
		fmt.Println("Pass match!")
	}
}
