package main

import (
	"fmt"
	"strings"
)

type myBigTestStruct struct {
	collectedStrings map[int]testStruct
}

type testStruct struct {
	sampleString string
	charCount    map[rune]int
}

func (t testStruct) printCharCount() {
	for key, value := range t.charCount {
		fmt.Printf("Key: %s, Value: %d\n", string(key), value)
	}
}

func (t testStruct) removeEmptySpaces() {
	// This works ? pass by value vs pass by reference
	// This is more to test / learn the deletion of keys
	// can be avoided entriely by not writing to the map
	// when an empty space is seen
	delete(t.charCount, ' ')

	// Seems to work because it is pass by reference
}

func (t *testStruct) countChars() {
	t.charCount = make(map[rune]int) // this was uncommented but caused issues
	for _, val := range strings.ToLower(t.sampleString) {
		t.charCount[val]++ // this works because the zero value = 0
	}
}

func main() {
	var theBigOne myBigTestStruct = myBigTestStruct{}

	testString1 := "This module is used to test new features / scratchpad"
	testString2 := "Right now I'm getting better with maps"

	var firstTestStruct testStruct = testStruct{
		sampleString: testString1,
	}
	firstTestStruct.countChars()

	// a more concise way to define similar to the above
	secondTestStruct := testStruct{
		sampleString: testString2,
	}
	secondTestStruct.countChars()

	// Remove whitespace count - arbitrary action for learning
	firstTestStruct.removeEmptySpaces()
	secondTestStruct.removeEmptySpaces()

	// Parse to the big one
	theBigOne.collectedStrings = map[int]testStruct{
		1: firstTestStruct,
		2: secondTestStruct,
	}

	// Print the values of the maps
	for key, value := range theBigOne.collectedStrings {
		fmt.Printf("Struct #%d - String: %s\n", key, value.sampleString)
		value.printCharCount()
		fmt.Println()
	}
}
