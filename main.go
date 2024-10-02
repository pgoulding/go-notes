package main

import (
	"fmt"
	"github.com/pgoulding/go-notes/algos"
	"github.com/pgoulding/go-notes/basic_examples"
)

func main() {
	fmt.Println("----- Main -----")
	// for {
	// 	fmt.Println("\n-------- Main Menu --------")
	// 	fmt.Println("1. Strings example")
	// 	fmt.Println("2. Structs example")
	// 	fmt.Println("3. Maps example")
	// 	fmt.Println("4. Interfaces Example")
	// 	fmt.Println("99. Exit")
	// 	fmt.Println("\nEnter your choice")

	// 	var choice int
	// 	_, err := fmt.Scan(&choice)

	// 	if err != nil {
	// 		fmt.Println("Invalid input, please enter a number.")
	// 		// Clear stdin buffer after an invalid input
	// 		var discard string
	// 		fmt.Scanln(&discard)
	// 		continue
	// 	}

	// 	switch choice {
	// 	case 1:
	// 		allTheStrings()
	// 	case 2:
	// 		//showStructs()
	// 	case 3:
	// 		//showMaps()
	// 	case 4:
	// 		//showInterfaces()
	// 	case 99:
	// 		fmt.Println("Exiting...")
	// 		os.Exit(0)
	// 	default:
	// 		fmt.Println("Invalid choice, please try again.")
	// 	}
	// }

	// allTheStrings()
	// timeAndDateExample()
	// http_resp_interface()
	// geoMain()
	// examplePointers()
	// exampleLabels()
	// traceExample()
	// exampleHigherOrder()
	// // you'll need to supply a file to the command line arguments to run this one.
	// textinput()
	// dup1()
	basicexamples.ExampleErrors()
	algorithms.BubbleSort([]int{1, 5, 3263, 2, 2435, 34, 64, 22, 12, 54})
}
