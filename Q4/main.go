package main

import "fmt"

func main() {
	fmt.Println("Quiston 4")
	 

	fmt.Print("Enter the mark: ")
	var mark float64

	fmt.Scan(&mark)
	if mark >= 50 && mark <= 100{
		fmt.Println("Passed")
	}else{
		fmt.Println("Failed")
	}

}
