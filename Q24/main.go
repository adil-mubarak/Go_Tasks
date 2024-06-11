package main

import (
	"fmt"
	"math"
)

type Area struct{}

func (a Area) circle(radius float64) float64 {
	return math.Pi * radius * radius
}

func (a Area) square(length float64) float64 {
	return length * length
}

func (a Area) rectangle(length, width float64) float64 {
	return length * width
}

func (a Area) triangle(base, height float64) float64 {
	return 0.5 * base * height
}

type MyClass struct {
	Area
}

func (m MyClass) main() {
	var choice int
	fmt.Println("Choose your item.")
	fmt.Println("1: circle")
	fmt.Println("2: square")
	fmt.Println("3: rectangle")
	fmt.Println("4: triangle")
	fmt.Print("Enter your choice:")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		m.circle()
	case 2:
		m.square()
	case 3:
		m.rectangle()
	case 4:
		m.triangle()
	default:
		fmt.Println("invalid number")
	}
}

func(m MyClass) circle(){
	var radius float64
	fmt.Print("Enter the radius of the circle:")
	fmt.Scan(&radius)
	area := m.Area.circle(radius)
	fmt.Printf("The area of the circle: %.2f\n",area)
}

func(m MyClass) square(){
	var length float64
	fmt.Print("Enter the length of the square:")
	fmt.Scan(&length)
	area := m.Area.square(length)
	fmt.Printf("The areamof the square: %.2f\n",area)
}

func(m MyClass) rectangle(){
	var length,width float64
	fmt.Print("Enter the length and width of the rectangle:")
	fmt.Scan(&length,&width)
	area := m.Area.rectangle(length,width)
	fmt.Printf("The area of the rectangle: %.2f\n",area)
}

func (m MyClass) triangle(){
	var base,height float64
	fmt.Print("Enter the base and height of the triangle:")
	fmt.Scan(&base,&height)
	area := m.Area.triangle(base,height)
	fmt.Printf("The area of the triangle: %.2f",area)
}
func main() {
	fmt.Println("Quistion 24")
	myClass := MyClass{}
	myClass.main()

}
