package main

import (
	"fmt"
)

func main() {
	fmt.Println(RepeatAlpha("abc"))
	fmt.Println(RepeatAlpha("Choumi."))
	fmt.Println(RepeatAlpha(""))
	fmt.Println(RepeatAlpha("abacadaba 01!"))
}

func RepeatAlpha(s string) string {
	var result string
	for i := 0; i < len(s); i++ {
		if s[i] >= 65 && s[i] <= 90 {
			for j := 0; j < int(s[i])-64; j++ {
				result += string(s[i])
			}
		} else if s[i] >= 97 && s[i] <= 122 {
			for j := 0; j < int(s[i])-96; j++ {
				result += string(s[i])
			}
		} else {
			result += string(s[i])
		}
	}

	return result
}
