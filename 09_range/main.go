package main

import "fmt"

func main() {
	ids := []int{16, 54, 34, 89, 64, 25, 935, 485}
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// if not using the index use the "_" in place of the i
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// add id's together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum", sum)

	// Range with map

	// define a map
	emails := make(map[string]string)

	//assign key values
	emails["Bob"] = "bob@gmail.com"
	emails["sharon"] = "sharon@gmail.com"
	emails["mike"] = "mike@gmail.com"

	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}

	// only get the key
	for k := range emails {
		fmt.Println("Name: " + k)
	}

	// only get the value

	for _, v := range emails {
		fmt.Println("Email: " + v)
	}
}
