package main

import (
	"bufio"
	"fmt"
	"os"
)

// prints out the count of text lines that appear more than once in the input
// It reads from stdin, or from a list of named files.
func dup2() {
	fmt.Println("----- dup1 -----")
	// establishes a 'counts'
	counts := make(map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
