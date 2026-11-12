package main

import (
	"fmt"
	"time"
)

func square(n int, ch chan string) {
	fmt.Println("Square goroutine started")
	time.Sleep(100 * time.Millisecond)
	ch <- fmt.Sprintf("Square: %d", n*n)
	fmt.Println("Square goroutine completed")
}

func cube(n int, ch chan string) {
	fmt.Println("Cube goroutine started")
	time.Sleep(150 * time.Millisecond)
	ch <- fmt.Sprintf("Cube: %d", n*n*n)
	fmt.Println("Cube goroutine completed")
}

func fibonacci(n int, ch chan string) {
	fmt.Println("Fibonacci goroutine started")
	time.Sleep(50 * time.Millisecond)
	a, b := 0, 1
	for i := 0; i < n; i++ {
		a, b = b, a+b
	}
	ch <- fmt.Sprintf("Fibonacci: %d", a)
	fmt.Println("Fibonacci goroutine completed")
}

func main() {
	ch := make(chan string, 3)
	go square(5, ch)
	go cube(5, ch)
	go fibonacci(10, ch)

	for i := 0; i < 3; i++ {
		fmt.Println(<-ch)
	}
	close(ch)
}
