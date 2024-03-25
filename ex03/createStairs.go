package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) > 1 {
		args := os.Args[1]
		num, err := strconv.Atoi(args)
		if err != nil {
			fmt.Println("Invalid format!")
		}
		if num <= 0 || num >= 10000 {
			return
		}
		count := 1
		for i := 1; i <= num && (((2 + (i - 1))/ 2 * i) <= num); i++ {
		fmt.Println(((2 + (i - 1))/ 2 * i))
			for j := 0; j < i; j++ {
				fmt.Printf("*")
				count++
			}
			fmt.Println()
		}
	} else {
		fmt.Println("No command line!")
	}
}
