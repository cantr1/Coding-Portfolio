package main

import "fmt"

func main() {
	fmt.Println("This module is used to test new features / scratchpad")
	nums := []int{1, 2, 3}

	a := nums[:2]

	a = append(a, 100)

	fmt.Println(nums)
	fmt.Println(a)
}
