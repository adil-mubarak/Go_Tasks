package main

import "fmt"

func main() {
	fmt.Println("Quistion 7")

	fmt.Println("Enter the number: ")
	var number int
	fmt.Scan(&number)

	for i := 1; i <= 10; i++ {
		fmt.Printf(" %v x %v = %v\n",i,number,i*number)
	}

}
