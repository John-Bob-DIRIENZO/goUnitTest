package maths

func Add(a, b int) int {
	return a + b
}

func Substract(a, b int) int {
	return a - b
}

func Multiply(a, b int) int {
	return a * b
}

func Divide(a, b int) float64 {
	if b == 0 {
		return 0
	}

	return float64(a / b)
}
