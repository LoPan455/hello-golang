package main

import (
	"example.com/greetings"
	"fmt"
	"log"
	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(quote.Go())

	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	message, err := greetings.Hello("Tom")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)

	// Let's try again, this time failing to pass a string to greetings.Hello()
	secondMessage, err := greetings.Hello("")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(secondMessage)
}
