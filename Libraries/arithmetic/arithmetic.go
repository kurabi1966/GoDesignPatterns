package arithmetic

import "errors"

// Sum function of the arithmetic package accepted zero to many integers and returns result as an integer.
// TODO: handle the overflow error
func Sum(args ...int) (res int) {
	for _, v := range args {
		res += v
	}
	return
}

// Subtract function of the arithmetic package accepted two to many integers and returns result as an integer. If number of Parameter is less than two, the function will return zero
func Subtract(args ...int) int {
	if len(args) < 2 {
		return 0
	}

	res := args[0]
	for i := 1; i < len(args); i++ {
		res -= args[i]
	}
	return res
}

// Multiply function of the arithmetic package accepted two to many integers and returns result as an integer. If number of Parameter is less than two, the function will return zero
func Multiply(args ...int) int {
	if len(args) < 2 {
		return 0
	}

	res := 1
	for i := 0; i < len(args); i++ {
		res *= args[i]
	}
	return res
}

// Divide function of the arithmetic package accepted two integers and returns result as an integer. If the devidor is zero, the function will return zero and error
func Divide(a, b int) (float64, error) {
	if b == 0 {
		return 0, errors.New("You cannot divide by zero")
	}
	return float64(a) / float64(b), nil
}
