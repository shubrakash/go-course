package basics

import "fmt"

/// executed before anything else per package
/// It is executed only once per package, even if the package is used at multiple places
func init() {
	fmt.Println("Initializing the package 1")
}

/// if multiples inits, they are executed in their order of occurence in code
func init() {
	fmt.Println("Initializing the package 2")
}

func init() {
	fmt.Println("Initializing the package 3")
}

func initMain() {
	fmt.Println("Inside main function")

}
