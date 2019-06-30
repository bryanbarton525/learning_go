package main

import "fmt"

//Define the type of input next to the expected input, then define the return type following that.
func greeting(name string) string {
	return "Hello " + name
}

func getSum(num1 int, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Println(greeting("Bryan"))
	fmt.Println(getSum(3, 4))
}
