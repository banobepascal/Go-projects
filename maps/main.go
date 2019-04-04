package main

import "fmt"

// func main() {
// 	var m = make(map[string]int)
// 	fmt.Println(m)

// 	if m == nil {
// 		fmt.Println("m is nil")
// 	}else {
// 		fmt.Println("m is not nil")
// 	}

// 	m["one hundred"] = 100
// 	fmt.Println(m)
// }

func main() {
	var faceMatch = make(map[string]string)

		// Adding keys to a map
		faceMatch["Bonnie"] = "Luque"
		faceMatch["Pascal"] = "Banobe"
		faceMatch["Brooke"] = "Stanley"
		faceMatch["Jose"] = "Kodek"
	
	fmt.Println(faceMatch)

	// Overwrite a key
	faceMatch["Bonnie"] = "Buhnie"
	delete(faceMatch, "Jose")
	fmt.Println(faceMatch)

	var employees = map[int]string{
		1004 : "John",
		1003 : "Pauline",
		1002 : "steve",
		1001 : "Doe",
	}

	printEmployee(employees, 1002)
	printEmployee(employees, 1001)

	if isEmployeeExists(employees, 10103) {
		fmt.Println("EmployeeId 1002 found")
	}

}

func printEmployee(employees map[int]string, employeeId int) {
	if name, ok := employees[employeeId]; ok {
		fmt.Printf("name = %s, and no. is %d, ok = %v\n", name, employeeId, ok)
	}else {
		fmt.Printf("Employee %d not found\n", employeeId)
	}
}

func isEmployeeExists(employees map[int]string, employeeId int) bool {
	_, ok := employees[employeeId]
	return ok
}