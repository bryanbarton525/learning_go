package main

import "fmt"

func main() {
	a := 5
	b := &a

	//normal print on pointer will show memory address
	fmt.Println(a, b)

	// print type
	fmt.Printf("%T\n", b)
	// use * to print value
	fmt.Println(*b)

	//change value of a using pointer
	*b = 10
	fmt.Println(*b)
}
