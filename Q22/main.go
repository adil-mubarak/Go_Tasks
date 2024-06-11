package main

import "fmt"

func Getarray(size int, num int) [][]int {
	array := make([][]int, size)
	fmt.Printf("Enter the values of array %d:\n", num)
	for i := 0; i < size; i++ {
		array[i] = make([]int, size)
		for j := 0; j < size; j++ {
			fmt.Scan(&array[i][j])
		}
	}
	return array
}

func Addarray(array1, array2 [][]int) [][]int {
	size := len(array1)
	add := make([][]int, size)
	for i := 0; i < size; i++ {
		add[i] = make([]int, size)
		for j := 0; j < size; j++ {
			add[i][j] = array1[i][j] + array2[i][j]
		}
	}
	return add
}

func Displayarray(array [][]int) {
	fmt.Println("Sum of array1 and array2:")
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array[i]); j++ {
			fmt.Printf("%d ", array[i][j])
		}
		fmt.Println()
	}
}

func main() {
	fmt.Println("Quistion 22")

	var size int
	fmt.Print("Enter the size of array:")
	_, err := fmt.Scan(&size)
	if err != nil {
		fmt.Println("Please enter an integer number.")
	}

	array1 := Getarray(size, 1)
	array2 := Getarray(size, 2)

	sumarray := Addarray(array1, array2)

	Displayarray(sumarray)

}
