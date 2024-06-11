package main

import "fmt"

func main() {
	fmt.Println("Quistion 21")

	var limit int

	fmt.Print("Enter the array limit: ")
	_, err := fmt.Scan(&limit)
	if err != nil {
		fmt.Println("Enter an integer number.")
	}

	array := make([]int, limit)
	fmt.Print("Enter the values of array: ")
	for i := 0; i < limit; i++ {
		_, err := fmt.Scan(&array[i])
		if err != nil {
			fmt.Println("Enter an integer numbers.")
		}
	}

	result := make([]int, len(array)-1)

	for i := 0; i < len(array)-1; i++ {
		result[i] = array[i] * array[i+1]
	}

	fmt.Println(result)

}
