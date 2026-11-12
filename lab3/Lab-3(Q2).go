package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Person struct {
	Name   string
	Age    int
	Job    string
	Salary float64
}

func (p Person) readData(reader *bufio.Reader) Person {
	fmt.Print("Enter Name of the person: ")
	p.Name, _ = reader.ReadString('\n')
	p.Name = strings.TrimSpace(p.Name)

	fmt.Print("Enter Age of the person: ")
	age, _ := reader.ReadString('\n')
	p.Age, _ = strconv.Atoi(strings.TrimSpace(age))

	fmt.Print("Enter Job of the person: ")
	p.Job, _ = reader.ReadString('\n')
	p.Job = strings.TrimSpace(p.Job)

	fmt.Print("Enter Salary of the person: ")
	salary, _ := reader.ReadString('\n')
	p.Salary, _ = strconv.ParseFloat(strings.TrimSpace(salary), 64)

	return p
}

func (p Person) printData() {
	fmt.Printf("Name: %v\n", p.Name)
	fmt.Printf("Age: %v\n", p.Age)
	fmt.Printf("Job: %v\n", p.Job)
	fmt.Printf("Salary: %.2f\n", p.Salary)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var p1, p2 Person

	fmt.Println("Enter details of Person 1:")
	p1 = p1.readData(reader)
	p1.printData()

	fmt.Println("\nEnter details of Person 2:")
	p2 = p2.readData(reader)
	p2.printData()
}
