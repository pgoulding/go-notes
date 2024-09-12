package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct {
}

type spanishBot struct {
}

func showInterfaces() {
	fmt.Println("\n----- Interfaces ------\n")
	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (eb englishBot) getGreeting() string {
	// custom logic for generating an english greeting,
	return "Hi there!"
}

func (sb spanishBot) getGreeting() string {
	return "Hola!"
}
