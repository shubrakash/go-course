package basics

import "fmt"

func rangeMain() {
	message := "Hello World!"

	for i, v := range message {
		fmt.Printf("Index: %d, Rune: %c\n", i, v)
	}
}
