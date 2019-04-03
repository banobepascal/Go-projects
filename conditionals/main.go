package main

import "fmt"

func main() {
	var x = 28

	if x < 5 {
		fmt.Println(" is a multiple of 5")
	}else if x == 28 && x >= 30 {
		fmt.Println(" is not a multiple of 5")
	}else if x > 30 || x <= 21 {
		fmt.Println("You are fake")
	}else {
		fmt.Println("you are obsolute")
	}
}

// func main() {

// 	if n := 9; n%2 == 0 {
// 		fmt.Printf("%d is a multiple of 2", n)
// 	} else {
// 		fmt.Println("Answer is wrong")
// 	}

// }
