package lab2

import "fmt"

func ReverseQuestion(str string) {
	rev := ""

	for i := len(str) - 1; i >= 0; i-- {
		rev = rev + string(str[i])
	}

	fmt.Println(rev)
}
