package main

import (
	"fmt"
	"go-tutorial/object/person"
)

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func (englishBot) getGreeting() string {
	return "Hello from england"
}

func (spanishBot) getGreeting() string {
	return "Orale Padron"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func main() {
	// struct example, struct is imported from object/person
	alex := person.Person{
		FirstName: "alex",
		LastName:  "anderson",
		ContactInfo: person.ContactInfo{
			Email: "msyamsula1995@gmail.com",
			Zip:   14095,
		},
	}

	alex.FirstName = "syamsul"
	(&alex).Update("wakwaw")
	alex.Print()

	// map example
	colors := map[string]string{
		"red":   "black",
		"white": "green",
	}

	colors["red"] = "blue"
	delete(colors, "purple")
	colors["purple"] = "ungu"

	for key, item := range colors {
		fmt.Printf("key = %v: item = %v\n", key, item)
	}

	fmt.Printf("%v\n", colors)

	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
	// spanish.Greeting()
	// english.Greeting()
}
