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
	fruitArr := [5]string{"Apple", "Orange", "Mango", "Grape", "Juice"}

	fmt.Println(fruitArr)
	fmt.Println(fruitArr[1:2])

	// Creating a slice from an array

	slice1 := fruitArr[1:4]
	slice2 := fruitArr[:3]
	slice3 := fruitArr[2:]
	slice4 := fruitArr[:]

	// numElementsCopied := copy(slice2, fruitArr)

	fmt.Println("slice1 =", slice1)
	fmt.Println("slice2 =", slice2)
	fmt.Println("slice3 =", slice3)
	fmt.Println("slice4 =", slice4)


	// Creating a slice using make
	s := make([]int, 5, 10)
	fmt.Printf("s = %v, len = %d, cap = %d\n", s, len(s), cap(s))

	// No array number
	friutArr := []string{"Apple", "Orange", "Mango"}

	fmt.Println(len(friutArr))
	fmt.Println(friutArr[1:2])

	// Iterating over an array
	names := [3]string{"Pascal", "Josh", "Mathew" }

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])

		//finding sum by Iterating over an array
		a := [5]float64{3.2, 4.5, 6.5, 3.6, 5.6}
		sum := float64(0)

		// for i := 0; i < len(a); i++ {
		// 	sum = sum + a[i]
		// }

		for _, value := range a {
			sum = sum + value
		}

		fmt.Printf("Sum of all the elements in array %v = %f\n", a, sum)
		break
	}
}