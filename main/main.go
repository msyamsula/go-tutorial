package main

import (
	"fmt"
	"go-tutorial/bot"
	"go-tutorial/bot/english"
	"go-tutorial/bot/spanish"
	"go-tutorial/channel"
	"go-tutorial/object/person"
)

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

	eb := english.EnglishBot{}
	sb := spanish.SpanishBot{}

	bot.PrintGreeting(eb)
	bot.PrintGreeting(sb)

	// http.HttpRequest()

	// channel.StatusCheck()
	channel.StatusCheckConcurent()
	// spanish.Greeting()
	// english.Greeting()
}
