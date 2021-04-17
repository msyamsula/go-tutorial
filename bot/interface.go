package bot

import "fmt"

// bot interface, it tell us that every struct that implement
// every function in Bot interface definition will automatically
// recognized as Bot along with its original struct
type Bot interface {
	GetGreeting() string
}

func PrintGreeting(b Bot) {
	fmt.Println(b.GetGreeting())
}
