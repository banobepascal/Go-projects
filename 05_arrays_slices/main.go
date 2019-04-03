package main

import "fmt"

func main() {
	//Arrays
	// var fruitArr [3]string

	// //Assign values
	// fruitArr[0] = "Apple"
	// fruitArr[1] = "Orange"
	// fruitArr[2] = "Mango"

	//Declare and assign
	// fruitArr := [3]string{"Apple", "Orange", "Mango"}

	// fmt.Println(fruitArr)
	// fmt.Println(fruitArr[2])

	friutArr := []string{"Apple", "Orange", "Mango"}

	fmt.Println(len(friutArr))
	fmt.Println(friutArr[1:2])
}