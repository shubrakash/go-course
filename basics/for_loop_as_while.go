package basics

import "fmt"

func forLoopAsWhile() {
	// infinite loop
	// for {
	// 	fmt.Println("Hello")
	// }

	// while loop with continue
	num := 0
	for num <= 10 {
		if num%2 == 0 {
			num++
			continue
		}
		fmt.Println("Odd number:", num)
		num++
	}
}
