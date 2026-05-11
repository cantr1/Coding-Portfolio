package basics

import (
	"fmt"
)

// Switch case
func getCreator(os string) string {
	var creator string
	switch os {
	case "linux":
		creator = "Linus Torvalds"
	case "windows":
		creator = "Bill Gates"

	// all three of these cases will set creator to "A Steve"
	case "macOS":
		fallthrough
	case "Mac OS X":
		fallthrough
	case "mac":
		creator = "A Steve"

	default:
		creator = "Unknown"
	}
	return creator
}

// Defer statement
// Defer is LIFO
func GetUsername(dstName, srcName string) (username string, err error) {
	// Open a connection to a database
	conn, _ := db.Open(srcName)

	// Close the connection *anywhere* the GetUsername function returns
	defer conn.Close()

	username, err = db.FetchUser()
	if err != nil {
		// The defer statement is auto-executed if we return here
		return "", err
	}

	// The defer statement is auto-executed if we return here
	return username, nil
}

// Embedded struct
type car struct {
  brand string
  model string
}

type truck struct {
  // "car" is embedded, so the definition of a
  // "truck" now also additionally contains all
  // of the fields of the car struct
  car
  bedSize int
}

lanesTruck := truck{
  bedSize: 10,
  car: car{
    brand: "Toyota",
    model: "Tundra",
  },
}

fmt.Println(lanesTruck.brand) // Toyota
fmt.Println(lanesTruck.model) // Tundra


/* Type assertions
type shape interface {
	area() float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	// ...
}

func printShapeInfo(s shape) {
	c, ok := s.(circle)
	if ok {
		radius := c.radius
		fmt.Printf("s is a circle, radius: %v\n", radius)
		return
	}
}
 */

func basics() {
	fmt.Println("Hello, Go!")
	// Type casting
	myInt := 42
	myFloat := float64(myInt)
	fmt.Printf("Integer: %d, Float: %.2f\n", myInt, myFloat)

	// Multiline declaration
	int1, int2, int3 := 1, 2, 3
	fmt.Println(int1, int2, int3)

	// Single line initial if
	if length := 22; length < 10 {
		fmt.Printf("Email must be at least 10 characters, is %d\n", length)
	}
}
