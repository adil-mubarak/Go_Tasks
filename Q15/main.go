package main

import "fmt"

func main() {
	fmt.Println("Quistion 15")
	var size int
	fmt.Print("Enter the size of the array: ")
	fmt.Scan(&size)
	array := make([]int, size)
	GetArray(array)
	DisplayArray(array)
}

func GetArray(array []int) {
	fmt.Println("Enter the values of array: ")
	for i := range array {
		fmt.Scan(&array[i])
	}
}

func DisplayArray(array []int) {
	fmt.Println("Array values are: ")
	for i, value := range array {
		fmt.Printf("Array[%d] = %d\n", i, value)
	}
}
