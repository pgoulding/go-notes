package main

import "fmt"

func showMaps() {
	fmt.Println("\n----- Maps -----\n")
	// Establishes that all the keys `[string]...` will be of Type String, and the values `[...]string` will be of Type String
	colors := map[string]string{
		"red":   "#fff0000",
		"green": "#00ff00",
	}
	fmt.Println(colors)
}
