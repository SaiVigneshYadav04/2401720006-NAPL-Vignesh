package lab2

import (
	"fmt"
	"slices"
)

func SliceQuestion() {
	fmt.Println("Created a Slice")
	s := []int{80, 87, 92}
	fmt.Println(s)

	fmt.Println("Slice before appending")
	fmt.Println(s)
	s = append(s, 20)
	fmt.Println("Slice after appending")
	fmt.Println(s)

	fmt.Println("Slice before deleting")
	fmt.Println(s)
	s = slices.Delete(s, 1, 2)
	fmt.Println("Slice after deleting")
	fmt.Println(s)

	fmt.Println("Slice before updating")
	fmt.Println(s)
	s[0] = 1
	fmt.Println("Slice after updating")
	fmt.Println(s)

}
