package main

import (
	"fmt"
)

func main() {
	fmt.Println("Quistion 14")

	var size int
	fmt.Print("Enter the size of arrays: ")
	fmt.Scan(&size)

	array1 := make([][]int, size)
	array2 := make([][]int, size)
	sum := make([][]int, size)


	fmt.Println("Enter the values for array1: ")
	for i := 0; i < size; i++ {
		array1[i] = make([]int, size)
		for j := 0; j < size; j++ {
			fmt.Scan(&array1[i][j])
		}
	}

	fmt.Println("Enter the values for array2: ")
	for i := 0; i < size; i++ {
		array2[i] = make([]int,size)
		for j := 0; j < size; j++ {
			fmt.Scan(&array2[i][j])
		}
	}

	for i := 0; i < size; i++ {
		sum[i] = make([]int, size)
		for j := 0; j < size; j++ {
			sum[i][j] = array1[i][j] + array2[i][j]
		}
	}

	

	fmt.Println("Sum of 2 arrays is:")
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			fmt.Printf("%d ",sum[i][j])
		}
		fmt.Println()
	}


}
