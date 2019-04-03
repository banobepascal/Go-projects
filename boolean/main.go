package main

import "fmt"

func main() {
	truth := 20 > 10
	falsehood := 20 != 40

	// shorthand

	res1 := 20 >= 10 && 10 == 10
	res2 := 40 >= 20 || 5%2 == 0

	fmt.Println(truth, falsehood, res1, res2)

}