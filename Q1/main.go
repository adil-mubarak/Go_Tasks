package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Quistion 1")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the Character: ")

	input, _ := reader.ReadString('\n')
	fmt.Println("You enterd character is: ", input)
}
