package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func showStructs() {
	fmt.Println("Structs")

	// // Updating Structs
	// // opt 1
	alex := person{
		firstName: "Alex",
		lastName:  "Anderson",
		contact: contactInfo{
			email:   "alex@gmail.com",
			zipCode: 94000,
		},
	}

	// // opt 2
	//alex := person{"Alex", "Anderson"}

	// // opt 3
	// var alex person
	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"

	fmt.Println(alex)
	fmt.Printf("%+v", alex)
}
