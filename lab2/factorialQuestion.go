package lab2

func FactorialQuestion(a int) int {
	if a == 0 {
		return 1
	}
	if a == 1 {
		return 1
	}
	return a * FactorialQuestion(a-1)
}
