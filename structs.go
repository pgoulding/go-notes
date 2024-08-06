package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

// for the contact field, you don't have to specify a field name,
// you can just enter contactInfo and it will assume it is a
// field name of contactInfo with the contactInfo struct
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
	alex.updateName("Alexander")
	alex.print()
	//fmt.Println(alex)
}

func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

// receiver function on struct
func (p person) print() {
	fmt.Printf("%+v", p)
}
