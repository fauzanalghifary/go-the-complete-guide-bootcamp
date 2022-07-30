package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}
type bot interface {
	getGreeting() string
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (eb englishBot) getGreeting() string {
	// VERY CUSTOM LOGIC
	return "Hi There!"
}

func (sb spanishBot) getGreeting() string {
	// VERY CUSTOM LOGIX
	return "Hola!"
}
