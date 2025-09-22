package main

import "fmt"

type Book struct {
	Title       string
	Author      string
	Pages       int
	CurrentPage int
}

func (b Book) isLong() bool {
	return b.Pages > 500
}

func (b Book) Progress() float64 {
	return ((float64(b.CurrentPage) / float64(b.Pages)) * 100)
}

func main() {
	b := Book{Title: "The Go Programming Language", Author: "Alan Donovan", Pages: 300, CurrentPage: 134}
	fmt.Printf("%s by %s, %d pages\n", b.Title, b.Author, b.Pages)
	if b.isLong() {
		fmt.Printf("%s is a long book.\n", b.Title)
	} else {
		fmt.Printf("%s is a not long book.\n", b.Title)
	}
	fmt.Printf("You are %.2f%% through the book\n", b.Progress())
}
