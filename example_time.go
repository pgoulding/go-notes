package main

import (
	"fmt"
	"time"
)

var week time.Duration

func timeAndDateExample() {

	// Time representations
	t := time.Now() // declare the time and initialize it at the start of the program
	fmt.Println(t)  // print out the time in default format

	fmt.Printf("%02d.%02d.%4d\n", t.Day(), t.Month(), t.Year()) // print out the format as Day.Month.Year

	t = time.Now().UTC() // set time in UTC format
	fmt.Println(t)

	fmt.Println(time.Now())

	// Calculating Times

	// declare the week variable and initialize it with a formula for seconds in a week (seconds, minutes, hours, days)
	week = 60 * 60 * 24 * 7 * 1e9 // multiply it by 1e9 because it must be in nanosecond

	// Add the previous formula to get a week from now
	weekFromNow := t.Add(week)

	// print out the form
	fmt.Println(weekFromNow)

	// formatting times
	fmt.Println(t.Format(time.RFC822)) // format the time is RFC822
	fmt.Println(t.Format(time.ANSIC))  // format the time in ANSIC

	// simplify the time to Year, Month, Day
	s := t.Format("2006 01 02")
	// print out the time.now and the simplified time
	fmt.Println(t, "=>", s)
}
