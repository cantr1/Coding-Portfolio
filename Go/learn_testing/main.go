package main

import "fmt"

func printName(name string) (string, error) {
	return fmt.Sprintf("Hello, %s", name), nil
}

func main() {
	var name string
	fmt.Println("Enter your name:")
	_, err := fmt.Scanln(&name)
	if err != nil {
		return
	}
	printName(name)
}
