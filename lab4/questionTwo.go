package lab4

import "errors"

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func SquareRoot(n int) (int, error) {
	if n < 0 {
		return 0, errors.New("negative number")
	}
	for i := 0; i*i <= n; i++ {
		if i*i == n {
			return i, nil
		}
	}
	return 0, errors.New("not a perfect square")
}
