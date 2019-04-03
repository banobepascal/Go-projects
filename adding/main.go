package main

import "fmt"

func getSum( a, b int) int {
	return a + b
}

func main() {
	fmt.Println(getSum(3, 4))
}