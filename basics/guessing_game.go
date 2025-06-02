package basics

import (
	"fmt"
	"math/rand"
	"time"
)

func guessingGame() {
	limit := 100

	fmt.Println("Welcome to the Guessing game!")
	fmt.Println("We have a number between 1 and", limit)
	fmt.Println("Can you guess what it is?")

	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	target := random.Intn(limit) + 1

	var guess int
	for {
		fmt.Println("Enter your guess: ")
		fmt.Scanln(&guess)

		if guess == target {
			fmt.Println("Congratulations! You have guessed the number correctly")
			break
		} else if guess < target {
			fmt.Println("Too low! Go higher")
		} else {
			fmt.Println("Too high! Go lower")
		}
	}
}
