package main

import "fmt"
import "math"

func main() {
	x := 15
	y := 10
	z := 15.8
	var w string = "something"

	fmt.Println(w)

	//if else
	if x <= y {
		fmt.Printf("%d is less then or equal to %d\n", x, y)
	} else {
		fmt.Printf("%d is less then %d\n", y, x)
	}

	color := "blue"

	if color == "red" {
		fmt.Printf("%s is a pretty color\n", color)
	} else if color == "blue" {
		fmt.Println("The color is blue")
	} else {
		fmt.Println("The color is not blue or red!")
	}

	// Switch
	switch color {
	case "red":
		fmt.Println("The color is red")
	case "blue":
		fmt.Println("The color is blue")
	default:
		fmt.Println("The color is NOT blue or red")
	}

	fmt.Println(math.Ceil(z))
}
