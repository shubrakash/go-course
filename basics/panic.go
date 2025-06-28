package basics

import "fmt"

// panic(interface{})
// interface{} is like `any`, means it can be of any type.

// once panicked, rest of the program terminates
func panicMain() {
	process1(10)
	// process(-2)

	process2(-3)
}

func process1(input int) {
	if input < 0 {
		panic("Input can't be negative")
	}
	fmt.Println("Processing input:", input)
}

func process2(input int) {
	defer fmt.Println("Deferring 1")
	defer fmt.Println("Deferring 2")
	if input < 0 {
		fmt.Println("Before panic...")
		panic("Input can't be negative")
		// fmt.Println("After panic...")	//anything after panic is unreachable
		// defer fmt.Println("After panic...")	//even deferred statement
	}
	fmt.Println("Processing input:", input)
}
