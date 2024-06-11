package main

import (
	"fmt"
)

func main() {
	fmt.Println("Quistion 18")

	var Writtentest float64
	var Labexam float64
	var Assignments float64
	var overall float64

	fmt.Print("Written test: ")
	fmt.Scan(&Writtentest)
	fmt.Print("Lab exam: ")
	fmt.Scan(&Labexam)
	fmt.Print("Assignments: ")
	fmt.Scan(&Assignments)

	overall = (Writtentest*70)/100 + (Labexam*20)/100 + (Assignments*10)/100
	fmt.Printf("Grade of the student is: %.2f\n", overall)

}
