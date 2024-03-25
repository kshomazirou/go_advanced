package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	noNewline := flag.Bool("n", false, "omit trailing newline")
	separator := flag.String("s", " ", "separator")

	flag.Parse()

	args := flag.Args()

	result := strings.Join(args, *separator)

	if !*noNewline {
		result += "\n"
	}

	fmt.Print(result)

}
