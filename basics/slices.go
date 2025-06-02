package basics

import (
	"fmt"
	"slices"
)

func slicesBasics() {
	//// ways of declaration
	// var variableName []type

	// var slice1 []int

	// var slice2 = []int{1, 2, 3}

	// slice3 := []int{9, 8, 7}

	// slice4 := make([]int, 5)

	//// slicing
	a := [5]int{1, 2, 3, 4, 5}
	slice5 := a[1:4]

	// fmt.Println("slice5", slice5)

	//// appending to slice
	slice5 = append(slice5, 6, 7, 8)
	// fmt.Println("slice5", slice5)

	//// copy of a slice
	sliceCopy := make([]int, len(slice5))
	copy(sliceCopy, slice5)

	// fmt.Println("sliceCopy", sliceCopy)

	//// slice without any value
	// var nilSlice []int

	// for i, v := range slice5 {
	// 	fmt.Println("Index:", i, "Value:", v)
	// }

	slice5[3] = 50
	fmt.Println("eLement at index 3 in slice5", slice5[3])

	//// equality: slices.Equal
	if slices.Equal(slice5, sliceCopy) {
		fmt.Println("slice5 and sliceCopy are equal") // doesn't print because slice5 was modified
	}

	//// multi-dimersional slice

	twoD := make([][]int, 5)
	for i := range len(twoD) {
		twoD[i] = make([]int, i+1)
		for j := range len(twoD[i]) {
			twoD[i][j] = i + j
		}
	}

	fmt.Println(twoD)

	slice6 := slice5[2:4]
	fmt.Println("length of slice6", len(slice6))   // O: 2; count of elements
	fmt.Println("capacity of slice6", cap(slice6)) // O: 6; size of the referencing array; for slice6, it was slice5 which has 6 elements

}
