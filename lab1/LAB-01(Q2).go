package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadInteger(reader *bufio.Reader, message string) int {
	for {
		fmt.Print(message)
		input, _ := reader.ReadString('\n')
		num, err := strconv.Atoi(strings.TrimSpace(input))

		if err == nil && num >= -100000 && num <= 100000 {
			return num
		} else {
			fmt.Println("Invalid input! Enter an integer between -100000 and 100000.")
		}
	}
}

func ReadFloat(reader *bufio.Reader, message string) float64 {
	for {
		fmt.Print(message)
		input, _ := reader.ReadString('\n')
		num, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

		if err == nil && num >= -100000 && num <= 100000 {
			return num
		} else {
			fmt.Println("Invalid input! Enter a number between -100000 and 100000.")
		}
	}
}

func CalculatorQuestion() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Integer Operations")

	a := ReadInteger(reader, "Enter first integer: ")
	b := ReadInteger(reader, "Enter second integer: ")

	fmt.Println("Addition:", a+b)
	fmt.Println("Subtraction:", a-b)
	fmt.Println("Multiplication:", a*b)

	fmt.Println("\nFloating Point Operations")

	x := ReadFloat(reader, "Enter first float: ")
	y := ReadFloat(reader, "Enter second float: ")

	fmt.Printf("Addition: %.2f\n", x+y)
	fmt.Printf("Subtraction: %.2f\n", x-y)
	fmt.Printf("Multiplication: %.2f\n", x*y)
}

func main() {
	CalculatorQuestion()
}
