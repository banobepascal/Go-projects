package main

import "fmt"

// func main() {
// 	switch dayOfWeek := 8; dayOfWeek {
// 	case 1: fmt.Println("Monday")
// 	case 2: fmt.Println("Tuesday")
// 	case 3: fmt.Println("wednesday")
// 	case 4: fmt.Println("Thursday")
// 	case 5: fmt.Println("Friday")
// 	case 6: {
// 		fmt.Println("Saturday")
// 		fmt.Println("Thats right boss")
// 	} 
// 	case 7: {
// 		fmt.Println("Sunday")
// 		fmt.Println("Weekend. Yaay!")
// 	}
// default: fmt.Println("Invalid day")
// 	}
// }

func main() {
	switch dayOfWeek := 8; dayOfWeek {
	case 1, 2, 3, 4, 5:
		fmt.Println("weekday")
	case 6, 7:
		fmt.Println("weekend")
	default:
		fmt.Println("invalid day")
	}
}