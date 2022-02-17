package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {

	engBot := englishBot{}
	spaBot := spanishBot{}

	printGreeting(engBot)
	printGreeting(spaBot)

}

func printGreeting(b bot) {

	fmt.Println(b.getGreeting())

}

func (englishBot) getGreeting() string {

	return "Hi!"

}
func (spanishBot) getGreeting() string {

	return "Hola!"
}
