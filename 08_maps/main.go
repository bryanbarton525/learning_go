package main

import "fmt"

func main() {
	// define a map
	emails := make(map[string]string)

	//assign key values
	emails["Bob"] = "bob@gmail.com"
	emails["sharon"] = "sharon@gmail.com"
	emails["mike"] = "mike@gmail.com"

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["Bob"])

	//delete from map
	delete(emails, "Bob")
	fmt.Println(emails)
}
