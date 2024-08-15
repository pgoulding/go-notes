package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

// Create a program that reads the content of a text file then prints its contents to the terminal.

// The file should be provided as an argument to the program when it is executed at the terminal.
// e.g. 'go run main.go myfile.txt' should open up the myfile.txt

// To read the arguments provided to a program, reference the variable 'os.Args' which is a slice of type string.

func textinput() {
	// os.Args is a slice that starts off with the compiled program name/location and then takes
	// the text input afterwards into a slice. Use the first index for the first argument.
	fileLocation := os.Args[1]

	// os.Open takes a file name and returns a POINTER to the file, and an error if unsuccessful.
	file, err := os.Open(fileLocation)
	// if an error occurs log it and exit the program.
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	// because the os.Open struct implements the Read function, we can use io.Copy
	// to take all the info in the file and pipe it to the os.Stdout variable.
	io.Copy(os.Stdout, file)
	fmt.Println(fileLocation)
}
