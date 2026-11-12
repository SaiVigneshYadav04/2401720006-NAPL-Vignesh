package main

import "fmt"

func CalculatorQuestion() {
	for {
		var choice int

		fmt.Println("\n1. Integer Operations")
		fmt.Println("2. Float Operations")
		fmt.Println("3. Exit")
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		if choice == 1 {
			var a, b int

			fmt.Print("Enter first integer: ")
			fmt.Scan(&a)

			fmt.Print("Enter second integer: ")
			fmt.Scan(&b)

			fmt.Println("Addition:", a+b)
			fmt.Println("Subtraction:", a-b)
			fmt.Println("Multiplication:", a*b)

		} else if choice == 2 {
			var x, y float32

			fmt.Print("Enter first float: ")
			fmt.Scan(&x)

			fmt.Print("Enter second float: ")
			fmt.Scan(&y)

			fmt.Printf("Addition: %f\n", x+y)
			fmt.Printf("Subtraction: %f\n", x-y)
			fmt.Printf("Multiplication: %f\n", x*y)

		} else if choice == 3 {
			break
		} else {
			fmt.Println("Invalid choice!")
		}
	}
}

func main() {
	CalculatorQuestion()
}
