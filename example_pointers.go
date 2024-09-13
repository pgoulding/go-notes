package main

import (
	"fmt"
)

func examplePointers() {
	i1 := 5
	// while the variable (i1) will stay the same,
	// the pointer (&i1) can be different every time because it's pointing to a distinct
	// place in memory where the integer is stored.
	fmt.Printf("An integer: %d, and it's location in memory: %p\n", i1, &i1)

	// pointer variable
	var intPtr *int
	// Storing address of i1 in pointer variable
	intPtr = &i1
	fmt.Printf("The value at memory location %p is %d\n", intPtr, *intPtr)

	// initalize the variable
	s := "good bye"
	fmt.Println(s)

	// Declare the pointer variable and set it to the address of s
	var sPtr *string = &s
	// dereference the sPtr and change the value to "ciao"
	*sPtr = "ciao"
	fmt.Printf("Here is the pointer sPtr: %p\n", sPtr)
	fmt.Printf("Here is the string *sPtr: %s\n", *sPtr)
	fmt.Printf("here is the string sPtr: %s\n", s)

}
