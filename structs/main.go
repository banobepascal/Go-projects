package main 

import "fmt"
import "strconv"

// Define person struct
type Person struct {
	firstName string
	lastName string
	city string
	gender string
	age int
}

// Greeting method (value reciever)
func (p Person) greet() string {
	return "Hello my name is " + p.firstName + " " + p.lastName + " and i am " +
	strconv.Itoa(p.age)
}

// hasBirthday method (pointer reciever)
func (p *Person) hasBirthday() {
	p.age++
} 

// getMarried method (pointer reciever)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "m" {
		return
	} else {
		p.lastName = spouseLastName	
	}
}

func main() {
	// Init person using struct
	person1 := Person{firstName: "Samantha", lastName: "Smith", city: "Boston", gender: "f",
	age: 25 }

 	// Alternative
	person2 := Person{"Bobby",  "Lash", "London", "m", 35 }
	
 	person1.hasBirthday()
	person1.getMarried("Bobby")
	fmt.Println(person1.greet())

	person2.getMarried("claire")
	fmt.Println(person2.greet())
} 