package main

import "fmt"

func showMaps() {
	fmt.Println("\n----- Maps -----")
	// // establishes an empty map with key and values of type string
	// var colors map[string]string

	// colors := make(map[string]string)

	// Establishes that all the keys `[string]...` will be of Type String, and the values `[...]string` will be of Type String
	colors := map[string]string{
		"red":   "#fff0000",
		"green": "#00ff00",
		"white": "#ffffff",
	}

	fmt.Println(colors)

	// // to delete an element in the map:
	// delete(colors, "green")

	//fmt.Println(colors)
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
