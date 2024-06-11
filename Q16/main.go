package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Quistion 16")

	var num int
	fmt.Print("Enter a number: ")
	_, err := fmt.Scan(&num)
	if err != nil {
		fmt.Println("Please enter an integer number.")
	}

	if prime(num) {
		fmt.Println("Enterd number is Prime number.")
	} else {
		fmt.Println("Enterd number is not a Prime number.")
	}
}

func prime(n int) bool {
	if n <= 1 {
		return false
	}

	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%2 == 0 {
			return false
		}
	}
	return true
}
