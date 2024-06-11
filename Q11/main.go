package main

import "fmt"

func main() {
	fmt.Println("Quistion 11")

	var size int
	fmt.Print("Enter the size of the array: ")
	_, err := fmt.Scan(&size)
	if err != nil {
		fmt.Println("Please enter an integer value.")
		return
	}

	array := make([]int, size)
	fmt.Print("Enter the array values: ")
	for i := 0; i < size; i++ {
		_, err := fmt.Scan(&array[i])
		if err != nil {
			fmt.Println("Please enter an integer value.")
			return
		}
	}

	evennumlen := 0
	var evennum []int
	for _, num := range array {
		if num%2 == 0 {
			evennum = append(evennum, num)
			evennumlen++
		}
	}

	fmt.Printf("array: %v \n", array)
	fmt.Println("The even numbers: ", evennum)
	fmt.Println("Number of even numbers in the given array: ", evennumlen)
}
