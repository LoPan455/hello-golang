package main

import (
	"example.com/greetings"
	"fmt"
	"log"
	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello, World!") // output
	fmt.Println(quote.Go())      // output

	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	message, err := greetings.Hello("Tom")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message) // output

	// Let's try again, this time failing to pass a string to greetings.Hello()
	//secondMessage, err := greetings.Hello("")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(secondMessage) // output

	names := []string{"Ed", "Sam", "Sally"}
	messages, err := greetings.Hellos(names)
	fmt.Println(messages) // output
}
