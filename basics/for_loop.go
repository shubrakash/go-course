package basics

import "fmt"

func forLoopMain() {
	// // simple iteration over a range
	// for i := 1; i <= 5; i++ {
	// 	fmt.Println(i)
	// }

	// // iteration over collection
	// numbers := []int{1, 2, 3, 4, 5, 6}
	// for i, val := range numbers {
	// 	fmt.Printf("index: %d, value: %v\n", i, val)
	// }

	// // continue and break
	// for i := 0; i <= 10; i++ {
	// 	if i%2 == 0 {
	// 		continue
	// 	}
	// 	fmt.Println("Odd number: ", i)
	// 	if i == 5 {
	// 		break
	// 	}
	// }

	// // pyramid pattern
	// rows := 10
	// // outer loop
	// for i := 1; i <= rows; i++ {
	// 	//inner loop for spaces before stars
	// 	for j := 1; j <= rows-i; j++ {
	// 		fmt.Print(" ")
	// 	}
	// 	for k := 1; k <= 2*i-1; k++ {
	// 		fmt.Print("*")
	// 	}

	// 	fmt.Println()
	// }

	// consise int range iteration
	for i := range 10 {
		fmt.Println(i)
	}
}
