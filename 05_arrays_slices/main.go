package main

import "fmt"

func main() {
	// Arrays
	// var fruitArr [2]string

	//assing values
	/* 	fruitArr[0] = "Apple"
	   	fruitArr[1] = "Orange" */

	// Declare and assign
	fruitArr := [2]string{"Apple", "Orange"}

	fmt.Println(fruitArr)
	fmt.Println(fruitArr[1])

	//Slices
	slicedFruitArr := []string{"Apple", "Orange", "Grape"}
	fmt.Println(slicedFruitArr)
	fmt.Println(slicedFruitArr[2])
}
