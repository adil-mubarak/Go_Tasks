package main

import "fmt"

func main() {
	fmt.Println("Quistion 8")

	fmt.Print("Enter the number: ")
	var num int
	fmt.Scan(&num)

	sum := 0

	for i := 1 ; i <= num; i++ {
     if i%2!=0{
		sum+=i
	 }
		
	}
	fmt.Println("The sum of odd number is: ",sum)
	

}
