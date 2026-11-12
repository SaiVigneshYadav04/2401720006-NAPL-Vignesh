package lab2

func PowerQuestion(a, b int) int {
	if a == 0 {
		return 1
	}
	if b == 0 {
		return 1
	}
	result := 1

	for i := 1; i <= b; i++ {
		result = result * a
	}

	return result
}
