package main

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func allTheStrings() {
	// https://pkg.go.dev/strings#pkg-functions
	fmt.Printf("\n------ Strings Example --------\n")
	fmt.Printf("\nTrue or False?\n\n")

	var str string = "This is an example of a string"
	// Prefixes and Suffixes
	fmt.Printf("Does the string \"%s\" have a prefix \"%s\"?", str, "Th")
	// HasPrefix tests whether the variable 's' begins with a prefix of `prefix`
	// example: strings.HasPrefix(s, prefix string) bool
	fmt.Printf("\n%t\n\n", strings.HasPrefix(str, "Th"))

	fmt.Printf("Does the string \"%s\" have the suffix \"%s\"?", str, "ting")
	// HasSuffix tests whether the variable 'str' ends with a suffix 'ting'
	// example: strings.HasSuffix(s, suffix string) bool
	fmt.Printf("\n%t\n\n", strings.HasSuffix(str, "ting"))

	// Contains
	fmt.Printf("Does the string \"%s\" contain the substring of \"%s\"?", str, "example")
	fmt.Printf("\n%t\n\n", strings.Contains(str, "example"))
	//strings.Contains(s, substr string) bool

	fmt.Println("\n\nOther Examples:")

	// Index, returns the index of the first instance of str in s, or -1 if str is not present in s:
	// strings.Index(s, str string) int
	fmt.Printf("\nThe first occurence of \"is\" in the string \"%s\" is at position: %d\n", str, strings.Index(str, "is"))

	// LastIndex returns the index of the last instance of str in s, or -1 if str is not present in s:
	// strings.LastIndex(s, str string) int
	fmt.Printf("\nThe last occurence of \"is\" in the string \"%s\"\" is at position: %d\n", str, strings.LastIndex(str, "is"))

	// Counts counts the number of non-overlapping instances of the substring str in s
	// strings.Count(s, str string) int
	fmt.Printf("\nThe number of \"i\"'s in the string %s is: %d\n", str, strings.Count(str, "i"))

	// Replace an old string with a new string
	// strings.Replace(str, old, new string, n int)

	// ToLower, returns a *copy* of the string s with all Unicode characters mapped to their lower case
	// strings.ToLower(s) string
	// All uppercase is obtained with:
	// strings.ToUpper(s) string
	var camel string = "The Quick Brown Fox Jumps Over The Lazy Dog"

	fmt.Printf("strings.ToLower returns a copy of the string \n\"%s\" \nAnd converts it to all lowercase: \n\"%s\"\n", camel, strings.ToLower(camel))
	fmt.Printf("While strings.ToUpper returns a copy of that string as \n\"%s\"\n", strings.ToUpper(camel))
	fmt.Printf("Note that strings are immutable in Go, so these only return a copy, and doesn't modify the original\n\"%s\"\n", camel)

	// TrimSpace can be used to remove all leading and trailing whitespaces as:
	// strings.TrimSpace(s)
	// Trim can also be used to trim a specific str from a s, using
	// strings.Trim(s, str)

	// Splitting a string

	fmt.Printf("The strings.fields(s) splits the string s around each instance of one or more consecutive white space characters\nand returns a slice of substrings []string of s or and empty list if s contains only one whitespace\n")
	//
	fieldsString := strings.Fields(camel)
	for i, s := range fieldsString {
		fmt.Printf("index: %d, value: %s\n", i, s)
	}

	fmt.Printf("\nThe *strings.Split(s, sep)* works the same as Fields, but splits around the \"sep\". \nThe sep can be a seprator of any character (:,;,-,...) or any other separator string\n")
	fruitString := "apples,bannanas,oranges,pears"
	fmt.Printf("For example the string \"%s\" can be turned into:\n", fruitString)
	splitString := strings.Split(fruitString, ",")

	for i, s := range splitString {
		fmt.Printf("\nindex: %d, value: %s", i, s)
	}

	// Joining over a slice
	fmt.Printf("\n\nThe strings.Join(sl []string, sep string) results in a string containing all the elements of the slice 'sl' seprated by 'sep'\n")
	fmt.Printf("So you can turn the previously split string back into: \n\"%s\"\n", strings.Join(splitString, ","))

	// Reading from a string
	fmt.Printf("The 'strings' package also has a function called 'strings.NewReader(str)'\nThis produces a *pointer* to a reader value, that provides amongst other things the following functions to operate on str:")
	fmt.Printf("\n - Read() to read a []byte\n - ReadByte() to read the next byte from the string.\n - ReadRune() to read the next rune from the string.\nhttps://pkg.go.dev/strings#NewReader\n")

	// Conversion to and from a string
	fmt.Printf("the package 'strconv' contains a few variables to calculte the size in bits of the int of the platform on which the program runs. \nLike this: %v\n", strconv.IntSize)
	strVal := "42"
	//intToStr := 69
	intVal, iErr := strconv.Atoi(strVal)
	initFloatNum := 42.69
	floatVal := strconv.FormatFloat(initFloatNum, 'g', 2, 64)

	if iErr != nil {
		fmt.Printf("Error: %v", iErr)
		os.Exit(0)
	}

	fmt.Printf("%s: %v\n%d: %v\n%v: %v\n", strVal, reflect.TypeOf(strVal), intVal, reflect.TypeOf(intVal), floatVal, reflect.TypeOf(floatVal))
}
