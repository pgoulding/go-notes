package main

import (
	"fmt"
	"strings"
)

func allTheStrings() {
	var str string = "This is an example of a string"

	fmt.Printf("True or False?\n\nDoes the string \"%s\" have a prefix \"%s\"?", str, "Th")
	fmt.Printf("\n%t\n\n", strings.HasPrefix(str, "Th"))

	fmt.Printf("Does the string \"%s\" have the suffix \"%s\"?", str, "ting")
	fmt.Printf("\n%t\n\n", strings.HasSuffix(str, "ting"))

}
