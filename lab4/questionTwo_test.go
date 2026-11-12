package lab4

import (
	"testing"
)

func TestDivide(t *testing.T) {
	cases := []struct {
		a, b, want int
	}{
		{10, 2, 5},
		{9, 3, 3},
		{0, 2, 0},
	}

	for _, c := range cases {
		got, err := Divide(c.a, c.b)
		if err != nil {
			t.Errorf("Error %v", err)
		}
		if got != c.want {
			t.Errorf("Divide %v by %v = %v but wanted %v", c.a, c.b, got, c.want)
		}
	}
}

func TestDivideByZero(t *testing.T) {
	cases := []struct {
		a, b int
	}{
		{10, 0},
		{5, 0},
		{7, 0},
	}
	for _, c := range cases {
		_, err := Divide(c.a, c.b)
		if err == nil {
			t.Errorf("Error excpected for division by zero")
		}
	}
}

func TestSquareRoot(t *testing.T) {
	cases := []struct {
		n, want int
	}{
		{25, 5},
		{16, 4},
		{0, 0},
	}

	for _, c := range cases {
		got, err := SquareRoot(c.n)
		if err != nil {
			t.Errorf("Error %v", err)
		}
		if got != c.want {
			t.Errorf("Square Root of %v is %v, but wanted %v", c.n, got, c.want)
		}
	}
}

func TestNegativeSquareRoot(t *testing.T) {
	cases := []struct {
		n int
	}{
		{-9},
		{-4},
		{-1},
	}

	for _, c := range cases {
		_, err := SquareRoot(c.n)
		if err == nil {
			t.Errorf("Expected error for SquareRoot %v", c.n)
		}
	}
}
