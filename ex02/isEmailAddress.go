package main

import (
	"fmt"
	"os"
	"regexp"
)

func isValidEmail(email string) bool {
	emailPatternn := `[\w\-\._]+@[\w\-\._]+\.[A-Za-z]+`
	match, err := regexp.MatchString(emailPattern, email)
	if err != nil {
		return false
	}
	return match
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("Invalid argument")
		return
	}
	for _, arg := range args {
		if isValidEmail(arg) {
			fmt.Printf("%s is a valid email address.\n", arg)
		} else {
			fmt.Printf("%s is NOT a valid email address.\n", arg)
		}
	}
}
