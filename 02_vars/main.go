package main

import "fmt"


func main() {
	
	age, isCool := 37, true
	name, email := "Pascal", "pascal@gmail.com"

	fmt.Println(name, age, isCool, email)
	fmt.Printf("%T\n", name)
}