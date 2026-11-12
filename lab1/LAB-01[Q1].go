package main

import "fmt"

func BasicQuestion() {
	fmt.Println("Hello, Vignesh")
}

func CalculatorQuestion() {
	var a, b int = 20, 10
	var x, y float64 = 15.5, 5.5

	fmt.Println("\nInteger Operations:")
	fmt.Println("Addition:", a+b)
	fmt.Println("Subtraction:", a-b)
	fmt.Println("Multiplication:", a*b)

	fmt.Println("\nFloating Point Operations:")
	fmt.Println("Addition:", x+y)
	fmt.Println("Subtraction:", x-y)
	fmt.Println("Multiplication:", x*y)
}

func main() {
	BasicQuestion()
	CalculatorQuestion()
}

