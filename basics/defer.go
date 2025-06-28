package basics

import "fmt"

func deferMain() {
	deferredProcess()
}

func deferredProcess() {
	i := 10
	defer fmt.Println("1st deferred statement", i)
	defer fmt.Println("2nd deferred statement", i)
	defer fmt.Println("3rd deferred statement")
	i++
	fmt.Println("Normal statement")
	fmt.Println(i)
}
