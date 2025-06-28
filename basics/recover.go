package basics

import "fmt"

func recoverMain() {
	process()
	fmt.Println("Working after process!!!!")
}

func process() {
	defer func() {
		// if r := recover(); r != nil {
		r := recover()
		if r != nil {
			fmt.Println("Recovered from: ", r)
		}
	}()

	fmt.Println("starting process...")
	panic("Something went wrong!")
	fmt.Println("Ending process")
}
