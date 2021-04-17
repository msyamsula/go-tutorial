package spanish

type SpanishBot struct{}

// because SpanishBot satisfy every function in Bot Interface
// english bot is Bot and SpanishBot at the same time
func (SpanishBot) GetGreeting() string {
	return "Orale Padron"
}
