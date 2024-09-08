package main

import (
	"bufio"
	"fmt"
	"os"
)

// Dup1 prints text of each line that appers more than once in the standard input
func dup1() {
	fmt.Println("----- dup1 -----")
	// establishes a 'counts'
	counts := make(map[string]int)
	// utilize the bufio package with a feature called Scanner
	// that reads inputs and breaks int into lines or words
	input := bufio.NewScanner(os.Stdin)

	// scans the input and puts each line into the counts map
	for input.Scan() {
		// the line is used a key for the map
		// and the value is an incrementing integer of times it's been found
		counts[input.Text()]++
	}

	// intiates a second loop, this time over the 'counts' map
	//
	for line, n := range counts {
		// if the value of that index is more than one
		if n > 1 {
			// display the %d decimal integer, %T add a tab, %s the string found
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

}
