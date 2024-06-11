package main

import (
	"fmt"
)

func main() {
	fmt.Println("Quistion 3")

	var P int
	var R float64
	var n float64
	var SI float64

	fmt.Print("Enter the Principal amount: ")
	_, err := fmt.Scan(&P)
	if err != nil {
		fmt.Println("Error reading Principle amount:", err)
		return
	}

	fmt.Print("Enter the Interest rate: ")
	_, err = fmt.Scan(&R)
	if err != nil {
		fmt.Println("Error reading the Interest rate:", err)
		return
	}

	fmt.Print("Enter the Number of year: ")
	_, err = fmt.Scan(&n)
	if err != nil {
		fmt.Println("Error reading the Number of year:", err)
		return
	}

	SI = (float64(P) * R * n / 100)
	fmt.Printf("The Simple Interset is: %.2f \n", SI)

}
