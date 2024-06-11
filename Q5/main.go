package main

import "fmt"

func main() {
	fmt.Println("Quistion 5")

	fmt.Print("Enter the mark: ")
	var Mark float64
	fmt.Scan(&Mark)

	if Mark <= 100 {
		fmt.Println("The grade is: A")
	} else if Mark >= 80 && Mark < 90 {
		fmt.Println("The grade is: B")
	} else if Mark >= 70 && Mark < 80 {
		fmt.Println("The grade is: C")
	} else if Mark >= 60 && Mark < 70 {
		fmt.Println("The grade is: D")
	} else if Mark >= 50 && Mark < 60 {
		fmt.Println("The grade is: E")
	} else if Mark <= 50{
		fmt.Println("Failed")
	}else{
		fmt.Println("Mark is outof 100")
	}
}
