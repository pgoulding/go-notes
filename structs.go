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
	fmt.Println("\n--------Structs--------")

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

	// updating structs
	// &variable = Give me the memory address of the value this variable is pointing at.
	alexPointer := &alex
	alexPointer.updateName("Alexander")
	alex.print()
	//fmt.Println(alex)
}

// *pointer = give me the value the memory address is pointing at
func (pointerToPerson *person) updateName(newFirstName string) {
	// *person = This is a type description, meaning we're working with a pointer to a person
	//*pointerToPerson = This is an operator, it mneans we want to manipulate actual the value the pointer is referencing.
	(*pointerToPerson).firstName = newFirstName
}

// receiver function on struct
func (p person) print() {
	fmt.Printf("\n%+v", p)
}
