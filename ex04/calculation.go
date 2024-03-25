package main

import (
	"fmt"
	"os"
	"strconv"
)

const ERROR_MSG string = "Arguments is invalid."

func calculationStr(args []string) (string, bool) {
	if len(args) != 3 {
		return "", false
	}
	num1, err1 := strconv.Atoi(args[1])
	num2, err2 := strconv.Atoi(args[2])

	if err1 != nil || err2 != nil {
		return "", false
	}
	sum := num1 + num2
	differrence := num1 - num2
	product := num1 * num2
	if num2 == 0 {
		return "", false
	}
	quotient := num1 / num2
	result := fmt.Sprintf("sum: %d\ndifference: %d\nproduct: %d\nquotient: %d\n", sum, differrence, product, quotient)
	return result, true
}

func main() {
	s, ok := calculationStr(os.Args)
	if !ok {
		fmt.Println(ERROR_MSG)
		os.Exit(1)
	}
	fmt.Print(s)
}
