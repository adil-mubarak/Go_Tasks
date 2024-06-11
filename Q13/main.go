package main

import (
	"fmt"
	"strings"
)

func reverse(s []string) []string {
	for i := 0; i < len(s)/2; i++ {
		j := len(s) - i - 1
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func main() {
	fmt.Println("Quistion 13")

	var sent string
	fmt.Print("Enter a string: ")
	fmt.Scan(&sent)

	spl := strings.Split(sent, "")
	reverse(spl)
	rev := strings.Join(spl, "")

	if sent == rev {
		fmt.Println("Enterd string is a Palindrome.")
	} else {
		fmt.Println("Enterd string is not a Palindrome.")
	}

}
