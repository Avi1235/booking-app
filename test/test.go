package main

import "fmt"


func main() {
	var testing uint

	n, err := fmt.Scan(&testing)
	fmt.Printf("%v\n", n)
	fmt.Printf("%v\n", err)
	if (testing > 0) {
		fmt.Printf("Hola")
	} else {
		var userFirstName string
		var userLastName string
		var userEmail string
	
		fmt.Printf("Enter your first name: ")
		fmt.Scan(&userFirstName)
		fmt.Printf("Enter your last name: ")
		fmt.Scan(&userLastName)
		fmt.Printf("Enter your email: ")
		fmt.Scan(&userEmail)
	}
}