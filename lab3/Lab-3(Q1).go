package go
package main

import "fmt"

func changeValue(num *int) {
	*num = 100
}

type Student struct {
	Name string
	Age  int
}

func main() {

	x := 50

	fmt.Println("Value of x:", x)
	fmt.Println("Address of x:", &x)
	fmt.Println("Value using *:", *(&x))

	fmt.Println("Before function:", x)

	changeValue(&x)

	fmt.Println("After function:", x)
	s := new(Student)

	s.Name = "Vignesh"
	s.Age = 20

	fmt.Println("Student Name:", s.Name)
	fmt.Println("Student Age:", s.Age)
}