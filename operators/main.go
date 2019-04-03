package main

import (
		"fmt"
		"math"
)

func main() {
	 var a, b = 4, 5
	 var res1 = (a + b) * (a + b)/2  // Arithmetic operations

	 a++ //increment by 1

	 b += 10 //increment by 10

	 var res2 = a ^ b //Bitwise XOR

	 var r = 3.5
	 var res3 = math.Pi * r * r //operations on floating-point type

	 fmt.Printf("res1 : %v, res2 : %v, res3 : %v\n", res1, res2, res3)

}