package basics

import "fmt"

func functionsMain() {
	// func <name>(parameters list) returnType {
	// code block
	// return value
	// }

	fmt.Println(add(1, 2))

	// ananoymous function directly called upon declaration
	func() {
		fmt.Println("Hello ananoymous function!")
	}()
	// ananoymous function assigned
	greet := func() {
		fmt.Println("Greeting ananoymous function!")
	}
	greet()

	/// assigning function to a variable
	operation := add

	result := operation(2, 3)
	fmt.Println(result)

	// using function which as another function as an argument
	result1 := applyOperation(2, 3, add)
	fmt.Println("2+3:", result1)

	multiplyBy2 := createMulitplier(2)
	fmt.Println("6*2: ", multiplyBy2(6))
}

func add(a, b int) int {
	return a + b
}

// function as an argument
func applyOperation(a, b int, operation func(int, int) int) int {
	return operation(a, b)
}

// function as a return type
func createMulitplier(factor int) func(int) int {
	return func(x int) int {
		return factor * x
	}
}
