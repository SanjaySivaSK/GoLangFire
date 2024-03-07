// Golang program to show how to
// access the fields of struct
package main

import "fmt"

// defining the struct
type Car struct {
	Name, Model, Color string
	WeightInKg         float64
}

// Main Function
func main() {
	c := Car{Name: "Ferrari", Model: "GTC4",
		Color: "Red", WeightInKg: 1920}

	// Accessing struct fields
	// using the dot operator
	fmt.Println("Car Name: ", c.Name)
	fmt.Println("Car Color: ", c.Color)

	// Assigning a new value
	// to a struct field
	c.Color = "Black"

	// Displaying the result
	fmt.Println("Car: ", c)
	c.Color = "reddddd"

	fmt.Println(c)

	a := 7
	fmt.Println(a)
	a = 8
	fmt.Println(check(a))

	fmt.Println(a)

}

func check(a int) int {
	a = 3

	return a

}
