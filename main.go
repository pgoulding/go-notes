package main

import (
	"fmt"
	"os"
)

func main() {

	for {
		fmt.Println("\n-------- Main Menu --------")
		fmt.Println("1. Strings example")
		fmt.Println("2. Structs example")
		fmt.Println("3. Maps example")
		fmt.Println("4. Interfaces Example")
		fmt.Println("99. Exit")
		fmt.Println("\nEnter your choice")

		var choice int
		_, err := fmt.Scan(&choice)

		if err != nil {
			fmt.Println("Invalid input, please enter a number.")
			// Clear stdin buffer after an invalid input
			var discard string
			fmt.Scanln(&discard)
			continue
		}

		switch choice {
		case 1:
			allTheStrings()
		case 2:
			showStructs()
		case 3:
			showMaps()
		case 4:
			showInterfaces()
		case 99:
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}

	//http_resp_interface()
	//geoMain()

	// // you'll need to supply a file to the command line arguments to run this one.
	// textinput()
	// dup1()
}
