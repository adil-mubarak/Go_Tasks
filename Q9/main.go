package main

import "fmt"

func main() {
	fmt.Println("Quistion 9")

	for i := 1; i <= 5; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(j)
		}
		fmt.Println()

	}
	

}

