package main

import "fmt"

func main() {
	// using var keyword, notice string and int are not needed bc it will be inferred.
	var name string = "Bryan"
	var age int = 37
	const isCool = true //const make variable constant, cannot be changed

	// using shorthand
	// this method must be used with a function
	secondName := "Steph"
	size := 21.3

	// variable can be assigned on a single line
	firstName, email := "John", "john@test.com"

	// print variables
	fmt.Println(name, age, isCool, secondName, size, firstName, email)
	// print format see the fmt documenation for proper verb flags
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", isCool)
	fmt.Printf("%T\n", secondName)
	fmt.Printf("%T\n", size)
	fmt.Printf("%T\n", firstName)
	fmt.Printf("%T\n", email)

}
