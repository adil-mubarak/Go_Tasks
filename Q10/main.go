package main

import "fmt"

func main() {
	fmt.Println("Quistion 10")

	fmt.Print("Enter the size of the array: ")
	var size int
	_, err := fmt.Scan(&size)
	if err != nil {
		fmt.Println("Please enter an integer value.")
		return
	}

	array1 := make([]int, size)
	array2 := make([]int, size)

	readArray := func(arraynum int, array []int) {
		fmt.Printf("Enter the values of array %d: ", arraynum)
		for i := 0; i < size; i++ {
			_, err := fmt.Scan(&array[i])
			if err != nil {
				fmt.Println("Please enter an integer values.")
				return
			}
		}
	}
	readArray(1, array1)
	readArray(2, array2)

	array1, array2 = array2, array1

	fmt.Println("The arrays after swapping")
	fmt.Printf("array1: %v\n ", array1)
	fmt.Printf("array2: %v\n ", array2)
}
