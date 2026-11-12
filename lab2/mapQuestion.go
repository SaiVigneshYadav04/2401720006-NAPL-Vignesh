package lab2

import "fmt"

func MapQuestion() {
	fmt.Println("Created a Map")
	subjects := map[string]int{}
	subjects["Java"] = 90
	subjects["Python"] = 85
	subjects["SQL"] = 70
	fmt.Println(subjects)

	fmt.Println("Map before appending")
	fmt.Println(subjects)
	subjects["English"] = 95
	fmt.Println("Map after appending")
	fmt.Println(subjects)

	fmt.Println("Map before deleting")
	fmt.Println(subjects)
	delete(subjects, "Python")
	fmt.Println("Map after deleting")
	fmt.Println(subjects)

	fmt.Println("Lookup")
	for k, v := range subjects {
		fmt.Println(k, v)
	}

	fmt.Println("Trying to look up an element which is not present")
	v, status := subjects["Physics"]
	fmt.Println(status, status, v)

}
