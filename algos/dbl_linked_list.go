package algorithms

// The function `insertListElements` creates a double-linked list, inserts elements from 1 to 'n', and returns the list.

import (
	"container/list" // Import the Go standard library package for double-linked lists
	"fmt"
)

// insertListElements creates a double-linked list and adds elements from 1 to n
func insertListElements(n int) *list.List {
	// Create a new instance of a double-linked list
	lst := list.New()

	// Loop from 1 to 'n' and insert each value into the list
	for i := 1; i <= n; i++ {
		// Insert the value 'i' at the back of the list (PushBack appends to the end)
		lst.PushBack(i)
	}

	// Return the populated list
	return lst
}

// LinkedList initializes the process and prints the elements of the linked list
func LinkedList() {
	fmt.Println("----- Double Linked List -----")
	// Set the total number of elements to be inserted in the list
	n := 5
	fmt.Printf("Creates a Linked List with %d elements in it.\n", n)
	// Call insertListElements to create a list and insert elements from 1 to 'n'
	myList := insertListElements(n)

	// Loop through the linked list from the front to the back
	// myList.Front() returns the first element, and e.Next() moves to the next element in the list
	for e := myList.Front(); e != nil; e = e.Next() {
		// Print the value of the current element in the list
		fmt.Println(e.Value)
	}
}
