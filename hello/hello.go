package hello

import (
	"fmt"
	"go-tutorial/greetings"

	"rsc.io/quote"
)

func Say_something() {
	fmt.Println("hello")
	fmt.Println(quote.Go())
	greetings.Hello("syamsul")
}
