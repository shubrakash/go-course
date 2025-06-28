package basics

import "fmt"

func variadicMain() {
	// Ellipsis ...
	// func functionName(parameter1 type1, parameter2 type2, parameter3 ...type3) returnType {
	// function body
	// }

	fmt.Println("Sum of 1,2,3,4,5 is", sumFromAStart(10, 1, 2, 3, 4, 5))
	fmt.Println("Sum of 18, 23 is", sumFromAStart(-10, 18, 23))

	/// slice
	nums := []int{1, 2, 4, 5, 6, 7}
	/// expanding a slice like `nums...`
	fmt.Println("Sum of", nums, "is", sumFromAStart(100, nums...))

	fmt.Println(sum("Sum of ", 1, 2, 3, 4, 5))
}

func sumFromAStart(startFrom int, nums ...int) (total int) {
	total = startFrom
	for _, v := range nums {
		total += v
	}
	return
}

func sum(prefixMsg string, nums ...int) (string, int) {
	total := 0
	returnString := prefixMsg
	for _, v := range nums {
		total += v
		returnString += fmt.Sprintf("%d, ", v)
	}
	return returnString + "is: ", total
}
