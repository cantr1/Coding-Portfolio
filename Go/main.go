package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	// Generate a random number 1 - 100
	var secret int = rand.Intn(100) + 1

	fmt.Println("GUESSING GAME")

	var guess int

	for {
		fmt.Printf("Enter your guess: ")
		fmt.Scan(&guess)
		if guess < secret {
			fmt.Println("Too low...")
		} else if guess > secret {
			fmt.Println("Too high...")
		} else {
			fmt.Println("Correct!")
			break
		}
	}
}
