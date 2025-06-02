package basics

import "fmt"

func arrays() {
	// //// declare array variable
	// // var variableName [size]type
	// var numbers [5]int
	// // auto initialized with 0
	// fmt.Println(numbers)

	// //// modify array at index
	// numbers[0] = 50
	// fmt.Println(numbers)

	// fruits := [4]string{"Apple", "Banana", "Orange", "Kiwi"}
	// fmt.Println("Fruits array:", fruits)

	// //// read value at index
	// fmt.Println("Fruit at 2:", fruits[1])

	// //// create copy of array
	// originalArray := [3]int{1, 2, 3}
	// copiedArray := originalArray

	// fmt.Println("Original array:", originalArray)
	// fmt.Println("Copied array:", copiedArray)

	// //// modifying copiedArray doesn't modify the orginal array
	// copiedArray[0] = 50
	// fmt.Println("Original array:", originalArray)
	// fmt.Println("Copied array:", copiedArray)

	// fmt.Println("length of array is", len(numbers))

	// for i := 0; i < len(numbers); i++ {
	// 	fmt.Println("Element at", i, ":", numbers[i])
	// }

	// for i, v := range numbers {
	// 	fmt.Printf("Index %d: Value %d\n", i, v)
	// }

	// //// underscore is for blank identifier, used to store not be used value
	// for _, v := range numbers {
	// 	fmt.Printf("Value %d\n", v)
	// }

	// //// multi-dimensional array
	// var matrix = [3][3]int{
	// 	{1, 2, 3},
	// 	{4, 5, 6},
	// 	{7, 8, 9},
	// }

	// fmt.Println("Matrix:", matrix)

	//// create pointer to an array
	originalArray := [3]int{1, 2, 3}

	// reference the original array
	var copiedArray *[3]int
	copiedArray = &originalArray

	// simple way
	// copiedArray := &originalArray

	copiedArray[0] = 100
	fmt.Println("Original array:", originalArray)
	fmt.Println("Copied array:", *copiedArray)
}
