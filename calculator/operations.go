package calculator

import (
	"errors"
	"math"
)

func Add(a, b float64) float64 {
	return a + b
}

func Sub(a, b float64) float64 {
	return a - b
}

func Multi(a, b float64) float64 {
	return a * b
}

func Div(a, b float64) (float64, int, error) {
	if b == 0 {
		return 0, 0, errors.New("Math Error: Can't divide by 0")
	}
	decAns := a / b
	remainder := int(a) % int(b)
	return decAns, remainder, nil
}

func Power(a, b float64) float64 {
	return math.Pow(a, b)
}
