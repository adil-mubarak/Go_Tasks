package main

import (
	"fmt"
)

type Array struct {
	size int
	arr  [][]int
}

func (a *Array) getArray() {
	a.arr = make([][]int, a.size)
	fmt.Println("Enter the values of array:")
	for i := 0; i < a.size; i++ {
		a.arr[i] = make([]int, a.size)
		for j := 0; j < a.size; j++ {
			fmt.Scan(&a.arr[i][j])
		}
	}
}

func (a *Array) Displayarray() {
	fmt.Println("Array elements are:")
	for i := 0; i < a.size; i++ {
		for j := 0; j < a.size; j++ {
			fmt.Printf("%d ", a.arr[i][j])
		}
		fmt.Println()
	}
}

func main() {
	fmt.Println("Quistion 23")
	var size int
	fmt.Print("Enter the size of the array: ")
	_, err := fmt.Scan(&size)
	if err != nil {
		fmt.Println("Please enter an integer number.")
	}

	array := Array{size: size}

	array.getArray()
	array.Displayarray()
}
