package greetings

import (
	"fmt"

	"rsc.io/quote"
)

func Hello(name string) string {
	message := fmt.Sprintf("Hello, %v. Welcome!", name)
	m := quote.Opt()
	fmt.Println(m)
	fmt.Println(message)
	return message
}

func Say() {
	fmt.Println("say")
}
