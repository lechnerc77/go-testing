package calculator

import "errors"

func Add(left, right int) int {
	sum := left + right
	return sum
}

func Divide(left, right int) (int, int, error) {

	if right == 0 {
		return 0, 0, errors.New("Cannot divide by zero")
	}

	quotient := left / right
	remainder := left % right
	return quotient, remainder, nil
}
