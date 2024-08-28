package main

import (
	"fmt"
	"os"
)

func main() {
	result := ""
	args := os.Args[1]
	if len(os.Args) > 2 {
		fmt.Println()
	}

	for i := 0; i < len(args); i++ {
		if string(args[i]) != " " {
			result += string(args[i])
		}
		if i != 0 && i < len(args)-1 && (string(args[i]) == " " && string(args[i+1]) != " ") {
			result += " "
		}

	}
	fmt.Println(result)
}
// hjjhjhjj

