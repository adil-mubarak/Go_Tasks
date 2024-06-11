package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Quistion 12")

	var size int
	fmt.Print("Enter the size of the array: ")
	_, err := fmt.Scan(&size)
	if err != nil {
		fmt.Println("Please enter an integer value.")
		return
	}

	array := make([]int, size)
	fmt.Print("Enter the values of array: ")
	for i := 0; i < size; i++ {
		_, err = fmt.Scan(&array[i])
		if err != nil {
			fmt.Println("Please enter an integer values.")
			return
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(array)))

	fmt.Println("The array is: ", array)

}
