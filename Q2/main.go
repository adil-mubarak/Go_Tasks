
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Question 2")

	var Number1 int
	var Number2 float64

	fmt.Print("Enter an integer: ")
	_, err := fmt.Scan(&Number1)
	if err != nil {
		fmt.Println("Error reading integer:", err)
		return
	}

	fmt.Print("Enter a float: ")
	_, err = fmt.Scan(&Number2)
	if err != nil {
		fmt.Println("Error reading float:", err)
		return
	}

	sum := float64(Number1) + Number2

	fmt.Printf("The sum is: %.2f\n", sum)
}



