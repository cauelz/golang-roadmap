package interfaces

// How interfaces work in Go:
// - Interfaces are not generic types. They are implicit, which means that we don't have to explicitly say that a type implements an interface.
// - Interfaces are a contract to help us manage types.
// - Interfaces are a way to create a function that can be used with different types.
// - Interfaces and Structs muust be defined in the same package.

// Interfaces in Go works with the following rules:
// - If a type defines a function with the same signature as the interface, then that type is said to implement the interface.
// - We don't need to explicitly say that a type implements an interface. If a type has a function with the same signature as the interface, then that type implements the interface.
type bot interface {
	getGreeting() string
}

type englishBot struct {}
type spanishBot struct {}

func printGreeting(b bot) {
	println(b.getGreeting())
}

// It is not necessary to declare a receiver variable if it is not used
// so we can remove the receiver variable from the function signature, for example:
// func (englishBot) getGreeting() string
// func (spanishBot) getGreeting() string
func (englishBot) getGreeting() string {
	return "Hi there!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}