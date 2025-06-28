package basics

import (
	"errors"
	"fmt"
)

func multipleReturnValuesMain() {

	// func <functionName>(parameter1 type1, parameter2 type2, ...) (returnType1, returnType2, ...) {
	// code block
	// return returnType1, returnType2, ...
	// }

	q, r := divide(10, 3)
	fmt.Printf("Quotient: %d, Remainder: %d\n", q, r)

	q, r = divide2(18, 3)
	fmt.Printf("Quotient: %d, Remainder: %d\n", q, r)

	result, err := whichIsGreater(3, 3)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}

func divide(a, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}

// named return value
// elegant way of defining name to return types and doing just return as `go` picks them automatically
func divide2(a, b int) (quotient, remainder int) {
	quotient = a / b
	remainder = a % b
	return
}

func whichIsGreater(a, b int) (string, error) {
	if a > b {
		return "a is greater than b", nil
	} else if b > a {
		return "b is greater than a", nil
	} else {
		return "", errors.New("Not able to compare which is greater")
	}
}
