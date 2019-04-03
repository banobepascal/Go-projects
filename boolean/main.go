package main

import "fmt"

func main() {
	var truth = 3 <= 5
	var falsehood = 10 != 10

	// shorthand

	var res1 = 10 > 20 && 20 == 20
	var res2 = 3 > 2 || 10%3 == 0
	
	fmt.Println(truth, falsehood, res1, res2)
}