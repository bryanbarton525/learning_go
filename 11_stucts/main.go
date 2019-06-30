package main

import (
	"fmt"
	"strconv"
)

// Define person struct
type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

// Greeting method (value reciever 'makes no changes just uses existing data')
func (p Person) greet() string {
	return "Hello my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// hasBirthday method (pointer reciever 'changes something')
func (p *Person) hasBirthday() {
	p.age++
}

//getMarried (pointer recvr)
func (p *Person) getMarried(spouceLastName string) {
	if p.gender == "m" {
		return
	} else {
		p.lastName = spouceLastName
	}
}

func main() {
	// Init person using struct
	person1 := Person{"Bryan", "Barton", "Glenville", "m", 29}
	// Alternate
	person2 := Person{firstName: "Elle", lastName: "Feurstein", city: "Glenville", gender: "f", age: 30}

	fmt.Println(person1)
	fmt.Println(person2)

	//Get info from person
	fmt.Println(person1.age)

	// use method
	person1.hasBirthday()
	person1.getMarried("Barton")
	person2.getMarried("Barton")

	fmt.Println(person1.greet())
	fmt.Println(person2.greet())

}
