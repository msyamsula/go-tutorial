package english

// english bot struct
type EnglishBot struct{}

// because english bot satisfy every function in Bot Interface
// english bot is Bot and EnglishBot at the same time
func (EnglishBot) GetGreeting() string {
	return "Hello from england"
}
