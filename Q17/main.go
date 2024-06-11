package main

import "fmt"

type Calculation struct{}

func (c Calculation) addition(a, b float64) float64 {
	return a + b
}
func (c Calculation) substraction(a, b float64) float64 {
	return a - b
}

func (c Calculation) multiplication(a, b float64) float64 {
	return a * b
}

func (c Calculation) division(a, b float64) float64 {
	if b != 0 {
		return a / b
	}
	fmt.Println("Error")
	return 0
}

func main() {
	fmt.Println("Quistion 17")

	var choose int
	var a, b float64
	Calculation := Calculation{}

	for {
		fmt.Println("Choose the operator")
		fmt.Println("1: ADDITION")
		fmt.Println("2: SUBSTRACTION")
		fmt.Println("3: MULTIPLICATION")
		fmt.Println("4: DIVISION")
		fmt.Println("5: EXIT")

		fmt.Print("Enter the choice: ")
		fmt.Scan(&choose)

		if choose == 5 {
			break
		}

		fmt.Print("Enter the First number: ")
		fmt.Scan(&a)
		fmt.Print("Enter the Second number: ")
		fmt.Scan(&b)

		switch choose {
		case 1:
			result := Calculation.addition(a, b)
			fmt.Printf("The calculated value is: %.2f \n", result)
		case 2:
			result := Calculation.substraction(a, b)
			fmt.Printf("The calculated value is: %.2f\n", result)
		case 3:
			result := Calculation.multiplication(a, b)
			fmt.Printf("The calculated value is: %.2f\n", result)
		case 4:
			result := Calculation.division(a, b)
			fmt.Printf("The calculated value is: %.2f\n", result)
		default:
			fmt.Println("Invalid number.")
		}
	}
	fmt.Println("The calculation is ending thanks...")
}
