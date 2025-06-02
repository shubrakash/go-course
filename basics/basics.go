package basics

import "fmt"

const GRAVITY = 9.8
const pi = 3.14

type Employee struct {
	Name       string
	Email      string
	employeeID int
}

// func foo() {
// 	const (
// 		monday    int    = 0
// 		tuesday          = 1
// 		wednesday        = "Wed"
// 		thursday  string = "Thur"
// 	)
// 	var firstName string = "Shubham"
// 	var lastName = "Prakash"

// 	employee1 := Employee{Name: "Shubham", Email: "shubham@gmail.com", employeeID: 123}
// 	employee2 := Employee{"Shubham", "shubham@gmail.com", 123}
// }

func arithematicOperations() {
	var a, b int = 10, 3
	var result int

	// addition
	result = a + b
	fmt.Println("Addtion: ", result)

	// subtraction
	result = a - b
	fmt.Println("Subtraction: ", result)

	// multiplication
	result = a * b
	fmt.Println("Multiplication: ", result)

	// division
	/// (dividing `int`` by `int` will always be int value which is floored.)
	result = a / b
	fmt.Println("Division: ", result)

	// remainder
	result = a % b
	fmt.Println("Division Remainder: ", result)

	// power
	result = a ^ b
	fmt.Println("Power: ", result)

	const p float64 = 22 / 7.0
	fmt.Println(p)
}
