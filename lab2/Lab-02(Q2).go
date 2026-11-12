package main

import "fmt"

func printMap(m map[string]int) {
	fmt.Print("Map: ")
	fmt.Print("map[")
	keys := []string{"Math", "Science", "English"}

	first := true
	for _, key := range keys {
		if value, ok := m[key]; ok {
			if !first {
				fmt.Print(" ")
			}
			fmt.Printf("%s:%d", key, value)
			first = false
		}
	}
	fmt.Println("]")
}

func main() {
	names := []string{"Rahul", "Aman", "Vikas"}

	fmt.Println("Initial Slice:", names)

	names = append(names, "Karan")
	fmt.Println("After Adding:", names)

	index := 1
	names = append(names[:index], names[index+1:]...)
	fmt.Println("After Removing Index 1:", names)

	names[1] = "Arjun"
	fmt.Println("After Updating Index 1:", names)

	fmt.Println()

	marks := make(map[string]int)

	marks["Math"] = 90
	fmt.Println("After Inserting Math:")
	printMap(marks)

	marks["Science"] = 85
	fmt.Println("After Inserting Science:")
	printMap(marks)

	marks["English"] = 88
	fmt.Println("After Inserting English:")
	printMap(marks)

	delete(marks, "Science")
	fmt.Println("After Deleting Science:")
	printMap(marks)

	fmt.Println("Lookup Math:", marks["Math"])
}
