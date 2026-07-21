package main

import "fmt"

func main() {
	var name string = "Naren Roy"
	fmt.Println("Hello, ", name)

	// Short hand declaration
	age := 25
	fmt.Println("Age: ", age)

	// Multiple variable declaration
	var (
		city    string = "New York"
		country string = "USA"
	)
	fmt.Println("City: ", city)
	fmt.Println("Country: ", country)

	// Constant declaration
	const pi float64 = 3.14159
	fmt.Println("Value of Pi: ", pi)

	// Zero value of variables
	var isActive bool
	var score int
	var price float64
	fmt.Println("Is Active: ", isActive)
	fmt.Println("Score: ", score)
	fmt.Println("Price: ", price)

	// Type conversion
	var num1 int = 10
	var num2 float64 = 3.14
	var result float64 = float64(num1) + num2
	fmt.Println("Result: ", result)

	// Using the blank identifier
	var unusedVar int = 42
	_ = unusedVar // Ignoring the unused variable

}
