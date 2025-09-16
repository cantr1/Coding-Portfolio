package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	secret := rand.Intn(100) + 1

	var guess int
	for {
		fmt.Print("Enter your guess: ")
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
