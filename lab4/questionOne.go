package lab4

import (
	"fmt"
	"time"
)

func squareWorker(num int, result chan<- int) {
	fmt.Println("Square Routine started")
	result <- num * num
	fmt.Println("Square Routine completed")
}

func cubeWorker(num int, result chan<- int) {
	fmt.Println("Cube Routine started")
	result <- num * num * num
	fmt.Println("Cube Routine completed")

}

func fibonacciWorker(num int, result chan<- int) {
	fmt.Println("Fibonacci Routine started")
	a, b := 0, 1
	for i := 0; i < num; i++ {
		a, b = b, a+b
	}
	result <- a
	fmt.Println("Fibonacci Routine completed")
}

func QuestionOne() {
	result := make(chan int)

	var input int
	fmt.Print("Enter an Integer ")
	fmt.Scan(&input)
	go squareWorker(input, result)
	go cubeWorker(input, result)
	go fibonacciWorker(input, result)

	for i := 1; i <= 3; i++ {
		fmt.Println(<-result)
		time.Sleep(time.Millisecond * 500)
	}
}
