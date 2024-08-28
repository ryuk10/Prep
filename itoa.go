package main

import (
	"fmt"
)

func main() {
	fmt.Println(Itoa(12345))
	fmt.Println(Itoa(0))
	fmt.Println(Itoa(-1234))
	fmt.Println(Itoa(987654321))
}

func Itoa(n int) string {
	negative := ""
	result := ""
	sli := []string{}
	if n < 0 {
		negative = "-"
		n = -n
	}
	for n/10 != 0 {
		sli = append(sli, string(n%10+'0'))
		n /= 10
	}
	sli = append(sli, string(n+'0'))

	for i := len(sli) - 1; i >= 0; i-- {
		result += sli[i]
	}
	return negative + result
}
