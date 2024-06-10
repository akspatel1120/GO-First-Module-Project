package main

import (
	"fmt"

	"golang.org/x/example/hello/reverse"
)

func main() {

	/*log.SetPrefix("Greetings: ")
	//SetFlags(0) disabled printing of time, source file, and  line number.
	log.SetFlags(0)

	//qoute.Go() is a function from the rsc.io/qoute/v4 module which returns a quote
	fmt.Println(quote.Go())

	names := []string{"Akshay", "Dipesh", "Chubby"}

	message, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)*/

	fmt.Println(reverse.String("Akshay"), reverse.Int(97))
}
