package main

import "fmt"

func main() {
	fmt.Println("Quistion 19")

	var income float64
	var taxamount float64

	fmt.Print("Enter the annual income: ")
	fmt.Scan(&income)

	if income <= 25000 {
		fmt.Println("No TAX")
	} else if income > 25000 && income <= 500000 {
		taxamount = income * 0.05
	} else if income > 500000 && income <= 1000000 {
		taxamount = income * 0.20
	} else if income > 1000000 && income <= 5000000 {
		taxamount = income * 0.30
	}

	fmt.Printf("Income Tax amount: %.2f\n", taxamount)

}
