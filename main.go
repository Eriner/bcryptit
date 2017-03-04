package main

import (
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	args := os.Args[1:]
	for _, arg := range args {
		hash, err := bcrypt.GenerateFromPassword([]byte(arg), bcrypt.DefaultCost)
		if err != nil {
			fmt.Errorf("unable to bcrypt: " + err.Error())
		}
		fmt.Println(arg + ": " + string(hash))
	}
}
