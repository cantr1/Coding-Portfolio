package main

import "fmt"

func main() {
	// Create var to store continuance
	var cont string = "Y"

	// Execute while loop
	for cont != "EXIT" {
		// Take operation
		var operation string
		fmt.Print("Enter operation: {+,-,*,/} ")
		fmt.Scanln(&operation)
		// Create variables
		var x int
		var y int
		fmt.Print("Enter two numbers: ")
		fmt.Scanf("%d %d", &x, &y)

		// Process calculation
		switch operation {
		case "+":
			fmt.Printf("%d + %d = %d\n", x, y, x+y)
		case "-":
			fmt.Printf("%d - %d = %d\n", x, y, x-y)
		case "*":
			fmt.Printf("%d * %d = %d\n", x, y, x*y)
		case "/":
			if y == 0 {
				fmt.Println("ERROR: Division by zero...")
				continue
			}
			fmt.Printf("%d / %d = %d\n", x, y, x/y)
		default:
			fmt.Println("Unrecognized operation...")
		}

		fmt.Printf("Would you like to continue: (Y or EXIT) ")
		fmt.Scanln("%s", &cont)
	}
}
