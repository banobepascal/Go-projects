package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Printf("First positive number divisible by both 3 and 5 is %d\n", i)
			break
		}	
	}
}